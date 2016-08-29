package parser

import (
	th "github.com/mrtazz/checkmake/testhelpers"
	"testing"
)

func TestParseSimpleMakefile(t *testing.T) {

	ret, err := Parse("../fixtures/simple.make")

	th.Expect(t, err, nil)
	th.Expect(t, len(ret.Rules), 4)
	th.Expect(t, len(ret.Variables), 4)
	th.Expect(t, ret.Rules[0].Target, "clean")
	th.Expect(t, ret.Rules[0].Body, []string{"rm bar", "rm foo"})

	th.Expect(t, ret.Rules[1].Target, "foo")
	th.Expect(t, ret.Rules[1].Body, []string{"touch foo"})
	th.Expect(t, ret.Rules[1].Dependencies, []string{"bar"})

	th.Expect(t, ret.Rules[2].Target, "bar")
	th.Expect(t, ret.Rules[2].Body, []string{"touch bar"})

	th.Expect(t, ret.Rules[3].Target, "all")
	th.Expect(t, ret.Rules[3].Dependencies, []string{"foo"})

	th.Expect(t, ret.Variables[0].Name, "expanded")
	th.Expect(t, ret.Variables[0].Assignment, "\"$(simple)\"")
	th.Expect(t, ret.Variables[0].SimplyExpanded, false)
	th.Expect(t, ret.Variables[0].SpecialVariable, false)

	th.Expect(t, ret.Variables[1].Name, "simple")
	th.Expect(t, ret.Variables[1].Assignment, "\"foo\"")
	th.Expect(t, ret.Variables[1].SimplyExpanded, true)
	th.Expect(t, ret.Variables[1].SpecialVariable, false)

	th.Expect(t, ret.Variables[2].Name, "PHONY")
	th.Expect(t, ret.Variables[2].Assignment, "all clean")
	th.Expect(t, ret.Variables[2].SimplyExpanded, false)
	th.Expect(t, ret.Variables[2].SpecialVariable, true)

	th.Expect(t, ret.Variables[3].Name, "DEFAULT_GOAL")
	th.Expect(t, ret.Variables[3].Assignment, "all")
	th.Expect(t, ret.Variables[3].SimplyExpanded, false)
	th.Expect(t, ret.Variables[3].SpecialVariable, true)
}
