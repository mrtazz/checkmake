// Package main tests, empty to at least have it be included in the build
package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/docopt/docopt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {

	args, err := docopt.Parse(usage, []string{"../../fixtures/simple.make"}, true,
		fmt.Sprintf("%s %s built at %s by %s with %s",
			"checkmake", version, buildTime, builder, goversion), false)

	require.Equal(t, nil, err, "docopt parsing should work")

	formatter, violations := parseArgsAndGetFormatter(args)

	assert.NotNil(t, formatter)
	assert.Equal(t, 0, len(violations))
}
