package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRootCommand verifies that the root cobra command is properly initialized
// with correct values and contains the expected subcommands.
func TestRootCommand(t *testing.T) {
	assert.NotNil(t, rootCmd, "Root command should not be nil")
	assert.NotEmpty(t, rootCmd.Use, "Root command should have a Use value")
	assert.NotEmpty(t, rootCmd.Short, "Root command should have a Short description")

	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "slice" {
			found = true
			break
		}
	}
	assert.True(t, found, "Root command should have 'slice' as a subcommand")
}
