module github.com/k2io/go-agent/v3/integrations/nrsnowflake

// snowflakedb/gosnowflake says it requires 1.12 but builds on 1.10
go 1.10

require (
	// v3.3.0 includes the new location of ParseQuery
	github.com/k2io/go-k2secure/v2 v2.0.0-20220929060728-f434fdbff641
	github.com/snowflakedb/gosnowflake v1.3.4
)
