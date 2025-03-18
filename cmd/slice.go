package cmd

import (
	"github.com/spf13/cobra"
	"move-tool/app"
)

var (
	input          string
	output         string
	numberOfSlices int
	presetName     string

	sliceCmd = &cobra.Command{
		Use:   "slice",
		Short: "Slices long sample into drum rack",
		Long:  `Slice long sample into given number of equal numberOfSlices and creates a drum rack preset`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.SliceSampleIntoDrumRack(input, output, numberOfSlices, presetName)
		},
	}
)

func init() {
	sliceCmd.Flags().StringVarP(&input, "input", "i", "", "input file")
	_ = sliceCmd.MarkFlagRequired("input")
	sliceCmd.Flags().StringVarP(&output, "output", "o", "", "Output directory")
	_ = sliceCmd.MarkFlagRequired("output")
	sliceCmd.Flags().IntVarP(&numberOfSlices, "numberOfSlices", "n", 16, "Number of numberOfSlices")
	_ = sliceCmd.MarkFlagRequired("numberOfSlices")
	sliceCmd.Flags().StringVarP(&presetName, "preset-name", "p", "", "Custom name for the preset (without extension)")
	rootCmd.AddCommand(sliceCmd)
}
