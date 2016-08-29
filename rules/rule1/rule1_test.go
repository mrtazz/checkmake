package rule1

import (
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	th "github.com/mrtazz/checkmake/testhelpers"
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

	ret := ruleset(makefile, rules.RuleConfig{})

	th.Expect(t, len(ret), 0)

}

func TestMissingOnePhonyTarget(t *testing.T) {

	makefile := parser.Makefile{
		Variables: []parser.Variable{parser.Variable{
			Name:       "PHONY",
			Assignment: "all"}},
		Rules: []parser.Rule{parser.Rule{
			Target: "all"}, parser.Rule{Target: "clean"},
		}}

	ret := ruleset(makefile, rules.RuleConfig{})

	th.Expect(t, len(ret), 1)

}
