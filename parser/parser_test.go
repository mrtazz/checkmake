package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSimpleMakefile(t *testing.T) {

	ret, err := Parse("../fixtures/simple.make")

	assert.Equal(t, err, nil)
	assert.Equal(t, ret.FileName, "../fixtures/simple.make")
	assert.Equal(t, len(ret.Rules), 5)
	assert.Equal(t, len(ret.Variables), 4)
	assert.Equal(t, ret.Rules[0].Target, "clean")
	assert.Equal(t, ret.Rules[0].Body, []string{"rm bar", "rm foo"})

	assert.Equal(t, ret.Rules[1].Target, "foo")
	assert.Equal(t, ret.Rules[1].Body, []string{"touch foo"})
	assert.Equal(t, ret.Rules[1].Dependencies, []string{"bar"})

	assert.Equal(t, ret.Rules[2].Target, "bar")
	assert.Equal(t, ret.Rules[2].Body, []string{"touch bar"})

	assert.Equal(t, ret.Rules[3].Target, "all")
	assert.Equal(t, ret.Rules[3].Dependencies, []string{"foo"})

	assert.Equal(t, ret.Variables[0].Name, "expanded")
	assert.Equal(t, ret.Variables[0].Assignment, "\"$(simple)\"")
	assert.Equal(t, ret.Variables[0].SimplyExpanded, false)
	assert.Equal(t, ret.Variables[0].SpecialVariable, false)

	assert.Equal(t, ret.Variables[1].Name, "simple")
	assert.Equal(t, ret.Variables[1].Assignment, "\"foo\"")
	assert.Equal(t, ret.Variables[1].SimplyExpanded, true)
	assert.Equal(t, ret.Variables[1].SpecialVariable, false)

	assert.Equal(t, ret.Variables[2].Name, "PHONY")
	assert.Equal(t, ret.Variables[2].Assignment, "all clean test")
	assert.Equal(t, ret.Variables[2].SimplyExpanded, false)
	assert.Equal(t, ret.Variables[2].SpecialVariable, true)

	assert.Equal(t, ret.Variables[3].Name, "DEFAULT_GOAL")
	assert.Equal(t, ret.Variables[3].Assignment, "all")
	assert.Equal(t, ret.Variables[3].SimplyExpanded, false)
	assert.Equal(t, ret.Variables[3].SpecialVariable, true)
}
