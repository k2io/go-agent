package nrsecureagent

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	secAgent "github.com/k2io/go-k2secure"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	"gopkg.in/yaml.v2"
)

type SecurityConfig struct {
	secAgent.SecurityAgentConfig
	Error error
}

func InitSecurityAgent(app *newrelic.Application, opts ...ConfigOption) error {
	c := SecurityConfig{}
	for _, fn := range opts {
		if nil != fn {
			fn(&c)
			if nil != c.Error {
				return c.Error
			}
		}
	}
	c.Security.Policy.Version = "overridden"
	secureAgent := secAgent.InitSecurityAgent(c.Security, app.Config().AppName, app.Config().License)
	newrelic.InitSecurityAgent(secureAgent)
	return nil
}

//config options

type ConfigOption func(*SecurityConfig)

func ConfigSecurityFromYaml() ConfigOption {
	return func(cfg *SecurityConfig) {
		confgFilePath := os.Getenv("NEW_RELIC_SECURITY_CONFIG_PATH")
		if confgFilePath == "" {
			cfg.Error = fmt.Errorf("Invalid value: NEW_RELIC_SECURITY_CONFIG_PATH can't be empty")
			return
		}
		data, err := ioutil.ReadFile(confgFilePath)
		if err == nil {
			err = yaml.Unmarshal(data, &cfg.Security)
			if err != nil {
				cfg.Error = fmt.Errorf("Error while unmarshal config file %s value: %s", confgFilePath, err)
				return
			}
		} else {
			cfg.Error = fmt.Errorf("Error while reading config file %s , %s", confgFilePath, err)
			return
		}
	}
}

// NEW_RELIC_SECURITY_ENABLE
// NEW_RELIC_SECURITY_LOG_LEVEL
// NEW_RELIC_SECURITY_MODE
// NEW_RELIC_SECURITY_SECURITY_HOME_PATH
// NEW_RELIC_SECURITY_VALIDATOR_SERCICE_END_POINT_URL
// NEW_RELIC_SECURITY_Detection_DisableRxss

func ConfigSecurityFromEnvironment() ConfigOption {
	return func(cfg *SecurityConfig) {

		assignBool := func(field *bool, name string) {
			if env := os.Getenv(name); env != "" {
				if b, err := strconv.ParseBool(env); nil != err {
					cfg.Error = fmt.Errorf("invalid %s value: %s", name, env)
				} else {
					*field = b
				}
			}
		}
		assignString := func(field *string, name string) {
			if env := os.Getenv(name); env != "" {
				*field = env
			}
		}

		assignBool(&cfg.Security.Enable, "NEW_RELIC_SECURITY_ENABLE")
		assignString(&cfg.Security.SecurityHomePath, "NEW_RELIC_SECURITY_SECURITY_HOME_PATH")
		assignString(&cfg.Security.LogLevel, "NEW_RELIC_SECURITY_LOG_LEVEL")
		assignString(&cfg.Security.ValidatorServiceEndpointUrl, "NEW_RELIC_SECURITY_VALIDATOR_SERCICE_END_POINT_URL")
		assignString(&cfg.Security.Mode, "NEW_RELIC_SECURITY_MODE")
		assignBool(&cfg.Security.ForceCompleteDisable, "NEW_RELIC_SECURITY_FORCE_COMPLETE_DISABLE")
		assignBool(&cfg.Security.Detection.DisableRxss, "NEW_RELIC_SECURITY_DETECTION_DISABLE_RXSS")
	}
}
