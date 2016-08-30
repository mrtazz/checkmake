package validator

import (
	"github.com/mrtazz/checkmake/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator(t *testing.T) {

	violations := Validate(parser.Makefile{}, Config{})

	assert.Equal(t, 0, len(violations))

}
