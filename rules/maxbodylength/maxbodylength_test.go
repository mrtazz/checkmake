package maxbodylength

import (
	"testing"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/stretchr/testify/assert"
)

func TestFooIsTooLong(t *testing.T) {

	makefile := parser.Makefile{
		Rules: []parser.Rule{parser.Rule{
			Target: "foo",
			Body: []string{"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'"},
			LineNumber: 1}},
	}

	rule := MaxBodyLength{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, 1, len(ret))
	assert.Equal(t, "Target bodies should be kept simple and short.",
		rule.Description())
	assert.Equal(t, "Target body for \"foo\" exceeds allowed length of 5 (7).", ret[0].Violation)
	assert.Equal(t, 1, ret[0].LineNumber)
}

func TestFooIsTooLongWithConfig(t *testing.T) {

	makefile := parser.Makefile{
		Rules: []parser.Rule{parser.Rule{
			Target: "foo",
			Body: []string{"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'",
				"echo 'foo'"},
			LineNumber: 1}},
	}

	rule := MaxBodyLength{}

	cfg := rules.RuleConfig{}
	cfg["maxBodyLength"] = "3"

	ret := rule.Run(makefile, cfg)

	assert.Equal(t, 1, len(ret))
	assert.Equal(t, "Target bodies should be kept simple and short.",
		rule.Description())
	assert.Equal(t, "Target body for \"foo\" exceeds allowed length of 3 (4).", ret[0].Violation)
	assert.Equal(t, 1, ret[0].LineNumber)
}
