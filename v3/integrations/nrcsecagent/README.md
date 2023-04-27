# v3/integrations/nrcsecagent [![GoDoc](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrcsecagent?status.svg)](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrcsecagent)


## Import and add [nrcsecagent](https://github.com/newrelic/go-agent/v3/integrations/nrcsecagent) dependency in application
Add security agent

```
go get github.com/newrelic/go-agent/v3/integrations/nrcsecagent 
```
Then import the package in your application:

```
import "github.com/newrelic/go-agent/v3/integrations/nrcsecagent"
```

### Step 2.1: Create an Application

Instantiate your application by running the following:
```
app, err := newrelic.NewApplication(
    newrelic.ConfigAppName("Your Application Name"),
    newrelic.ConfigLicense("NEW_RELIC_LICENSE_KEY"),
    newrelic.ConfigDebugLogger(os.Stdout),
)
```

After instantiating your app, Init nrcsecagent as given below

```
    err := nrcsecagent.InitSecurityAgent(
        app,
       	nrcsecagent.ConfigSecurityMode("IAST"),
        nrcsecagent.ConfigSecurityValidatorServiceEndPointUrl("wss://csec.nr-data.net"),
        nrcsecagent.ConfigSecurityEnable(true),
    )
```

Generate traffic against your application for the IAST agent to detect vulnerabilities. Once vulnerabilities are detected they will be reported in the vulnerabilities list.

For more information, see
[godocs](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrcsecagent).
