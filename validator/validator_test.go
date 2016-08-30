package validator

import (
	"testing"

	"github.com/mrtazz/checkmake/parser"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	violations := Validate(parser.Makefile{}, Config{})
	assert.Equal(t, 3, len(violations))
}
