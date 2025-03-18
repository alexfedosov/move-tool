package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	// Test that the root command exists
	assert.NotNil(t, rootCmd, "Root command should not be nil")

	// Test the root command has a use value
	assert.NotEmpty(t, rootCmd.Use, "Root command should have a Use value")

	// Test the root command has a short description
	assert.NotEmpty(t, rootCmd.Short, "Root command should have a Short description")

	// Test that the root command has the slice command as a subcommand
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "slice" {
			found = true
			break
		}
	}
	assert.True(t, found, "Root command should have 'slice' as a subcommand")
}
