// This sqlx example is a separate module to avoid adding sqlx dependency to the
// nrpgx go.mod file.

module github.com/k2io/go-agent/v3/integrations/nrpgx/example/sqlx

go 1.13

require (
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/k2io/go-agent/v3 v3.3.0
	github.com/k2io/go-agent/v3/integrations/nrpgx v0.0.0
)

replace github.com/k2io/go-agent/v3 => ../../../../

replace github.com/k2io/go-agent/v3/integrations/nrpgx => ../../
