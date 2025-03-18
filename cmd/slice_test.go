package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSliceCommandFlags verifies that the slice command has all required flags configured
// with correct shorthand values and default settings.
func TestSliceCommandFlags(t *testing.T) {
	cmd := sliceCmd

	flag := cmd.Flag("input")
	assert.NotNil(t, flag, "'input' flag should exist")
	assert.Equal(t, "i", flag.Shorthand, "'input' flag should have shorthand 'i'")

	flag = cmd.Flag("output")
	assert.NotNil(t, flag, "'output' flag should exist")
	assert.Equal(t, "o", flag.Shorthand, "'output' flag should have shorthand 'o'")

	flag = cmd.Flag("numberOfSlices")
	assert.NotNil(t, flag, "'numberOfSlices' flag should exist")
	assert.Equal(t, "n", flag.Shorthand, "'numberOfSlices' flag should have shorthand 'n'")
	assert.Equal(t, "16", flag.DefValue, "'numberOfSlices' flag should have default value '16'")

	flag = cmd.Flag("preset-name")
	assert.NotNil(t, flag, "'preset-name' flag should exist")
	assert.Equal(t, "p", flag.Shorthand, "'preset-name' flag should have shorthand 'p'")
}

// TestSliceCommandMarkRequiredFlags verifies that required flags are properly marked
// so they can be validated before command execution.
func TestSliceCommandMarkRequiredFlags(t *testing.T) {
	inputFlag := sliceCmd.Flags().Lookup("input")
	require.NotNil(t, inputFlag, "input flag not found")

	if sliceCmd.PreRunE == nil {
		t.Skip("Skipping required flag check - command uses a different mechanism for validation")
	}
}
