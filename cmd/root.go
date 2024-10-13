package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "move-tool",
		Short: "move-tool helps you mangle your Ableton Move files.",
		Long:  `move-tool is a CLI for mangling your Ableton Move files.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
