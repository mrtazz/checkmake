package rule1

import (
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllTargetsArePhony(t *testing.T) {

	makefile := parser.Makefile{
		Variables: []parser.Variable{parser.Variable{
			Name:       "PHONY",
			Assignment: "all clean"}},
		Rules: []parser.Rule{parser.Rule{
			Target: "all"}, parser.Rule{Target: "clean"},
		}}

	rule := Rule1{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, len(ret), 0)

}

func TestMissingOnePhonyTarget(t *testing.T) {

	makefile := parser.Makefile{
		Variables: []parser.Variable{parser.Variable{
			Name:       "PHONY",
			Assignment: "all"}},
		Rules: []parser.Rule{parser.Rule{
			Target: "all"}, parser.Rule{Target: "clean"},
		}}

	rule := Rule1{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, len(ret), 1)

}
