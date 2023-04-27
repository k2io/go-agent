# v3/integrations/nrsecurityagent [![GoDoc](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrsecurityagent?status.svg)](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrsecurityagent)


## Import and add [nrsecurityagent](https://github.com/newrelic/go-agent/v3/integrations/nrsecurityagent) dependency in application
Add security agent

```
go get github.com/newrelic/go-agent/v3/integrations/nrsecurityagent 
```
Then import the package in your application:

```
import "github.com/newrelic/go-agent/v3/integrations/nrsecurityagent"
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

After instantiating your app, Init nrsecurityagent as given below

```
    err := nrsecurityagent.InitSecurityAgent(
        app,
       	nrsecurityagent.ConfigSecurityMode("IAST"),
        nrsecurityagent.ConfigSecurityValidatorServiceEndPointUrl("wss://csec.nr-data.net"),
        nrsecurityagent.ConfigSecurityEnable(true),
    )
```

Generate traffic against your application for the IAST agent to detect vulnerabilities. Once vulnerabilities are detected they will be reported in the vulnerabilities list.

For more information, see
[godocs](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrsecurityagent).
