package internal

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/alexfedosov/move-tool/internal/ablmodels"
	"io"
	"os"
	"path"
	"path/filepath"
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

	archiver := presetBundleArchiver{
		sourceDirectory: directory,
		zipWriter:       zipWriter,
	}
	return filepath.Walk(directory, archiver.addPath)
}

type presetBundleArchiver struct {
	sourceDirectory string
	zipWriter       *zip.Writer
}

func (a presetBundleArchiver) addPath(filePath string, info os.FileInfo, walkErr error) error {
	if walkErr != nil {
		return walkErr
	}
	if info.IsDir() {
		return nil
	}
	return a.copyFile(filePath)
}

func (a presetBundleArchiver) copyFile(filePath string) error {
	relativePath, err := filepath.Rel(a.sourceDirectory, filePath)
	if err != nil {
		return err
	}
	zipFileWriter, err := a.zipWriter.Create(relativePath)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(zipFileWriter, file)
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

func removeDirectory(path string) error {
	return os.RemoveAll(path)
}
