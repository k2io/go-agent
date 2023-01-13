package newrelic

import (
	"net/http"
)

var SecureAgent SecurityAgent = NoOpAgent{}

type SecurityAgent interface {
	RefreshState(map[string]string) bool
	DeactivateSecurity()
	SendEvent(string, ...interface{}) interface{}
	IsSecurityActive() bool
	DistributedTraceHeaders(hdrs *http.Request, secureAgentevent interface{})
	SendExitEvent(interface{}, error)
}

func InitSecurityAgent(s SecurityAgent) {
	if s != nil {
		SecureAgent = s
	}
}

func getLinkedMetaData(app *app) map[string]string {
	runningAppData := map[string]string{}
	if app != nil && app.run != nil {
		runningAppData["hostname"] = app.run.Config.hostname
		runningAppData["entityName"] = app.run.firstAppName
		if app.run != nil {
			runningAppData["entityGUID"] = app.run.Reply.EntityGUID
			runningAppData["agentRunId"] = app.run.Reply.RunID.String()
			runningAppData["accountId"] = app.run.Reply.AccountID
		}
	}
	return runningAppData
}

// NoOpAgent
type NoOpAgent struct {
}

func (t NoOpAgent) RefreshState(connectionData map[string]string) bool {
	return false
}

func (t NoOpAgent) DeactivateSecurity() {

}

func (t NoOpAgent) SendEvent(caseType string, data ...interface{}) interface{} {
	return nil
}

func (t NoOpAgent) IsSecurityActive() bool {
	return false
}

func (t NoOpAgent) DistributedTraceHeaders(hdrs *http.Request, secureAgentevent interface{}) {
}

func (t NoOpAgent) SendExitEvent(secureAgentevent interface{}, err error) {

}
