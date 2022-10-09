module github.com/k2io/go-agent/v3/integrations/nrmongo

// As of Dec 2019, 1.10 is the mongo-driver requirement:
// https://github.com/mongodb/mongo-go-driver#requirements
go 1.17

require (

	// mongo-driver does not support modules as of Nov 2019.
	go.mongodb.org/mongo-driver v1.10.2
	github.com/k2io/go-k2secure/v2 v2.0.0-20221009062520-14e980febfb4 // indirect
	github.com/k2io/go-k2secure/k2secure/k2secure_mongowrap/v2 v2.0.0-20221009062520-14e980febfb4 // indirect
)
