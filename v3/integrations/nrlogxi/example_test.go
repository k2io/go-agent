// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrlogxi_test

import (
	nrlogxi "github.com/k2io/go-agent/v3/integrations/nrlogxi"
	newrelic "github.com/k2io/go-agent/v3/newrelic"
	log "github.com/mgutz/logxi/v1"
)

func Example() {
	// Create a new logxi logger:
	l := log.New("newrelic")
	l.SetLevel(log.LevelInfo)

	newrelic.NewApplication(
		newrelic.ConfigAppName("Example App"),
		newrelic.ConfigLicense("__YOUR_NEWRELIC_LICENSE_KEY__"),
		// Use nrlogxi to register the logger with the agent:
		nrlogxi.ConfigLogger(l),
	)
}
