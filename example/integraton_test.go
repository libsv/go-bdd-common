// +build godog

package integration

import (
	"os"
	"testing"

	cgodog "github.com/libsv/go-bdd-common/godog"
)

func TestMain(t *testing.M) {
	exitCode := 1
	defer func() { os.Exit(exitCode) }()

	// this spins up a compose environment, this needs to be made configurable.
	exitCode = cgodog.NewSuite(cgodog.Options{
		ServiceName: "blocks",
		ServicePort: ":80",
	}).Run()
}
