package app

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"move-tool/ablmodels"
	"strings"
	"regexp"
)

func SliceSampleIntoDrumRack(inputFilePath string, outputFolderPath string, numberOfSlices int) (err error) {
	err = gofakeit.Seed(0)
	if err != nil {
		return err
	}
	presetName := strings.ToLower(fmt.Sprintf("%s_%s", gofakeit.HipsterWord(), gofakeit.AdverbPlace()))
	
	re := regexp.MustCompile(`[^a-z]+`)
	presetName = re.ReplaceAllString(strings.ToLower(presetName), "_")
	
	presetFolderPath, err := createFolderIfNotExist(outputFolderPath, presetName)
	if err != nil {
		return err
	}
	samplesFolderPath, err := createFolderIfNotExist(presetFolderPath, "Samples")
	if err != nil {
		return err
	}
	samples, err := writeAudioFileSlices(inputFilePath, samplesFolderPath, numberOfSlices, presetName)
	if err != nil {
		return err
	}

	preset := ablmodels.NewDrumRackDevicePresetWithSamples(*samples)

	err = writePresetFile(preset, presetFolderPath)
	if err != nil {
		return err
	}
	err = archivePresetBundle(presetName, presetFolderPath, outputFolderPath)
	if err != nil {
		return err
	}

	err = removeDirectory(presetFolderPath)
	if err != nil {
		return err
	}

	return nil
}
