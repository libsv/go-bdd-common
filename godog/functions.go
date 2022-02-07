package cgodog

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/libsv/go-bk/crypto"
	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/minio/minio-go/v7"
	"github.com/oliveagle/jsonpath"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/toorop/go-bitcoind"
	"github.com/yazgazan/jaydiff/diff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
)

func IMakeAGRPCRequestTo(ctx *Context) func(string) error {
	return func(proto string) error {
		ctx.GRPC = NewGRPCContext(proto, grpcClientConn)
		return makeGRPCRequest(ctx, "")
	}
}

func IMakeAGRPCRequestToWithData(ctx *Context) func(string, string) error {
	return func(proto, fileName string) error {
		ctx.GRPC = NewGRPCContext(proto, grpcClientConn)
		return makeGRPCRequest(ctx, fileName)
	}
}

func IMakeAGRPCRequestToOnPortWithData(ctx *Context) func(string, int, string) error {
	return func(proto string, port int, fileName string) error {
		conn, err := grpc.Dial(
			fmt.Sprintf(":%d", port),
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:    10 * time.Second,
				Timeout: 10 * time.Second,
			}),
		)
		if err != nil {
			return err
		}

		ctx.GRPC = NewGRPCContext(proto, conn)
		return makeGRPCRequest(ctx, fileName)
	}
}

func IMakeAHTTPRequestTo(ctx *Context, port string) func(string, string) error {
	return func(method, endpoint string) error {
		ctx.HTTP = NewHTTPContext(method, endpoint)

		tmpl, err := ctx.Template.Parse(endpoint)
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
			return err
		}

		req, err := http.NewRequest(method, fmt.Sprintf("http://0.0.0.0%s%s", port, out.String()), nil) // nolint:noctx
		if err != nil {
			return err
		}

		for k, v := range ctx.Headers {
			req.Header.Add(k, v)
		}

		resp, err := ctx.HTTP.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close() //nolint:errcheck

		if ctx.ResponseData, err = ioutil.ReadAll(resp.Body); err != nil {
			return err
		}

		ctx.HTTP.ResponseCode = resp.StatusCode

		return nil
	}
}

func IMakeAHTTPRequestToWithData(ctx *Context, port string) func(string, string, string) error {
	return func(method, endpoint, fileName string) error {
		ctx.HTTP = NewHTTPContext(method, endpoint)

		b, err := readFile(ctx, "data", fileName+".json")
		if err != nil {
			return err
		}

		ctx.RequestData = b

		req, err := http.NewRequest(method, fmt.Sprintf("http://0.0.0.0%s%s", port, endpoint), bytes.NewReader(ctx.RequestData)) //nolint:noctx
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/json")
		for k, v := range ctx.Headers {
			req.Header.Add(k, v)
		}

		resp, err := ctx.HTTP.client.Do(req)
		if err != nil {
			return errors.Wrap(err, "error performing request")
		}
		defer resp.Body.Close() //nolint:errcheck

		if ctx.ResponseData, err = ioutil.ReadAll(resp.Body); err != nil {
			if !errors.Is(err, io.EOF) {
				return errors.Wrap(err, "error decoding response")
			}
		}

		ctx.HTTP.ResponseCode = resp.StatusCode

		return nil
	}
}

func IUseTheAliases(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		conn, err := grpc.Dial(
			":9008",
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:    10 * time.Second,
				Timeout: 10 * time.Second,
			}),
		)
		if err != nil {
			return err
		}

		for _, row := range table.Rows[1:] {
			alias := row.Cells[0].Value
			if err := createAlias(ctx, conn, alias); err != nil {
				return err
			}
		}
		return nil
	}
}

func IUseRandomAliases(ctx *Context) func(int) error {
	return func(n int) error {
		conn, err := grpc.Dial(
			":9008",
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:    10 * time.Second,
				Timeout: 10 * time.Second,
			}),
		)
		if err != nil {
			return err
		}

		for i := 0; i < n; i++ {
			alias := uuid.NewString()
			if err := createAlias(ctx, conn, alias); err != nil {
				return err
			}
		}
		return nil
	}
}

func IFundSatoshis(ctx *Context) func(int) error {
	return func(amount int) error {
		if err := ISendBtcToWallet(ctx)(0.1, 0); err != nil {
			return err
		}

		if err := IStoreTheOfTXForTemplating(ctx)("TxID", len(ctx.BTC.CompletedTransactions)-1, "txid"); err != nil {
			return err
		}

		ctx.TemplateVars["amount"] = amount

		if err := IMakeAGRPCRequestToWithData(ctx)("proto.FundingService/AddFund", "add_fund"); err != nil {
			return err
		}

		if err := TheGRPCCodeShouldBe(ctx)("OK"); err != nil {
			return err
		}

		if err := TheDataShouldMatchJSON(ctx)("add_fund"); err != nil {
			return err
		}

		return nil
	}
}

func IFundAndSplitBySatoshis(ctx *Context) func(int) error {
	return func(denomination int) error {
		if err := ISendBtcToWallet(ctx)(0.5, 0); err != nil {
			return err
		}

		if err := IStoreTheOfTXForTemplating(ctx)("TxID", len(ctx.BTC.CompletedTransactions)-1, "txid"); err != nil {
			return err
		}

		if err := IStoreTheHexOfTXAsForTemplating(ctx)(len(ctx.BTC.CompletedTransactions)-1, "tx_hex"); err != nil {
			return err
		}

		ctx.TemplateVars["denomination"] = denomination

		if err := IMakeAGRPCRequestToWithData(ctx)("proto.FundingService/AddAndSplitFund", "add_and_split_fund_by"); err != nil {
			return err
		}

		if err := TheGRPCCodeShouldBe(ctx)("OK"); err != nil {
			return err
		}

		return nil
	}
}

func createAlias(ctx *Context, conn *grpc.ClientConn, alias string) error {
	var err error

	tmpl, err := ctx.Template.Parse(alias)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
		return err
	}

	alias = out.String()

	ctx.RequestData = []byte(`{"alias":"` + alias + `"}`)
	ctx.GRPC = NewGRPCContext("proto.CryptoService/CreateAlias", conn)
	if err = makeGRPCRequest(ctx, ""); err != nil {
		return err
	}

	fn := TheGRPCCodeShouldBe(ctx)
	if err = fn("OK"); err != nil {
		if err = fn("ALREADY_EXISTS"); err != nil {
			return err
		}
	}

	ctx.RequestData = []byte("{\"key_context\":{\"alias\":\"" + alias + "\",\"path\":\"0\"}}")
	ctx.GRPC = NewGRPCContext("proto.CryptoService/GetPublicKey", conn)
	if err = makeGRPCRequest(ctx, ""); err != nil {
		return err
	}

	createAliasResp := make(map[string]interface{})
	if err = json.NewDecoder(bytes.NewBuffer(ctx.ResponseData)).Decode(&createAliasResp); err != nil {
		return err
	}
	pk, err := base64.StdEncoding.DecodeString(createAliasResp["public_key"].(string))
	if err != nil {
		return err
	}

	s, err := bscript.NewP2PKHFromPubKeyBytes(pk)
	if err != nil {
		return err
	}

	ctx.BTC.Accounts = append(ctx.BTC.Accounts, BitcoinAccount{
		Alias:         alias,
		LockingScript: *s,
		Address:       createAliasResp["address"].(string),
	})

	return nil
}

func TheHeaders(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		ctx.Headers = make(map[string]string)
		for _, row := range table.Rows[1:] {
			val := row.Cells[1].Value
			tmpl, err := ctx.Template.Parse(val)
			if err != nil {
				return err
			}

			var out bytes.Buffer
			if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
				return err
			}
			ctx.Headers[strings.ToLower(row.Cells[0].Value)] = out.String()
		}
		return nil
	}
}

func IDeleteTheHeaders(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		for _, row := range table.Rows[1:] {
			raw := row.Cells[0].Value
			tmpl, err := ctx.Template.Parse(raw)
			if err != nil {
				return err
			}

			var out bytes.Buffer
			if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
				return err
			}

			delete(ctx.Headers, out.String())
		}
		return nil
	}
}

func TheTemplateValues(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		for _, row := range table.Rows[1:] {
			ctx.TemplateVars[row.Cells[0].Value] = row.Cells[1].Value
		}
		return nil
	}
}

func TheGRPCCodeShouldBe(ctx *Context) func(string) error {
	return func(respCode string) error {
		var exp codes.Code
		if err := exp.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, respCode))); err != nil {
			return err
		}

		if exp != ctx.GRPC.Status.Code() {
			return fmt.Errorf("expected %s, got %s: %s", exp, ctx.GRPC.Status.Code(), string(ctx.ResponseData))
		}
		return nil
	}
}

func TheHTTPResponseCodeShouldBe(ctx *Context) func(int) error {
	return func(respCode int) error {
		if respCode != ctx.HTTP.ResponseCode {
			return fmt.Errorf("expected %d, got %d", respCode, ctx.HTTP.ResponseCode)
		}

		return nil
	}
}

func TheDataShouldMatchJSON(ctx *Context) func(string) error {
	return func(fileName string) error {
		var actual, exp interface{}

		b, err := readFile(ctx, "responses", fileName+".json")
		if err != nil {
			return errors.Wrapf(err, "failed to read file %s.json", fileName)
		}
		if err := json.Unmarshal(b, &exp); err != nil {
			return errors.Wrap(err, "failed to unmarhsal expected data")
		}

		if len(ctx.ResponseData) == 0 {
			ctx.ResponseData = []byte("{}")
		}
		if err := json.Unmarshal(ctx.ResponseData, &actual); err != nil {
			return errors.Wrap(err, "failed to unmarshal response data")
		}

		return checkDiff(actual, exp)
	}
}

func IStoreFromTheResponseForTemplating(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		for _, v := range table.Rows[1:] {
			key := v.Cells[0].Value
			var jsonData interface{}
			if err := json.Unmarshal(ctx.ResponseData, &jsonData); err != nil {
				return err
			}
			val, err := jsonpath.JsonPathLookup(jsonData, "$"+key)
			if err != nil {
				return errors.Wrapf(err, "couldn't find key %s in response %s", key, string(ctx.ResponseData))
			}

			ctx.TemplateVars[v.Cells[1].Value] = val
		}
		return nil
	}
}

func ThereAreBitcoinWallets(ctx *Context) func(string, int) error {
	return func(_ string, n int) error {
		for i := 0; i < n; i++ {
			addr, err := ctx.BTC.client.GetNewAddress()
			if err != nil {
				return err
			}

			privkey, err := ctx.BTC.client.DumpPrivKey(addr)
			if err != nil {
				return err
			}

			wif, err := wif.DecodeWIF(privkey)
			if err != nil {
				return err
			}

			ctx.BTC.Accounts = append(ctx.BTC.Accounts, BitcoinAccount{
				Address: addr,
				WIF:     wif,
			})
		}
		return nil
	}
}

func ISendBtcToAccounts(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		bsvMap := make(map[string]float64)
		accountMap := make(map[string]BitcoinAccount)
		for _, account := range ctx.BTC.Accounts {
			if account.Alias != "" {
				accountMap[account.Alias] = account
			}
			accountMap[account.Address] = account
		}

		for _, row := range table.Rows[1:] {
			v, err := strconv.ParseFloat(row.Cells[1].Value, 64)
			if err != nil {
				return err
			}
			account := accountMap[row.Cells[0].Value]
			bsvMap[account.Address] = v
		}
		txID, err := ctx.BTC.client.SendMany("", bsvMap, 0, "")
		if err != nil {
			return err
		}

		rawTx, err := ctx.BTC.client.GetRawTransaction(txID, false)
		if err != nil {
			return err
		}

		tx, err := bt.NewTxFromString(rawTx.(string))
		if err != nil {
			return err
		}

		ctx.BTC.CompletedTransactions = append(ctx.BTC.CompletedTransactions, tx)
		return nil
	}
}

func ISendBtcToWallet(ctx *Context) func(float64, int) error {
	return func(amount float64, to int) error {
		txID, err := ctx.BTC.client.SendToAddress(
			ctx.BTC.Accounts[to].Address, amount, "", "",
		)
		if err != nil {
			return err
		}

		rawTx, err := ctx.BTC.client.GetRawTransaction(txID, false)
		if err != nil {
			return err
		}

		tx, err := bt.NewTxFromString(rawTx.(string))
		if err != nil {
			return err
		}

		ctx.BTC.CompletedTransactions = append(ctx.BTC.CompletedTransactions, tx)
		return nil
	}
}

func WalletSendsBtcToWallet(ctx *Context) func(int, float64, int) error {
	return func(from int, amount float64, to int) error {
		txID, err := ctx.BTC.client.SendFrom(
			ctx.BTC.Accounts[from].Address,
			ctx.BTC.Accounts[to].Address,
			amount,
			0,
			"",
			"",
		)
		if err != nil {
			return err
		}

		rawTx, err := ctx.BTC.client.GetRawTransaction(txID, false)
		if err != nil {
			return err
		}

		tx, err := bt.NewTxFromString(rawTx.(bitcoind.RawTransaction).Hex)
		if err != nil {
			return err
		}

		ctx.BTC.CompletedTransactions = append(ctx.BTC.CompletedTransactions, tx)
		return nil
	}
}

func IStoreTheOfWalletForTemplating(ctx *Context) func(string, int, string) error {
	return func(f string, i int, k string) error {
		acc := ctx.BTC.Accounts[i]
		ctx.TemplateVars[k] = reflect.ValueOf(acc).FieldByName(f).Interface()
		return nil
	}
}

func IStoreTheHexOfTXAsForTemplating(ctx *Context) func(int, string) error {
	return func(i int, k string) error {
		ctx.TemplateVars[k] = ctx.BTC.CompletedTransactions[i].String()
		return nil
	}
}

func IStoreTheOfTXForTemplating(ctx *Context) func(string, int, string) error {
	return func(f string, i int, k string) error {
		trans := ctx.BTC.CompletedTransactions[i]
		fn := reflect.ValueOf(trans).MethodByName(f)
		ctx.TemplateVars[k] = fn.Call([]reflect.Value{})[0].Interface()
		return nil
	}
}

func IStoreTheOfTheWorkingTXForTemplating(ctx *Context) func(string, string) error {
	return func(f string, k string) error {
		trans := ctx.BTC.WorkingTX.TX
		fn := reflect.ValueOf(trans).MethodByName(f)
		ctx.TemplateVars[k] = fn.Call([]reflect.Value{})[0].Interface()
		return nil
	}
}

func IUseTheOutputsOfTransactionAsInputs(ctx *Context) func(int) error {
	return func(n int) error {
		baseTx := ctx.BTC.CompletedTransactions[n]
		fromPKHash160 := crypto.Hash160(ctx.BTC.WorkingTX.From.WIF.SerialisePubKey())

		var totalSatoshis uint64

		for i, utxo := range baseTx.Outputs {
			b, err := utxo.LockingScript.PublicKeyHash()
			if err != nil {
				return err
			}

			if !bytes.Equal(b, fromPKHash160) {
				continue
			}

			if err := ctx.BTC.WorkingTX.TX.From(baseTx.TxID(),
				uint32(i), utxo.LockingScript.String(), utxo.Satoshis); err != nil {
				return err
			}

			totalSatoshis = totalSatoshis + utxo.Satoshis
		}

		txAmount := totalSatoshis - uint64(10000)
		if err := ctx.BTC.WorkingTX.TX.PayToAddress(ctx.BTC.WorkingTX.To.Address, txAmount); err != nil {
			return err
		}

		return nil
	}
}

func TheTransactionIsSignedByAccount(ctx *Context) func(int) error {
	return func(n int) error {
		if err := ctx.BTC.WorkingTX.TX.UnlockAll(context.Background(), &bt.LocalUnlockerGetter{
			PrivateKey: ctx.BTC.Accounts[n].WIF.PrivKey,
		}); err != nil {
			return err
		}
		ctx.TemplateVars["transaction"] = ctx.BTC.WorkingTX.TX.Bytes()
		return nil
	}
}

func IPrepareATranscationFromWalletToWallet(ctx *Context) func(int, int) error {
	return func(from, to int) error {
		ctx.BTC.WorkingTX = WorkingTransaction{
			From: ctx.BTC.Accounts[from],
			To:   ctx.BTC.Accounts[to],
			TX:   bt.NewTx(),
		}
		return nil
	}
}

func IPrepareATransaction(ctx *Context) func() error {
	return func() error {
		ctx.BTC.WorkingTX = WorkingTransaction{
			TX: bt.NewTx(),
		}
		return nil
	}
}

func IAddRandomDataToTheTransaction(ctx *Context) func() error {
	return func() error {
		var data [][]byte
		for i := 0; i < 10; i++ {
			data = append(data, []byte(uuid.NewString()))
		}
		return ctx.BTC.WorkingTX.TX.AddOpReturnPartsOutput(data)
	}
}

func TheDatabaseIsSeededWith(ctx *Context) func(string) error {
	return func(fileName string) error {
		b, err := readFile(ctx, "sql", fileName+".sql")
		if err != nil {
			return err
		}

		_, err = dbClientConn.Exec(string(b))
		return err
	}
}

func ACleanAccountDirectoryInS3(ctx *Context) func() error {
	return func() error {
		path := path.Join("accounts", ctx.ScenarioID)

		ch := ctx.S3.client.ListObjects(context.Background(), "ps-bucket", minio.ListObjectsOptions{
			Prefix:    path,
			Recursive: true,
		})

		errChan := ctx.S3.client.RemoveObjects(context.Background(), "ps-bucket", ch, minio.RemoveObjectsOptions{})

		for err := range errChan {
			if err.Err != nil {
				return err.Err
			}
		}

		return nil
	}
}

func IDeleteFromS3(ctx *Context) func(*godog.Table) error {
	return func(table *godog.Table) error {
		for _, row := range table.Rows[1:] {
			file := row.Cells[0].Value
			tmpl, err := ctx.Template.Parse(file)
			if err != nil {
				return err
			}

			var out bytes.Buffer
			if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
				return err
			}

			ch := ctx.S3.client.ListObjects(context.Background(), "ps-bucket", minio.ListObjectsOptions{
				Prefix:    out.String(),
				Recursive: true,
			})

			errChan := ctx.S3.client.RemoveObjects(
				context.Background(),
				"ps-bucket",
				ch,
				minio.RemoveObjectsOptions{},
			)

			for err := range errChan {
				if err.Err != nil {
					return err.Err
				}
			}
		}

		return nil
	}
}

func TheFileExistsInS3(ctx *Context) func(string) error {
	return func(file string) error {
		tmpl, err := ctx.Template.Parse(file)
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
			return err
		}

		_, err = ctx.S3.client.StatObject(
			context.Background(),
			"ps-bucket",
			out.String(),
			minio.StatObjectOptions{},
		)
		if err != nil {
			return err
		}

		return nil
	}
}

func TheFileInS3ContainsTheContent(ctx *Context) func(string, string) error {
	return func(s3Path string, contentPath string) error {
		tmpl, err := ctx.Template.Parse(s3Path)
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err = tmpl.Execute(&out, ctx.TemplateVars); err != nil {
			return err
		}

		obj, err := ctx.S3.client.GetObject(context.Background(), "ps-bucket", out.String(), minio.GetObjectOptions{})
		if err != nil {
			return err
		}

		stat, err := obj.Stat()
		if err != nil {
			return err
		}

		var content []byte = make([]byte, int(stat.Size))
		if _, err = obj.Read(content); err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}
		}

		exp, err := readFile(ctx, "s3", contentPath)
		if err != nil {
			return err
		}

		// Strip ending newline
		if exp[len(exp)-1] == 10 {
			exp = exp[:len(exp)-1]
		}

		return checkDiff(content, exp)
	}
}

func IWaitForSeconds(ctx *Context) func(int) error {
	return func(n int) error {
		time.Sleep(time.Duration(n) * time.Second)
		return nil
	}
}

func ISendAKafkaMessage(ctx *Context) func(string, string) error {
	return func(topic string, fileName string) error {
		if fileName != "" {
			b, err := readFile(ctx, "data", fileName+".json")
			if err != nil {
				return err
			}
			ctx.RequestData = b
		}
		var msg *Message
		if err := json.Unmarshal(ctx.RequestData, &msg); err != nil {
			return errors.Wrap(err, "failed to unmarshall message")
		}
		ctx.Kafka.ContextID = uuid.NewString()
		for key, val := range ctx.Headers {
			msg.Identifiers[key] = val
		}
		msg.Identifiers["Ps-Kensai_id"] = ctx.Kafka.ContextID
		bb, err := json.Marshal(msg)
		if err != nil {
			return errors.Wrap(err, "failed to marshall message")
		}
		if err := ctx.Kafka.Publisher.WriteMessages(context.Background(), kafka.Message{
			Topic: topic,
			Key:   []byte(ctx.Kafka.ContextID),
			Value: bb,
			Time:  time.Now().UTC(),
		}); err != nil {
			var errs kafka.WriteErrors
			if errors.As(err, &errs) {
				sb := strings.Builder{}
				for _, e := range errs {
					sb.WriteString(e.Error())
				}
				return fmt.Errorf("failed to send message to kafka: %s", sb.String())
			}
			return errors.Wrap(err, "failed to send message to kafka")
		}
		return nil
	}
}

func IReadAKafkaMessage(ctx *Context) func(string, int) error {
	return func(topic string, waitMs int) error {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{ctx.Kafka.Broker},
			GroupID:        uuid.NewString(),
			Topic:          topic,
			MaxWait:        2 * time.Second,
			StartOffset:    0,
			IsolationLevel: kafka.ReadCommitted,
		})
		defer reader.Close()
		c := context.Background()
		newCtx, cancelfn := context.WithTimeout(c, time.Millisecond*time.Duration(waitMs))
		defer cancelfn()

		found := false
		for !found {
			km, err := reader.ReadMessage(newCtx)
			if err != nil {
				return err
			}
			var msg *Message
			if err := json.Unmarshal(km.Value, &msg); err != nil {
				return errors.Wrap(err, "failed to unmarshall message when reading kafka")
			}
			if msg.Identifiers["Ps-Kensai_id"] == ctx.Kafka.ContextID {
				found = true
				delete(msg.Identifiers, "Ps-Kensai_id")
				bb, err := json.Marshal(msg)
				if err != nil {
					return errors.Wrap(err, "failed to marshall message when reading kafka")
				}
				ctx.ResponseData = bb
				return nil
			}
		}
		return fmt.Errorf("message not received for topic %s", topic)
	}
}

func readFile(ctx *Context, dir, fileName string) ([]byte, error) {
	if dir == "" || fileName == "" {
		return []byte{}, nil
	}

	b, err := ioutil.ReadFile(filepath.Clean(path.Join(dir, fileName)))
	if err != nil {
		return nil, err
	}

	tmpl, err := ctx.Template.Parse(string(b))
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, ctx.TemplateVars); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func makeGRPCRequest(ctx *Context, fileName string) error {
	ctx.ResponseData = make([]byte, 0)
	if fileName != "" {
		b, err := readFile(ctx, "data", fileName+".json")
		if err != nil {
			return err
		}
		ctx.RequestData = b
	}

	f, err := grpcreflect.NewClient(
		context.Background(),
		reflectpb.NewServerReflectionClient(ctx.GRPC.Conn),
	).FileContainingSymbol(ctx.GRPC.Service)
	if err != nil {
		return err
	}

	md := make(metadata.MD)
	for k, v := range ctx.Headers {
		md[k] = []string{v}
	}

	sd := f.FindSymbol(ctx.GRPC.Service).(*desc.ServiceDescriptor)
	mthd := sd.FindMethodByName(ctx.GRPC.Method)

	ext := dynamic.NewExtensionRegistryWithDefaults()
	ext.AddExtensionsFromFileRecursively(f)

	messageFactory := dynamic.NewMessageFactoryWithExtensionRegistry(ext)
	req := messageFactory.NewMessage(mthd.GetInputType())
	resp := messageFactory.NewMessage(mthd.GetOutputType())

	var rawJSON json.RawMessage
	if err = json.NewDecoder(bytes.NewBuffer(ctx.RequestData)).Decode(&rawJSON); err != nil {
		return err
	}
	if err = jsonpb.Unmarshal(bytes.NewBuffer(rawJSON), req); err != nil {
		return err
	}

	var respHeaders, respTrailers metadata.MD
	err = ctx.GRPC.Conn.Invoke(metadata.NewOutgoingContext(context.Background(), md),
		ctx.GRPC.Proto,
		req,
		resp,
		grpc.Trailer(&respTrailers),
		grpc.Header(&respHeaders),
	)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return err
		}
		ctx.GRPC.Status = st
	} else {
		ctx.GRPC.Status = status.New(codes.OK, "")
	}

	if ctx.GRPC.Status.Code() != codes.OK {
		errorResponse := map[string]interface{}{
			"error": ctx.GRPC.Status.Message(),
		}
		if ctx.ResponseData, err = json.Marshal(errorResponse); err != nil {
			return err
		}
		return nil
	}

	m := jsonpb.Marshaler{
		OrigName: true,
	}
	respStr, err := m.MarshalToString(resp)
	if err != nil {
		return err
	}

	if ctx.ResponseData, err = ioutil.ReadAll(bytes.NewBufferString(respStr)); err != nil {
		return err
	}

	return nil
}

func checkDiff(l, r interface{}) error {
	d, err := diff.Diff(l, r)
	if err != nil {
		return err
	}

	if d.Diff() == diff.Identical {
		return nil
	}

	report := d.StringIndent("", "", diff.Output{
		Indent:     "  ",
		ShowTypes:  true,
		JSON:       true,
		JSONValues: true,
	})

	return fmt.Errorf(report)
}

// Message defines a standard kensei event.
type Message struct {
	Identifiers map[string]string `json:"identifiers"`
	Key         string            `json:"key"`
	Body        json.RawMessage   `json:"body"`
	Origin      string            `json:"origin"`
	PublishMeta *PublishMeta      `json:"publishMeta,omitempty"`
}

// PublishMeta contains know meta objects that we can expect in a Message,
// more can be added as we need them.
//
// They should be nullable and not present if they have a nil or empty value.
type PublishMeta struct {
	Partition string           `json:"partition,omitempty"`
	Failures  []MessageFailure `json:"failures,omitempty"`
	Attempts  *int             `json:"attempts,omitempty"`
}

// MessageFailure defines a consumer failure and is added to the message
// if it has failed to be handled correctly in an application.
// By using WithFailure you can pass the error and this will create and
// add the MessageFailure to the message.
type MessageFailure struct {
	FailureTime time.Time `json:"failureTime"`
	Reason      string    `json:"reason"`
	HandledBy   string    `json:"handledBy"`
}
