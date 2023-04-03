// Copyright 2022 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrcsecagent

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	securityAgent "github.com/k2io/go-k2secure"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	"gopkg.in/yaml.v2"
)

type SecurityConfig struct {
	securityAgent.SecurityAgentConfig
	Error error
}

// defaultSecurityConfig creates a Security Config populated with default settings.
func defaultSecurityConfig() SecurityConfig {
	cfg := SecurityConfig{}
	cfg.Security.Enabled = false
	cfg.Security.Validator_service_url = "wss://csec.nr-data.net"
	cfg.Security.Mode = "IAST"
	cfg.Security.Agent.Enabled = true
	cfg.Security.Detection.Rxss.Enabled = true
	return cfg
}

// InitSecurityAgent initilized the nrcsecagent agent with provied config.
func InitSecurityAgent(app *newrelic.Application, opts ...ConfigOption) error {
	c := defaultSecurityConfig()
	for _, fn := range opts {
		if nil != fn {
			fn(&c)
			if nil != c.Error {
				return c.Error
			}
		}
	}
	secureAgent := securityAgent.InitSecurityAgent(c.Security, app.Config().AppName, app.Config().License, app.Config().Logger.DebugEnabled())
	newrelic.InitSecurityAgent(secureAgent)
	return nil
}

type ConfigOption func(*SecurityConfig)

//Getting config through config YAML file
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

/* getting config through env variables */

// ConfigSecurityFromEnvironment populates the config based on environment variables:

//		NEW_RELIC_SECURITY_ENABLED							sets Security.Enabled
//		NEW_RELIC_SECURITY_VALIDATOR_SERVICE_URL			sets Security.Validator_service_url
//		NEW_RELIC_SECURITY_MODE								sets Security.Mode
//		NEW_RELIC_SECURITY_AGENT_ENABLED					sets Security.Agent.Enabled
//		NEW_RELIC_SECURITY_DETECTION_RXSS_ENABLED			sets cfg.Security.Detection.Rxss.Enabled

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

		assignBool(&cfg.Security.Enabled, "NEW_RELIC_SECURITY_ENABLED")
		assignString(&cfg.Security.Validator_service_url, "NEW_RELIC_SECURITY_VALIDATOR_SERVICE_URL")
		assignString(&cfg.Security.Mode, "NEW_RELIC_SECURITY_MODE")
		assignBool(&cfg.Security.Agent.Enabled, "NEW_RELIC_SECURITY_AGENT_ENABLED")
		assignBool(&cfg.Security.Detection.Rxss.Enabled, "NEW_RELIC_SECURITY_DETECTION_RXSS_ENABLED")
	}
}

/* getting config through config methods */

// ConfigSecurityMode sets security mode default: IAST
func ConfigSecurityMode(mode string) ConfigOption {
	return func(cfg *SecurityConfig) {
		cfg.Security.Mode = mode
	}
}

// ConfigSecurityValidatorServiceEndPointUrl sets security validator service endpoint.
func ConfigSecurityValidatorServiceEndPointUrl(url string) ConfigOption {
	return func(cfg *SecurityConfig) {
		cfg.Security.Validator_service_url = url
	}
}

// ConfigSecurityDetectionDisableRxss used to disable rxss validation.
func ConfigSecurityDetectionDisableRxss(isEnabled bool) ConfigOption {
	return func(cfg *SecurityConfig) {
		cfg.Security.Detection.Rxss.Enabled = isEnabled
	}
}

// ConfigSecurityEnable sets Security.Enabled
func ConfigSecurityEnable(isEnabled bool) ConfigOption {
	return func(cfg *SecurityConfig) {
		cfg.Security.Enabled = isEnabled
	}
}
