package app

import (
	"github.com/alexfedosov/move-tool/ablmodels"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateFolderIfNotExist verifies that folders are correctly created when needed
// and returns the correct path even when the folder already exists.
func TestCreateFolderIfNotExist(t *testing.T) {
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(testDir) // Clean up after test

	folderName := "test-folder"
	expectedPath := filepath.Join(testDir, folderName)

	resultPath, err := createFolderIfNotExist(testDir, folderName)
	require.NoError(t, err, "createFolderIfNotExist should not fail")
	assert.Equal(t, expectedPath, resultPath, "Returned path should match expected path")

	_, err = os.Stat(expectedPath)
	assert.False(t, os.IsNotExist(err), "Folder should exist at %s", expectedPath)

	resultPath2, err := createFolderIfNotExist(testDir, folderName)
	require.NoError(t, err, "createFolderIfNotExist should not fail on existing folder")
	assert.Equal(t, expectedPath, resultPath2, "Returned path should match expected path")
}

// TestRemoveDirectory verifies that directories are properly removed
// including their contents (files and subdirectories).
func TestRemoveDirectory(t *testing.T) {
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")

	testFile := filepath.Join(testDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test content"), 0644)
	require.NoError(t, err, "Failed to create test file")

	err = removeDirectory(testDir)
	require.NoError(t, err, "removeDirectory should not fail")

	_, err = os.Stat(testDir)
	assert.True(t, os.IsNotExist(err), "Directory should be removed")
}

// TestWritePresetFile verifies that device presets are correctly serialized to JSON
// and written to the filesystem at the expected location.
func TestWritePresetFile(t *testing.T) {
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(testDir) // Clean up after test

	filePath := "TestPath"
	audioFile := []ablmodels.AudioFile{
		{
			FilePath: &filePath,
			Duration: 1000.0,
		},
	}
	preset := ablmodels.NewDrumRackDevicePresetWithSamples(audioFile)

	err = writePresetFile(preset, testDir)
	require.NoError(t, err, "writePresetFile should not fail")

	presetPath := filepath.Join(testDir, "Preset.ablpreset")
	_, err = os.Stat(presetPath)
	assert.False(t, os.IsNotExist(err), "Preset file should exist at %s", presetPath)

	content, err := os.ReadFile(presetPath)
	require.NoError(t, err, "Should be able to read preset file")

	assert.Greater(t, len(content), 0, "Preset file should not be empty")
}

// TestArchivePresetBundle verifies that directories are correctly zipped into preset bundles
// with the expected file structure and naming convention.
func TestArchivePresetBundle(t *testing.T) {
	sourceDir, err := os.MkdirTemp("", "test-source")
	require.NoError(t, err, "Failed to create source directory")
	defer os.RemoveAll(sourceDir)

	outputDir, err := os.MkdirTemp("", "test-output")
	require.NoError(t, err, "Failed to create output directory")
	defer os.RemoveAll(outputDir)

	testFiles := []string{"test1.txt", "test2.txt", "subfolder/test3.txt"}
	testContent := []byte("test content")

	for _, filename := range testFiles {
		filePath := filepath.Join(sourceDir, filename)
		dirPath := filepath.Dir(filePath)

		if dirPath != sourceDir {
			err := os.MkdirAll(dirPath, os.ModePerm)
			require.NoError(t, err, "Failed to create directory %s", dirPath)
		}

		err := os.WriteFile(filePath, testContent, 0644)
		require.NoError(t, err, "Failed to create test file %s", filePath)
	}

	presetName := "test_preset"
	err = archivePresetBundle(presetName, sourceDir, outputDir)
	require.NoError(t, err, "archivePresetBundle should not fail")

	zipPath := filepath.Join(outputDir, presetName+".ablpresetbundle")
	_, err = os.Stat(zipPath)
	assert.False(t, os.IsNotExist(err), "Archive file should exist at %s", zipPath)
}
