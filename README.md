# go-bdd-common

[![Release](https://img.shields.io/github/release-pre/libsv/go-bdd-common.svg?logo=github&style=flat&v=1)](https://github.com/libsv/go-bdd-common/releases)
[![Go](https://github.com/libsv/go-bt/actions/workflows/run-tests.yml/badge.svg?branch=master)](https://github.com/libsv/go-bt/actions/workflows/run-tests.yml)
[![Go](https://github.com/libsv/go-bt/actions/workflows/run-tests.yml/badge.svg?event=pull_request)](https://github.com/libsv/go-bt/actions/workflows/run-tests.yml)

[![Report](https://goreportcard.com/badge/github.com/libsv/go-bdd-common?style=flat&v=1)](https://goreportcard.com/report/github.com/libsv/go-bdd-common)
[![Go](https://img.shields.io/github/go-mod/go-version/libsv/go-bdd-common?v=1)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-libsv-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/libsv)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat&v=3)](https://gobitcoinsv.com/#sponsor)
[![Mergify Status][mergify-status]][mergify]

[mergify]: https://mergify.io
[mergify-status]: https://img.shields.io/endpoint.svg?url=https://gh.mergify.io/badges/libsv/go-bdd-common&style=flat
<br/>

go-bdd-common is a generic [bdd](https://en.wikipedia.org/wiki/Behavior-driven_development) tool helping to write scalable tests for microservices applications. The framework is implemented in golang using [godog](https://github.com/cucumber/godog).

It runs within Docker Compose allowing you to setup and a tear down a deterministic known state environment within which to run your integration tests.

It supports 3 transports at the moment:

* Http
* Grpc
* Kafka

These allow you to send and receive data to services and test the responses, headers and response codes they may return.

It uses templated json files for request data and can validate responses against json files.

## To use

In order to use the library, add it to your project:

```shell
  go get github.com/libsv/go-bdd-common
```

We recommend adding to the service you are testing to keep everything together and adding a `testing` folder at the root with an `integration` sub folder. This 
will allow you to keep all your tests in one place and keep the integration tests separate.

### Folder structure

Setup the below folder structure

* testing
  * integration
    * data (for json request files)
    * features (for .feature files)
    * responses (for json responses)
    * env.yml (see below, allows you to setup real / mocked services)
    * integration_test.go (test entry point)
    
### integration_test.go

This is the main entry point to the tests and is picked up when running go test, note the build flag at the top.

This ensures the tests are only ran when you explicitly want to run them, otherwise, when running unit tests the integration suite would also be spun up.

```golang
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

	exitCode = cgodog.NewSuite(cgodog.Options{
		ServiceName:        "service-name", // your service
		ServicePort:        ":1234", // your port
		InitBitcoinBackend: true, // do we want bitcoin node running
		InitS3:             true, // do we want to setup minio
		InitKafka:          true, // initialise kafka
	}).Run()
}
```

## Environments

You can choose to run real or mocked services when running the integration tests.

NOTE: these services must exist in the docker compose file you are using, if you supply a service that isn't present, it will fail.

At the root of your integration test folder, add an `env.yml` file.

An example is shown:

```yaml
real:
  - s3
  - kafka
mocks:
  - this-service
  - that-service
vars:
  - MY_ENV_VAR=true
```

Note the 3 sections:

 * real - actual services to run, these will be setup as docker containers
 * mocks - mockec services, we can setup predicates for these to return canned responses
 * vars - custom environment variables

### Real Services 

These services must exist in the docker compose file being used.

Being in this section means the tests will setup a container for this service and run the actual code.

Useful if you *NEED* to run another service your service depends on like kafka, sql etc.

### Mocking Services

As much as possible, you want to exercise the service being tested only, this is where you can then mock downstream dependencies.

For example, given the below dependency:

  my-service -> other-service

We want to test my-service but don't want to test other-service as this is a test for my-service.

By mocking other-service you ensure that a failure in other-service won't impact my-service.

BUT by mocking, you can FORCE failures to check how my-service works. For example, for a request you could send a header or body data, setup the mock to return a 500 and test your servic ereacts properly.

#### Example mock predicate

Add the mock .json files to the /mock folder.

We recommend a file per mocked service, but you can split these as required.

The server we used is [based off grpc-mock](https://github.com/Tigh-Gherr/grpc-mock).

Some examples can be seen by following the above link.

NOTE: only grpc mocking is supported right now.

### Environment Variables

These are environment variables and they are passed ONLY to the service being tested.

## Custom Steps

The pre-canned steps are documented below.

You can of course add your own specific step definitions as required.

These are hooked into the `cgodog.NewSuite(` call in `integration_tests.go`

Our general steps should cover most scenarios, but of course there are edge cases and specifics that don't make sense here but you need.

### Init Scenario

The hook is an initScenerio function, the signature is shown below and is ran before each scenario is ran:

```go
func initScenario(sCtx *godog.ScenarioContext, tCtx *cgodog.Context) 
```

You can then add to the initSuite call as shown:

```go
  cgodog.NewSuite(cgodog.Options{
      ServiceName:         "my-service",
      ServicePort:         ":1234",
      ScenarioInitializer: initScenario, // add function hook
  }).Run())
  
  func initScenario(sCtx *godog.ScenarioContext, tCtx *cgodog.Context) {
    // can also add some pre-scenario code that is ran before every scenario is ran
  
    ctx.Step(`^I make a RANDOM request to "([^"]*)"$`, IMakeARANDOMRequestTo(testCtx))
    
  }

  // IMakeARANDOMRequestTo 
  func IMakeARANDOMRequestTo(ctx *Context) func(string) error {
    return func(proto string) error {
      // implementation of step
    }
  }

```

You can then add your new function to the feature files in the service.

## Steps

There are a number of pre-canned steps ready to use, these should cover most basic tests required when
testing microservices.

### GRPC

| Step                                                  | Description                                                                                  | Example                                                                                                |
|-------------------------------------------------------|----------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| I make a GRPC request to ""                           | Sends a GRPC request, with an empty body                                                     | I make a GRPC request to "proto.MyService/GetThing"                                                    |
| I make a GRPC request to "" with JSON ""              | Sends a GRPC request, with a json body, read from data/{filename}.json                       | I make a GRPC request to "proto.MyService/AddThing" with JSON "the_request"                            |
| I make a GRPC request to "" on port 1234 with JSON "" | Sends a GRPC request, with a json body, read from data/{filename}.json to a specific port    | I make a GRPC request to "proto.MyService/AddThing" on port 1234 with JSON "the_request"               |
| the GRPC code should be xx                            | Ran after one of the above steps, will check the GRPC response code matches the one supplied | the GRPC code should be OK - codes are listed https://grpc.github.io/grpc/core/md_doc_statuscodes.html |

### Http

| Step                                                        | Description                                                                                                       | Example                                                                 |
|-------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| I make a (GET POST PATCH DELETE) request to ""              | Sends an HTTP request with the supplied method to the endpoint defined.                                           | I make a GET request to "/api/v1/endpoint"                              |
| I make a (GET POST PATCH DELETE) request to "" with JSON "" | Sends an HTTP request with the supplied method to the endpoint defined with a request body defined in a JSON file | I make a POST request to "/api/v1/endpoint" with JSON "my_request"      |
| the HTTP response code should be 000                        | Added after one of the above steps, will validate the http response code matches what we expect                   | the HTTP response code should be 201 - one of https://httpstatuses.com/ |

### Bitcoin

| Step                                                      | Description                                                                   | Example                              |
|-----------------------------------------------------------|-------------------------------------------------------------------------------|--------------------------------------|
| there (are is) ([0-9]+) bitcoin account(s)?$              | Will setup a bitcoin node with n accounts and assign them to ctx.BTC.Accounts | there are 4 bitcoin accounts         |
| I send ([0-9.]+) BTC to account ([0-9]+)$                 | Sends the supplied about of bitcoin to the account with index                 | I send 0.001 BTC to account 1        |
| account ([0-9]+) sends ([0-9.]+) BTC to account ([0-9]+)$ | Send n number of bitcoin to account index from account                        | account 0 sends 0.1 BTC to account 1 |
| I send BTC to accounts:                                   | A table that defines account aliases and amount to top them up with           |                                      |

#### BTC To Accounts

```gherkin
And I send BTC to accounts:
| alias        | amount     |
| unused1      | 2          |
| unused2      | 3          |
```

### Transaction

| Step                                                      | Description                                                                   | Example                              |
|-----------------------------------------------------------|-------------------------------------------------------------------------------|--------------------------------------|
| there (are is) ([0-9]+) bitcoin account(s)?$              | Will setup a bitcoin node with n accounts and assign them to ctx.BTC.Accounts | there are 4 bitcoin accounts         |
| I send ([0-9.]+) BTC to account ([0-9]+)$                 | Sends the supplied about of bitcoin to the account with index                 | I send 0.001 BTC to account 1        |
| account ([0-9]+) sends ([0-9.]+) BTC to account ([0-9]+)$ | Send n number of bitcoin to account index from account                        | account 0 sends 0.1 BTC to account 1 |
| I send BTC to accounts:                                   | A table that defines account aliases and amount to top them up with           |                                      |

### Database

Database files should be stored under the /sql folder with the suffix .sql. The suffix does not need to be added to the name in the below steps

| Step                                 | Description                                                                                 | Example                                     |
|--------------------------------------|---------------------------------------------------------------------------------------------|---------------------------------------------|
| the database is seeded with "([^"]*) | Run the named sql file against the database, useful for seeding a database with known state | Give the database is seeded with "bad_data" |

### S3 / minio

Any files content checks will look for a folder named s3 in the root of the Integration test folder. Unlike json and sql files, suffix is required here as there
could be different file content types we are checking.

| Step                                                       | Description                                                                                                                           | Example                                                               |
|------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|
| I delete from s3:                                          | Remove a list of file names from S3                                                                                                   | Give I delete from s3:  (full example below)                          |
| the file "([^"]+)" should exist in s3                      | Checks that a file with the key provided exists, this can be of the form of a file path                                               | And the file "test/thing" should exist in s3                          |
| the file "([^"]+)" in s3 contains the content of "([^"]+)" | Checks that a file at path has content matching the file provided. These files should be added to an s3 folder and include the suffix | And the file "test/thing" in s3 contains the content of "post_tx.txt" |

### Kafka

| Step                                                          | Description                                                                                                                                                                                  | Example                                                                |
|---------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| I send a Kafka message to topic "([^"]*)" with JSON "([^"]*)" | Publishes the supplied JSON content to the kafka topic provided. Messages sent are given a unique identifier allowing the value to be checked by the below step as the identifier propagates | When I send a Kafka message to topic "test.topic" with JSON "my_data"  |
| I listen to topic "([^"]*)" and wait (\d+) ms for our message | Sets up a listener on the provided topic and can be given a configurable wait.                                                                                                               | And I listen to topic "test.response" and wait 5000 ms for our message |

### General

| Step                                      | Description                                                                                                | Example                                                                                                              |
|-------------------------------------------|------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| the headers:                              | Adds key value pairs to all outgoing requests following                                                    | Given the headers: (example below)                                                                                   |
| I delete the headers:                     | A list of header names to remove                                                                           | I delete the headers:                                                                                                |
| the template values:                      | Key value pairs of values to write to use in request or responses                                          | Given the template values:                                                                                           |
| I store from the response for templating: | Used for responses, can supply a JSON path and a name to give it, this can then be referenced in responses | And I store from the response for templating:                                                                        |
| the data should match JSON "([^"]*)"      | Will check a response (from http, grpc or kafka) matches the json at the file named.                       | Then I listen to topic "test.accepted" and wait 2000 ms for our message And the data should match JSON "my_expected" |
| I wait for (\d+) second(s)                | Pauses the current scenario for n seconds                                                                  | And I wait for 1 second                                                                                              |

#### the headers:
Example for the headers:

```gherkin
    And the headers:
      | key           | val |
      | header1 | abc |
      | header2 | integration-tx-broadcast |
```

#### I delete the headers:

Example

```gherkin
And I delete the headers:
      | key     |
      | header1 |
      | header2 |
```

#### the template values:

Example

This will write the value 100 to the variable named balance as shown:

```gherkin
    Given I fund 100 satoshis
    And the template values:
      | key      | val |
      | balance  | 100 |
```

This variable can then be referenced in data or response json files like this:

```json
{
  "myprop": "test",
  "balance": "{{ .balance }}"
  
}
```
If you need to check the same response template over and over with different values, this is a convenient way of doing it

#### I store from the response for templating:

```gherkin
    And I store from the response for templating:
      | jsonpath | name |
      | body.TxBytes | transaction |
```

This stores a value in the json response names TxBytes to a new variable called transaction.

This can then be referenced in json files as shown:

```json
{
  "myprop": "test",
  "tx": "{{ .transaction }}"
  
}
```