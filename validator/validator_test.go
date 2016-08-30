package validator

import (
	"github.com/mrtazz/checkmake/config"
	"github.com/mrtazz/checkmake/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator(t *testing.T) {

	violations := Validate(parser.Makefile{}, &config.Config{})

	assert.Equal(t, 0, len(violations))

}
