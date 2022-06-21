package cgodog

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v2"
)

var dockerSnapshot = make(map[string]bool)

// Env the environment declaration
type Env struct {
	// Real a list of services to be brought up via docker compose [eg, s3]
	Real []string
	// Mocks a list of services to be mocked
	Mocks []string
	Vars  []string
}

func (s *Suite) initEnv() {
	s.takeSnapshot()
	s.teardownEnv()
	vars := []string{"S3_ENDPOINT=http://s3:9000", "ENV_ENVIRONMENT=local"}
	env := s.initDeps()
	s.env.Vars = append(s.env.Vars, vars...)
	s.env.Vars = append(s.env.Vars, env...)

	// Init test container
	cont, err := s.dc.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: fmt.Sprintf("local.%s", s.serviceName),
			Healthcheck: &container.HealthConfig{
				Test: []string{"NONE"},
			},
			ExposedPorts: nat.PortSet{
				nat.Port(s.servicePort[1:]): {},
			},
			Env: s.env.Vars,
		},
		&container.HostConfig{
			AutoRemove:  true,
			NetworkMode: DockerNetwork,
			PortBindings: nat.PortMap{
				nat.Port(s.servicePort[1:]): []nat.PortBinding{{
					HostPort: fmt.Sprintf("3%s", s.servicePort[1:]),
				}},
			},
		},
		&network.NetworkingConfig{},
		nil,
		fmt.Sprintf("%s-test", s.serviceName),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = s.dc.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}
}

func (s *Suite) initDeps() []string {
	f, err := os.Open("env.yml")
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		log.Fatal(err)
	}

	if err = yaml.NewDecoder(bufio.NewReader(f)).Decode(&s.env); err != nil {
		log.Fatal(err)
	}

	_, err = s.dc.NetworkCreate( // nolint:gosec,errcheck
		context.Background(),
		DockerNetwork,
		types.NetworkCreate{
			CheckDuplicate: true,
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	s.initReal()
	return s.initMocks()
}

func (s *Suite) teardownEnv() {
	fmt.Println("tearing down containers to provide deterministic test env")
	cc, err := s.dc.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal("failed to list docker containers ", err)
	}
	for _, c := range cc {
		if err := s.dc.ContainerKill(context.Background(), c.ID, "SIGKILL"); err != nil {
			log.Fatal("failed to kill docker container ", err)
		}
	}
	// clear all volumes so we have a deterministic state
	if _, err := s.dc.VolumesPrune(context.Background(), filters.Args{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("environment cleared")
}

func (s *Suite) initReal() {
	if s.env.Real == nil || len(s.env.Real) == 0 {
		return
	}

	args := []string{
		"-f", "../../../bowstave_local_dev/docker-compose.yml",
		"-f", "../../../bowstave_local_dev/docker-compose.dev.yml",
		"-f", "../../../bowstave_local_dev/docker-compose.rt.yml",
		"-f", "../../../bowstave_local_dev/docker-compose.rt.override.yml",
	}

	if _, err := os.Stat("../../../bowstave_local_dev/docker-compose.dev.override.yml"); err == nil {
		args = append(args, "-f", "../../../bowstave_local_dev/docker-compose.dev.override.yml")
	}

	args = append(args, "up", "-d")
	args = append(args, s.env.Real...)

	var cmd *exec.Cmd

	// Check if "docker compose" is valid, then use it over "docker-compose"
	if err := exec.Command("docker", "compose").Run(); err != nil {
		cmd = exec.Command("docker-compose", args...)
	} else {
		args = append([]string{"compose"}, args...)
		cmd = exec.Command("docker", args...)
	}

	var b bytes.Buffer
	cmd.Stderr = &b
	if err := cmd.Run(); err != nil {
		fmt.Println(b.String())
		log.Fatal(err)
	}

	var waitKafka bool
	for _, svc := range s.env.Real {
		waitKafka = svc == "kafka"
		if waitKafka {
			break
		}
	}

	if !waitKafka {
		return
	}
	fmt.Println("waiting for kafka to become available")
	kkCtx := context.Background()
	var kkErr error = nil
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 10)
		kkErr = kakfaReady(kkCtx)
		if kkErr == nil {
			break
		}
	}

	if kkErr != nil {
		log.Fatal(errors.Wrap(kkErr, "waiting for kafka timed out"))
	}
	fmt.Println("kafka ready")
}

func (s *Suite) initMocks() []string {
	if s.env.Mocks == nil || len(s.env.Mocks) == 0 {
		return []string{}
	}

	protoSource, _ := filepath.Abs("../../../proto")
	stubSoruce, _ := filepath.Abs("mock")

	svcs := make([]string, 0)
	protos := make([]string, 0)
	for _, mock := range s.env.Mocks {
		svcName := strings.TrimSuffix(strings.ToUpper(mock), "_SERVICE")
		svcs = append(svcs, fmt.Sprintf("%s_SERVICE_ADDRESS=%s:%s", svcName, "mock-server", "22222"))
		protos = append(protos, fmt.Sprintf("%s.proto", mock))
	}
	cont, err := s.dc.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "tigh/grpc-mock:latest",
			ExposedPorts: nat.PortSet{
				nat.Port("22222"): {},
			},
			Entrypoint: []string{
				"grpc-mock",
				"-mock-addr=:22222",
				"-import-path=/proto",
				"-stub-dir=/stubs",
				fmt.Sprintf("-proto=%s", strings.Join(protos, ",")),
			},
		},
		&container.HostConfig{
			AutoRemove:  true,
			NetworkMode: DockerNetwork,
			PortBindings: nat.PortMap{
				nat.Port("22222"): []nat.PortBinding{{
					HostPort: "22222",
				}},
			},
			Mounts: []mount.Mount{{
				ReadOnly: true,
				Source:   protoSource,
				Target:   "/proto",
				Type:     "bind",
			}, {
				ReadOnly: true,
				Source:   stubSoruce,
				Target:   "/stubs",
				Type:     "bind",
			}},
		},
		&network.NetworkingConfig{},
		nil,
		"mock-server",
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.dc.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}

	return svcs
}

func (s *Suite) takeSnapshot() {
	cc, err := s.dc.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range cc {
		dockerSnapshot[c.ID] = true
	}
}

func kakfaReady(context.Context) error {
	conn, err := kafka.Dial("tcp", "localhost:29092")
	if err != nil {
		return err
	}
	defer conn.Close() //nolint:errcheck

	brokers, err := conn.Brokers()
	if err != nil {
		return err
	}
	if len(brokers) == 0 {
		return errors.New("kafka not ready")
	}

	return nil
}
