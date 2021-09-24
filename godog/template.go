package cgodog

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

func base64Encode(i interface{}) (string, error) {
	s := fmt.Sprintf("%s", i)
	return base64.StdEncoding.EncodeToString([]byte(s)), nil
}

func subtract(a, b interface{}) (string, error) {
	l, err := strconv.ParseInt(fmt.Sprintf("%v", a), 10, 64)
	if err != nil {
		return "", err
	}

	r, err := strconv.ParseInt(fmt.Sprintf("%v", b), 10, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", l-r), nil
}

func workingTxHex(ctx *Context) func() (string, error) {
	return func() (string, error) {
		return ctx.BTC.WorkingTX.TX.String(), nil
	}
}

func workingTxBytes(ctx *Context) func() (string, error) {
	return func() (string, error) {
		return string(ctx.BTC.WorkingTX.TX.Bytes()), nil
	}
}
