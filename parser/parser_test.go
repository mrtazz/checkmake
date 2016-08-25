package parser

import (
	"reflect"
	"testing"
)

// Expect provides a simple way to write unit test assertions
// gratefully taken and adapted from
// http://keighl.com/post/mocking-http-responses-in-golang/
func Expect(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("Expected: %v (type %v) - Got: %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestParseSimpleMakefile(t *testing.T) {

	ret, err := Parse("../fixtures/simple.make")

	Expect(t, err, nil)
	Expect(t, len(ret.Rules), 4)
	Expect(t, len(ret.Variables), 4)
	Expect(t, ret.Rules[0].Target, "clean")
	Expect(t, ret.Rules[0].Body, []string{"rm bar", "rm foo"})

	Expect(t, ret.Rules[1].Target, "foo")
	Expect(t, ret.Rules[1].Body, []string{"touch foo"})
	Expect(t, ret.Rules[1].Dependencies, []string{"bar"})

	Expect(t, ret.Rules[2].Target, "bar")
	Expect(t, ret.Rules[2].Body, []string{"touch bar"})

	Expect(t, ret.Rules[3].Target, "all")
	Expect(t, ret.Rules[3].Dependencies, []string{"foo"})

	Expect(t, ret.Variables[0].Name, "expanded")
	Expect(t, ret.Variables[0].Assignment, "\"$(simple)\"")
	Expect(t, ret.Variables[0].SimplyExpanded, false)
	Expect(t, ret.Variables[0].SpecialVariable, false)

	Expect(t, ret.Variables[1].Name, "simple")
	Expect(t, ret.Variables[1].Assignment, "\"foo\"")
	Expect(t, ret.Variables[1].SimplyExpanded, true)
	Expect(t, ret.Variables[1].SpecialVariable, false)

	Expect(t, ret.Variables[2].Name, "PHONY")
	Expect(t, ret.Variables[2].Assignment, "all clean")
	Expect(t, ret.Variables[2].SimplyExpanded, false)
	Expect(t, ret.Variables[2].SpecialVariable, true)

	Expect(t, ret.Variables[3].Name, "DEFAULT_GOAL")
	Expect(t, ret.Variables[3].Assignment, "all")
	Expect(t, ret.Variables[3].SimplyExpanded, false)
	Expect(t, ret.Variables[3].SpecialVariable, true)
}
