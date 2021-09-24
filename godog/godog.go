package cgodog

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors" //nolint:misspell
	"github.com/docker/docker/client"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	flag "github.com/spf13/pflag"
	"github.com/toorop/go-bitcoind"

	_ "github.com/lib/pq" // database driver
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// GodogOption a flag for determining if an option has been set
type GodogOption int

const (
	// DockerNetwork the name of the docker network
	DockerNetwork = "platform-services-network"

	// RandomEnabled forces the test to to run in random order, ignoring the --godog.random flag
	RandomEnabled GodogOption = 1
	// RandomDisabled forces the test to to run in senquential order, ignoring the --godog.random flag
	RandomDisabled GodogOption = 2
)

var godogOpts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
	Paths:  []string{"features"},
	Strict: true,
}

func init() {
	godog.BindCommandLineFlags("godog.", &godogOpts)
}

var (
	grpcClientConn *grpc.ClientConn
	btcClientConn  *bitcoind.Bitcoind
	dbClientConn   *sql.DB
)

// Suite the test suite to be ran
type Suite struct {
	serviceName string
	servicePort string
	btcInit     bool
	s3Init      bool

	dbCfg *DatabaseConfig
	DB    *sql.DB

	ts     godog.TestSuite
	tsInit func(*godog.TestSuiteContext)
	scInit func(*godog.ScenarioContext, *Context)

	dc *client.Client

	env Env
}

// Options config options for the test suite
type Options struct {
	// ServiceName the name of the service being tested
	ServiceName string
	// ServicePort the port for the service to be ran on
	ServicePort string
	// InitBitcoinBackend if true a bitcoin client will be created
	InitBitcoinBackend bool
	// InitS3 if true a minio client will be created
	InitS3 bool

	// ConcurrencyOverride if supplied, overrides the --godog.concurrency command line flag
	ConcurrencyOverride int

	// RandomOverride if supplied, overrides the --godog.random command line flag
	RandomOverride GodogOption

	// DatabaseConfig the config for the database. If provided a psql client will
	// be created
	DatabaseConfig *DatabaseConfig

	// TestSuiteInitializer an optional handler func for adding performing service specific
	// prep work.
	//
	// Any BeforeSuite calls defined here are performed after services have been initialised
	// Any AfterSuite calls defined here are performed before services have been destroyed
	TestSuiteInitializer func(*godog.TestSuiteContext)

	// ScenarioInitializer an optional handler func for adding performing service specific
	// prep work for scenarios.
	ScenarioInitializer func(*godog.ScenarioContext, *Context)
}

// DatabaseConfig the config for the local database
type DatabaseConfig struct {
	// Host the host name
	Host string
	// Port the port
	Port string
	// Username the authenticating username
	Username string
	// Password the authenticating password
	Password string
	// DatabaseName the name of the database
	DatabaseName string
}

// NewSuite created and returns a new *Suite based on the options provided
func NewSuite(opts Options) *Suite {
	flag.Parse()
	godogOpts.Paths = flag.Args()

	s := &Suite{
		serviceName: opts.ServiceName,
		servicePort: opts.ServicePort,
		btcInit:     opts.InitBitcoinBackend,
		s3Init:      opts.InitS3,

		dbCfg: opts.DatabaseConfig,

		tsInit: opts.TestSuiteInitializer,
		scInit: opts.ScenarioInitializer,
	}

	if opts.ConcurrencyOverride > 0 {
		godogOpts.Concurrency = opts.ConcurrencyOverride
	}
	if opts.RandomOverride > 0 {
		godogOpts.Randomize = int64(opts.RandomOverride) - 2
	}

	ts := godog.TestSuite{
		Name:                 opts.ServiceName,
		TestSuiteInitializer: s.initTestSuite,
		ScenarioInitializer:  s.initScenario,
		Options:              &godogOpts,
	}

	s.ts = ts

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	s.dc = cli
	return s
}

// Run runs the test suite
func (s *Suite) Run() int {
	return s.ts.Run()
}

func (s *Suite) initTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		s.initEnv()
		var err error
		if s.serviceName != "sars" {
			grpcClientConn, err = grpc.Dial(
				fmt.Sprintf(":3%s", s.servicePort[1:]),
				grpc.WithInsecure(),
				grpc.WithBlock(),
				grpc.WithKeepaliveParams(keepalive.ClientParameters{
					Time:    10 * time.Second,
					Timeout: 10 * time.Second,
				}),
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		if s.btcInit {
			btcClientConn, err = bitcoind.New(
				"localhost", 18332, "bitcoin", "bitcoin", false,
			)
			if err != nil {
				log.Fatal(err)
			}

			var req *http.Request
			req, err = http.NewRequest( //nolint:noctx
				"POST",
				"http://127.0.0.1:18332/",
				bytes.NewBuffer(
					[]byte(`{"id": "godog-testsuite", "jsonrpc": "1.0", "method": "generate", "params":[101]}`),
				),
			)
			if err != nil {
				log.Fatal(err)
			}
			req.SetBasicAuth("bitcoin", "bitcoin")
			req.Header.Add("Content-Type", "text/plain")

			var resp *http.Response
			resp, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close() //nolint:errcheck
		}

		if s.dbCfg != nil {
			if dbClientConn, err = sql.Open(
				"postgres",
				fmt.Sprintf(
					"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
					s.dbCfg.Username, s.dbCfg.Password, s.dbCfg.DatabaseName, s.dbCfg.Host, s.dbCfg.Port,
				),
			); err != nil {
				log.Fatal(err)
			}

			s.DB = dbClientConn
		}
	})

	if s.tsInit != nil {
		s.tsInit(ctx)
	}

	// Allow any user `AfterSuite` to be loaded in before tearing down the environment
	// incase the user wants to do perform some cleanup operations.
	ctx.AfterSuite(func() {
		if grpcClientConn != nil {
			if err := grpcClientConn.Close(); err != nil {
				fmt.Println(err)
			}
		}
		if dbClientConn != nil {
			if err := dbClientConn.Close(); err != nil {
				fmt.Println(err)
			}
		}
		s.teardownEnv()
	})
}

func (s *Suite) initScenario(ctx *godog.ScenarioContext) {
	testCtx := NewContext()
	ctx.BeforeScenario(func(sn *godog.Scenario) {

		testCtx.ScenarioID = sn.GetId()
		testCtx.TemplateVars["godog_scenario_id"] = sn.GetId()

		testCtx.AttachTemplater(sn.GetId())

		if s.s3Init {
			minioClient, err := minio.New("0.0.0.0:9000", &minio.Options{
				Creds:  credentials.NewStaticV4("admin", "bsvisking", ""),
				Secure: false,
			})
			if err != nil {
				log.Fatal(err)
			}

			testCtx.S3 = NewS3Context(minioClient)
		}

	})
	ctx.BeforeStep(func(st *godog.Step) {
		// Our stuff
	})

	if s.scInit != nil {
		s.scInit(ctx, testCtx)
	}

	ctx.AfterStep(func(st *godog.Step, err error) {
		// Our stuff
	})
	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		// Our stuff
	})

	// GRPC
	ctx.Step(`^I make a GRPC request to "([^"]*)"$`, IMakeAGRPCRequestTo(testCtx))
	ctx.Step(`^I make a GRPC request to "([^"]*)" with JSON "([^"]*)"$`, IMakeAGRPCRequestToWithData(testCtx))
	ctx.Step(`^I make a GRPC request to "([^"]*)" on port ([0-9]+) with JSON "([^"]*)"$`, IMakeAGRPCRequestToOnPortWithData(testCtx))
	ctx.Step(`^the GRPC code should be (\w*)$`, TheGRPCCodeShouldBe(testCtx))

	// HTTP
	ctx.Step(`^I make a (GET|POST|PATCH|DELETE) request to "([^"]*)"$`, IMakeAHTTPRequestTo(testCtx, fmt.Sprintf(":3%s", s.servicePort[1:])))
	ctx.Step(`^I make a (GET|POST|PATCH|DELETE) request to "([^"]*)" with JSON "([^"]*)"$`, IMakeAHTTPRequestToWithData(testCtx, fmt.Sprintf(":3%s", s.servicePort[1:])))
	ctx.Step(`^the HTTP response code should be (\d*)$`, TheHTTPResponseCodeShouldBe(testCtx))

	// Platform
	ctx.Step(`^I use the aliases:$`, IUseTheAliases(testCtx))
	ctx.Step(`^I use ([0-9]+) random alias(es)?$`, IUseRandomAliases(testCtx))
	ctx.Step(`^I fund (\d+) satoshis$`, IFundSatoshis(testCtx))
	ctx.Step(`^I fund and split by (\d+) satoshis$`, IFundAndSplitBySatoshis(testCtx))

	// Bitcoin
	ctx.Step(`^there (are|is) ([0-9]+) bitcoin account(s)?$`, ThereAreBitcoinWallets(testCtx))
	ctx.Step(`^I send ([0-9.]+) BTC to account ([0-9]+)$`, ISendBtcToWallet(testCtx))
	ctx.Step(`^I send BTC to accounts:$`, ISendBtcToAccounts(testCtx))
	ctx.Step(`^account ([0-9]+) sends ([0-9.]+) BTC to account ([0-9]+)$`, WalletSendsBtcToWallet(testCtx))
	ctx.Step(`^I store the "([^"]*)" of account (\d+) as "([^"]+)" for templating$`, IStoreTheOfWalletForTemplating(testCtx))

	// Tx
	ctx.Step(`^I prepare a transaction$`, IPrepareATransaction(testCtx))
	ctx.Step(`^I prepare a transaction from account ([0-9]+) to account ([0-9]+) for templating$`, IPrepareATranscationFromWalletToWallet(testCtx))
	ctx.Step(`^I add random data to the transaction$`, IAddRandomDataToTheTransaction(testCtx))
	ctx.Step(`^I use the outputs of transaction ([0-9]+) as inputs$`, IUseTheOutputsOfTransactionAsInputs(testCtx))
	ctx.Step(`^the transaction is signed by account ([0-9]+)$`, TheTransactionIsSignedByAccount(testCtx))
	ctx.Step(`^I store the hex of tx ([0-9]+) as "([^"]+)" for templating$`, IStoreTheHexOfTXAsForTemplating(testCtx))
	ctx.Step(`^I store the "([^"]+)" of tx ([0-9]+) as "([^"]+)" for templating$`, IStoreTheOfTXForTemplating(testCtx))
	ctx.Step(`^I store the "([^"]+)" of the working tx as "([^"]+)" for templating$`, IStoreTheOfTheWorkingTXForTemplating(testCtx))

	// Postgresql
	ctx.Step(`^the database is seeded with "([^"]*)"$`, TheDatabaseIsSeededWith(testCtx))

	// Minio
	ctx.Step(`^I delete from s3:$`, IDeleteFromS3(testCtx))
	ctx.Step(`^a clean account directory in s3$`, ACleanAccountDirectoryInS3(testCtx))
	ctx.Step(`^the file "([^"]+)" should exist in s3$`, TheFileExistsInS3(testCtx))
	ctx.Step(`^the file "([^"]+)" in s3 contains the content of "([^"]+)"$`, TheFileInS3ContainsTheContent(testCtx))

	// General
	ctx.Step(`^the headers:$`, TheHeaders(testCtx))
	ctx.Step(`^I delete the headers:$`, IDeleteTheHeaders(testCtx))
	ctx.Step(`^the template values:$`, TheTemplateValues(testCtx))
	ctx.Step(`^I store from the response for templating:$`, IStoreFromTheResponseForTemplating(testCtx))
	ctx.Step(`^the data should match JSON "([^"]*)"$`, TheDataShouldMatchJSON(testCtx))
	ctx.Step(`^I wait for (\d+) second(s)?$`, IWaitForSeconds(testCtx))
}
