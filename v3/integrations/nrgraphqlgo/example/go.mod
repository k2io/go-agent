module github.com/k2io/go-agent/v3/integrations/nrgraphqlgo/example

go 1.13

require (
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/graphql-go-handler v0.2.3
	github.com/k2io/go-k2secure/v2 v2.0.0-20220929060728-f434fdbff641
	github.com/k2io/go-agent/v3/integrations/nrgraphqlgo v1.0.0
)

replace github.com/k2io/go-agent/v3 => ../../../

replace github.com/k2io/go-agent/v3/integrations/nrgraphqlgo => ../
