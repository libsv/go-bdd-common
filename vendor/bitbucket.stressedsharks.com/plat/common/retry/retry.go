// Package retry defines a Run method in which you can execute a function n number of times until it either
// successfully completes, or hits max attempts and returns an error.
//
// It supports RetryOpts in which you can override default behaviours by supplying them at run time as shown
//  Run(ctx, myFunc, WithBackOff(time.Minute * 2), WithAttempts(2), WithoutExponentialBackoff())
//
// To implement a RetryFunc is very simple, just create a function that either natches the signature or returns a RetryFunc:
//  setupFunc := func(ctx context.Context) error {
//		logger.Info("attempting to setup thing")
//		t, err := thing.Setup()
//		if err != nil {
//			logger.Error("failed to setup thing, retrying")
//			return err
//		}
//		logger.Info("thing setup successfully")
//		return nil
//	}
//
// This can then be passed to the Run func as shown,using default args and in a go routine so it isn't blocking:
//  go func() {
//	if err := Retry(context.Background(), setupFunc); err != nil {
//		logger.Error("failed to start thing")
//	}
//  }()
package retry

import (
	"context"
	"time"
)

// RetryerFunc defines a function that can be retried n number of times.
// The retryer will run the RetryFunc iteratively until it hits the max attempts or
// it completes without error.
type RetryerFunc func(ctx context.Context) error

// RetryerOptFunc can be implemented to override default retry options.
type RetryerOptFunc func(opt *retryOptions)

// WithBackOff can be supplied when running the retry function to change
// the default wait time of 10 seconds between retries.
func WithBackOff(backoff time.Duration) RetryerOptFunc {
	return func(opt *retryOptions) {
		opt.backoff = backoff
	}
}

// WithExponentialBackoff will run the retry function by
// exponentially increasing the backoff.
//
// If this method is not supplied the retryer will run in default mode where
// it waits for the backoff time to elapse before retrying.
// Given a backoff of 10 seconds, in exponential backoff mode, the first time will be 10 seconds,
// 2nd time will be 20 seconds, 3rd will be 60 seconds etc. It will run with
// backoff = backoff * attemptNumber.
func WithExponentialBackoff() RetryerOptFunc {
	return func(opt *retryOptions) {
		opt.exponentialBackoff = true
	}
}

// WithAttempts will override the default attempts of 3 with the
// value of attempts supplied.
func WithAttempts(attempts int) RetryerOptFunc {
	return func(opt *retryOptions) {
		opt.maxAttempts = attempts
	}
}

type retryOptions struct {
	backoff            time.Duration
	exponentialBackoff bool
	maxAttempts        int
}

// Run will execute the function fn with a default of 3 attempts using an exponential backoff
// with a default backoff time of 10 seconds between attempts. the 2nd and 3rd attempts will
// wait backoff time * attempt number.
// These defaults can be overridden by passing one or more RetryOptFuncs such as WithAttempts.
// If the result of the run has been a success, no error will be returned.
// If after n attempts the RetryFunc still returns an error, it will be returned.
// Its recommended to run this in a go routine, though not essential as it depends on use case.
func Run(ctx context.Context, fn RetryerFunc, opts ...RetryerOptFunc) error {
	retryOpts := &retryOptions{
		backoff:            time.Second * 10,
		maxAttempts:        3,
		exponentialBackoff: false,
	}
	// apply any override options
	for _, opt := range opts {
		opt(retryOpts)
	}
	var err error
	backoff := retryOpts.backoff
	for i := 0; i < retryOpts.maxAttempts; i++ {
		if e := fn(ctx); e != nil {
			err = e
			// exponential backoff
			if retryOpts.exponentialBackoff {
				backoff = backoff * time.Duration(i+1)
			}
			time.Sleep(backoff)
			continue
		}
		return nil
	}
	return err
}
