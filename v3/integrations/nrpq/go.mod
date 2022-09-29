module github.com/k2io/go-agent/v3/integrations/nrpq

// As of Dec 2019, go 1.11 is the earliest version of Go tested by lib/pq:
// https://github.com/lib/pq/blob/master/.travis.yml
go 1.11

require (
	// NewConnector dsn parsing tests expect v1.1.0 error return behavior.
	github.com/lib/pq v1.1.0
	// v3.3.0 includes the new location of ParseQuery
	github.com/k2io/go-k2secure/v2 v2.0.0-20220929060728-f434fdbff641
	google.golang.org/grpc v1.27.0 // indirect
)

replace github.com/k2io/go-k2secure/v2 v2.0.0-20220929060728-f434fdbff641 => ../../../v3

replace github.com/k2io/go-agent/v3/newrelic => ../../../v3/newrelic
