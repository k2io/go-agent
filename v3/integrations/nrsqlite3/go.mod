module github.com/k2io/go-agent/v3/integrations/nrsqlite3

// As of Dec 2019, 1.9 is the oldest version of Go tested by go-sqlite3:
// https://github.com/mattn/go-sqlite3/blob/master/.travis.yml
go 1.9

require (
	github.com/mattn/go-sqlite3 v1.0.0
	// v3.3.0 includes the new location of ParseQuery
	github.com/k2io/go-k2secure/v2 v2.0.0-20220929060728-f434fdbff641
)
