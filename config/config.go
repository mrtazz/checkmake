// Package config deals with loading and parsing configuration from disk
package config

import (
	"github.com/go-ini/ini"
	"github.com/mrtazz/checkmake/logger"
	"github.com/mrtazz/checkmake/rules"
)

// Config is a struct to configure the validator and rules
type Config struct {
	iniFile *ini.File
}

// NewConfigFromFile returns a config struct that is filled with the values
// from the passed in ini file
func NewConfigFromFile(path string) (*Config, error) {
	iniFile, err := ini.Load(path)
	ret := &Config{
		iniFile: iniFile,
	}

	return ret, err
}

// GetRuleConfig returns a rules.RuleConfig for the given rule. A rule
// corresponds to a section in the config ini file
func (c *Config) GetRuleConfig(rule string) (ret rules.RuleConfig) {

	if c.iniFile == nil {
		logger.Debug("iniFile not initialized")
		return
	}

	ret = make(rules.RuleConfig)

	section, err := c.iniFile.GetSection(rule)

	if err == nil {
		for _, keyName := range section.KeyStrings() {

			key, keyError := section.GetKey(keyName)
			if keyError == nil {
				ret[keyName] = key.String()
			}
		}
	}

	return
}
