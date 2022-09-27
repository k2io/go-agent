// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrnats_test

import (
	"fmt"
	"time"

	"github.com/k2io/go-agent/v3/integrations/nrnats"
	"github.com/k2io/go-agent/v3/newrelic"
	"github.com/nats-io/nats.go"
)

func currentTransaction() *newrelic.Transaction { return nil }

func ExampleStartPublishSegment() {
	nc, _ := nats.Connect(nats.DefaultURL)
	txn := currentTransaction()
	subject := "testing.subject"

	// Start the Publish segment
	seg := nrnats.StartPublishSegment(txn, nc, subject)
	err := nc.Publish(subject, []byte("Hello World"))
	if nil != err {
		panic(err)
	}
	// Manually end the segment
	seg.End()
}

func ExampleStartPublishSegment_defer() {
	nc, _ := nats.Connect(nats.DefaultURL)
	txn := currentTransaction()
	subject := "testing.subject"

	// Start the Publish segment and defer End till the func returns
	defer nrnats.StartPublishSegment(txn, nc, subject).End()
	m, err := nc.Request(subject, []byte("request"), time.Second)
	if nil != err {
		panic(err)
	}
	fmt.Println("Received reply message:", string(m.Data))
}
