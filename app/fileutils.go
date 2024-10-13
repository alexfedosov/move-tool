package app

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"move-tool/ablmodels"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func createFolderIfNotExist(basePath string, folderName string) (path string, err error) {
	path = filepath.Join(basePath, folderName)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
	}
	return path, err
}

func archivePresetBundle(presetName string, directory string, output string) error {
	zipFilePath := path.Join(output, fmt.Sprintf("%s.ablpresetbundle", presetName))
	fmt.Printf("Creating preset bundle %s\n", zipFilePath)
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relativePath := strings.TrimPrefix(path, fmt.Sprintf("%s/", directory))
		if info.IsDir() {
			return nil
		}
		zipFileWriter, err := zipWriter.Create(relativePath)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(zipFileWriter, file)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func writePresetFile(preset *ablmodels.DevicePreset, presetFolderPath string) error {
	presentJSON, err := json.MarshalIndent(preset, "", "  ")
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/Preset.ablpreset", presetFolderPath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(presentJSON)
	if err != nil {
		return err
	}
	return nil
}
