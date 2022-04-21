// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package newrelic

import (
	"runtime"

	"github.com/k2io/go-agent/v3/internal"
	_ "github.com/k2io/go-k2secure/v2"
)

const (
	// Version is the full string version of this Go Agent.
	Version = "3.15.2"
)

var (
	goVersionSimple = minorVersion(runtime.Version())
)

func init() {
	internal.TrackUsage("Go", "Version", Version)
	internal.TrackUsage("Go", "Runtime", "Version", goVersionSimple)
	internal.TrackUsage("Go", "gRPC", "Version", grpcVersion)
}
