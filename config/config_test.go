// Package config testing
package config

import (
	"testing"

	"github.com/mrtazz/checkmake/rules"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleConfig(t *testing.T) {

	cfg, err := NewConfigFromFile("../fixtures/exampleConfig.ini")

	require.Equal(t, nil, err, "Parsing of the fixture config file should have worked.")

	ruleCfg := cfg.GetRuleConfig("phonydeclared")

	assert.Equal(t, "true", ruleCfg["disabled"])
	assert.Equal(t, "bla", ruleCfg["foo"])

}

func TestFailConfig(t *testing.T) {

	cfg, err := NewConfigFromFile("../fixtures/idontexist.ini")

	assert.NotEqual(t, nil, err)

	val := cfg.GetRuleConfig("bla")

	assert.Equal(t, rules.RuleConfig(nil), val)
}
