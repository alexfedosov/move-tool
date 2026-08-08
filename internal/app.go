package internal

import (
	"fmt"
	"github.com/alexfedosov/move-tool/internal/ablmodels"
	"github.com/brianvoe/gofakeit/v7"
	"strings"
)

func sanitizePresetName(presetName string) string {
	var result strings.Builder

	for _, char := range presetName {
		if (char >= 'a' && char <= 'z') || char == '_' {
			result.WriteRune(char)
		} else {
			result.WriteRune('_')
		}
	}

	return result.String()
}

func SliceSampleIntoDrumRack(inputFilePath string, outputFolderPath string, numberOfSlices int, customPresetName string) (err error) {
	err = gofakeit.Seed(0)
	if err != nil {
		return err
	}

	workspace, err := preparePresetWorkspace(outputFolderPath, customPresetName)
	if err != nil {
		return err
	}
	samples, err := writeAudioFileSlices(inputFilePath, workspace.SamplesFolderPath, numberOfSlices, workspace.PresetName)
	if err != nil {
		return err
	}

	preset := ablmodels.NewDrumRackDevicePresetWithSamples(*samples)
	return writeDrumRackPresetBundle(preset, workspace, outputFolderPath)
}

type presetWorkspace struct {
	PresetName        string
	PresetFolderPath  string
	SamplesFolderPath string
}

func preparePresetWorkspace(outputFolderPath string, customPresetName string) (presetWorkspace, error) {
	presetName := drumRackPresetName(customPresetName)
	presetFolderPath, err := createFolderIfNotExist(outputFolderPath, presetName)
	if err != nil {
		return presetWorkspace{}, err
	}
	samplesFolderPath, err := createFolderIfNotExist(presetFolderPath, "Samples")
	if err != nil {
		return presetWorkspace{}, err
	}
	return presetWorkspace{
		PresetName:        presetName,
		PresetFolderPath:  presetFolderPath,
		SamplesFolderPath: samplesFolderPath,
	}, nil
}

func drumRackPresetName(customPresetName string) string {
	if customPresetName != "" {
		return customPresetName
	}
	generatedName := strings.ToLower(fmt.Sprintf("%s_%s", gofakeit.HipsterWord(), gofakeit.AdverbPlace()))
	return sanitizePresetName(generatedName)
}

func writeDrumRackPresetBundle(preset *ablmodels.DevicePreset, workspace presetWorkspace, outputFolderPath string) error {
	if err := writePresetFile(preset, workspace.PresetFolderPath); err != nil {
		return err
	}
	if err := archivePresetBundle(workspace.PresetName, workspace.PresetFolderPath, outputFolderPath); err != nil {
		return err
	}
	return removeDirectory(workspace.PresetFolderPath)
}
