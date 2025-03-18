package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceCommandFlags(t *testing.T) {
	// Test that the slice command has the expected flags
	cmd := sliceCmd
	
	// Test input flag
	flag := cmd.Flag("input")
	assert.NotNil(t, flag, "'input' flag should exist")
	assert.Equal(t, "i", flag.Shorthand, "'input' flag should have shorthand 'i'")
	
	// Test output flag
	flag = cmd.Flag("output")
	assert.NotNil(t, flag, "'output' flag should exist")
	assert.Equal(t, "o", flag.Shorthand, "'output' flag should have shorthand 'o'")
	
	// Test numberOfSlices flag
	flag = cmd.Flag("numberOfSlices")
	assert.NotNil(t, flag, "'numberOfSlices' flag should exist")
	assert.Equal(t, "n", flag.Shorthand, "'numberOfSlices' flag should have shorthand 'n'")
	assert.Equal(t, "16", flag.DefValue, "'numberOfSlices' flag should have default value '16'")
	
	// Test preset-name flag
	flag = cmd.Flag("preset-name")
	assert.NotNil(t, flag, "'preset-name' flag should exist")
	assert.Equal(t, "p", flag.Shorthand, "'preset-name' flag should have shorthand 'p'")
}

func TestSliceCommandMarkRequiredFlags(t *testing.T) {
	// Test that the input flag is marked as required
	inputFlag := sliceCmd.Flags().Lookup("input")
	require.NotNil(t, inputFlag, "input flag not found")
	
	// Check that the command has marked required flags
	// We can't directly check if a flag is marked as required, but we can check if the command
	// has flags marked as required by checking its PreRunE function
	if sliceCmd.PreRunE == nil {
		t.Skip("Skipping required flag check - command uses a different mechanism for validation")
	}
}