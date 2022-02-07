package cgodog

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/minio/minio-go/v7"
	"github.com/segmentio/kafka-go"
	"github.com/toorop/go-bitcoind"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// NewContext created a new *Context
func NewContext() *Context {
	return &Context{
		Headers:      make(map[string]string),
		TemplateVars: make(map[string]interface{}),
		RequestData:  make([]byte, 0),
		BTC:          NewBitcoinContext(),
	}
}

// AttachTemplater attaches a template engine to the test ctx
func (c *Context) AttachTemplater(name string) {
	c.Template = template.New(
		fmt.Sprintf("%s_templater", name),
	).Funcs(template.FuncMap{
		"base64Encode":   base64Encode,
		"sub":            subtract,
		"workingTxID":    workingTxID(c),
		"workingTxHex":   workingTxHex(c),
		"workingTxBytes": workingTxBytes(c),
	})
}

// NewGRPCContext created a new *GRPCCon
func NewGRPCContext(proto string, gc *grpc.ClientConn) *GRPCContext {
	svc, mthd := splitProto(proto)
	return &GRPCContext{
		Service: svc,
		Method:  mthd,
		Proto:   proto,
		Conn:    gc,
	}
}

// NewHTTPContext created a new *HTTPContext
func NewHTTPContext(method, endpoint string) *HTTPContext {
	return &HTTPContext{
		Method:   method,
		Endpoint: endpoint,
		client:   &http.Client{},
	}
}

// NewBitcoinContext created a new *BitcoinContext
func NewBitcoinContext() *BitcoinContext {
	return &BitcoinContext{
		client:                btcClientConn,
		CompletedTransactions: make([]*bt.Tx, 0),
		Accounts:              make([]BitcoinAccount, 0),
	}
}

func NewS3Context(client *minio.Client) *S3Context {
	return &S3Context{
		client: client,
	}
}

func NewKafkaContext(host string) *KafkaContext {
	w := &kafka.Writer{
		Addr:         kafka.TCP(host),
		Balancer:     kafka.Murmur2Balancer{},
		MaxAttempts:  5,
		BatchSize:    100,
		WriteTimeout: 10 * time.Second,
		Async:        false,
		Logger:       nil,
		ErrorLogger:  nil,
		Transport:    nil,
	}
	return &KafkaContext{
		Publisher: w,
		Consumer:  nil,
		// ContextID is a uniqueID to identify a scenarios messages.
		ContextID: uuid.NewString(),
		Broker:    host,
	}
}

// Context holds all variables and response data for a given scenario
type Context struct {
	// ScenarioID the id of the current godog scenario
	ScenarioID string
	// TemplateVars a map of variables which are placed into every call to the templater
	TemplateVars map[string]interface{}
	// Template the templater for the scenario
	Template *template.Template

	// HTTP the *HTTPContext used for making HTTP calls
	HTTP *HTTPContext
	// GRPC the *GRPCContext used for making GRPC calls
	GRPC *GRPCContext
	// BTC the *BitcoinContext used for making BTC requests
	BTC *BitcoinContext
	// S3 the *S3Context use for making S3 requests
	S3 *S3Context
	// Kafka stores the context used to send and receive kafka messages.
	Kafka *KafkaContext

	// Headers a map of key value pairs which are used as headers for every HTTP/GRPC call
	Headers map[string]string
	// RequestData the raw data which will be used for the next HTTP/GRPC call
	RequestData []byte
	// ResponseData is the raw response data returned from the previous HTTP/GRPC call
	ResponseData []byte
}

// HTTPContext stores HTTP specific information for making a HTTP call
type HTTPContext struct {
	// Endpoint the endpoint to be hit
	Endpoint string
	// Method the HTTP method to use
	Method string
	// ResponseCode the response code of the HTTP request
	ResponseCode int

	client *http.Client
}

// GRPCContext stores the GRPC specific information for making a GRPC call
type GRPCContext struct {
	// Service the service name to be accessed
	Service string
	// Method the method name to be called
	Method string
	// Proto
	Proto string
	// Conn the GRPC connection
	Conn *grpc.ClientConn
	// Status the status of the response
	Status *status.Status
}

// KafkaContext stores information for use across a scenario.
type KafkaContext struct {
	Publisher *kafka.Writer
	Consumer  *kafka.Reader
	ContextID string
	Broker    string
}

// BitcoinContext stores bitcoin node information used when interfacing
type BitcoinContext struct {
	client *bitcoind.Bitcoind
	// Accounts the list of accounts created during the scenario. Any new accounts are
	// appended, so it is up to the user to know which account is stored at which index
	Accounts []BitcoinAccount
	// CompletedTransactions the list of transactions which have been completed. Any new
	// transactions are appeneded, so it is up to the user to know which tx is stored at
	// which index
	CompletedTransactions []*bt.Tx
	// WorkingTX the current working transaction. Any tx building steps write to this.
	// Any tx broadcasting steps send this.
	WorkingTX WorkingTransaction
}

// S3Context stores the minio information
type S3Context struct {
	client *minio.Client
}

// WorkingTransaction stores information needed for TX building
type WorkingTransaction struct {
	// From the sending BitcoinAcccount
	From BitcoinAccount
	// To the sending BitcoinAccount
	To BitcoinAccount
	// TX the raw *bt.TX
	TX *bt.Tx
}

// BitcoinAccount information regarding a Bitcoin account on the node
type BitcoinAccount struct {
	// Alias the internal alias of the account (if exists)
	Alias string
	// LockingScript the accounts locking script
	LockingScript []byte
	// WIF the accounts *bsvutil.WIF
	WIF *wif.WIF
	// Address the accounts address
	Address string
}

func splitProto(proto string) (string, string) {
	pos := strings.LastIndex(proto, "/")
	if pos > 0 {
		return proto[:pos], proto[pos+1:]
	}

	pos = strings.LastIndex(proto, ".")
	if pos > 0 {
		return proto[:pos], proto[pos+1:]
	}

	return "", ""
}
