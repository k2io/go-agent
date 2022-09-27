package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/k2io/go-agent/v3/integrations/logcontext-v2/nrzerolog"
	"github.com/k2io/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
)

func main() {
	baseLogger := zerolog.New(os.Stdout)

	app, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
		newrelic.ConfigAppName("NRZerolog Example"),
		newrelic.ConfigInfoLogger(os.Stdout),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app.WaitForConnection(5 * time.Second)

	nrHook := nrzerolog.NewRelicHook{
		App: app,
	}

	nrLogger := baseLogger.Hook(nrHook)
	nrLogger.Info().Msg("Hello World")

	// With transaction context
	txn := app.StartTransaction("My Transaction")
	ctx := newrelic.NewContext(context.Background(), txn)

	nrTxnHook := nrzerolog.NewRelicHook{
		App:     app,
		Context: ctx,
	}

	txnLogger := baseLogger.Hook(nrTxnHook)
	txnLogger.Debug().Msg("This is a transaction log")

	txn.End()

	nrLogger.Info().Msg("Goodbye")
	app.Shutdown(10 * time.Second)
}
