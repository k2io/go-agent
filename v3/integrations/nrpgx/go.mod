module github.com/k2io/go-agent/v3/integrations/nrpgx

// As of Dec 2019, go 1.11 is the earliest version of Go tested by lib/pq:
// https://github.com/lib/pq/blob/master/.travis.yml
go 1.11

require (
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/k2io/go-k2secure/v2 v2.0.0-20220929060728-f434fdbff641 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e // indirect
	golang.org/x/text v0.3.7 // indirect
)
