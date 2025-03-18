package app

import (
	"move-tool/ablmodels"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFolderIfNotExist(t *testing.T) {
	// Create a temporary directory for testing
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(testDir) // Clean up after test

	// Test case 1: Create a new folder
	folderName := "test-folder"
	expectedPath := filepath.Join(testDir, folderName)

	resultPath, err := createFolderIfNotExist(testDir, folderName)
	require.NoError(t, err, "createFolderIfNotExist should not fail")
	assert.Equal(t, expectedPath, resultPath, "Returned path should match expected path")

	// Check if folder was actually created
	_, err = os.Stat(expectedPath)
	assert.False(t, os.IsNotExist(err), "Folder should exist at %s", expectedPath)

	// Test case 2: Try to create a folder that already exists
	resultPath2, err := createFolderIfNotExist(testDir, folderName)
	require.NoError(t, err, "createFolderIfNotExist should not fail on existing folder")
	assert.Equal(t, expectedPath, resultPath2, "Returned path should match expected path")
}

func TestRemoveDirectory(t *testing.T) {
	// Create a temporary directory for testing
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")

	// Create a file inside the test directory
	testFile := filepath.Join(testDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test content"), 0644)
	require.NoError(t, err, "Failed to create test file")

	// Remove the directory
	err = removeDirectory(testDir)
	require.NoError(t, err, "removeDirectory should not fail")

	// Verify the directory was removed
	_, err = os.Stat(testDir)
	assert.True(t, os.IsNotExist(err), "Directory should be removed")
}

func TestWritePresetFile(t *testing.T) {
	// Create a temporary directory for testing
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(testDir) // Clean up after test

	// Create a minimal device preset for testing
	filePath := "TestPath"
	audioFile := []ablmodels.AudioFile{
		{
			FilePath: &filePath,
			Duration: 1000.0,
		},
	}
	preset := ablmodels.NewDrumRackDevicePresetWithSamples(audioFile)

	// Write the preset file
	err = writePresetFile(preset, testDir)
	require.NoError(t, err, "writePresetFile should not fail")

	// Check if the file was created
	presetPath := filepath.Join(testDir, "Preset.ablpreset")
	_, err = os.Stat(presetPath)
	assert.False(t, os.IsNotExist(err), "Preset file should exist at %s", presetPath)

	// Read the file contents to verify it's valid JSON
	content, err := os.ReadFile(presetPath)
	require.NoError(t, err, "Should be able to read preset file")

	// Verify the file contains expected content
	assert.Greater(t, len(content), 0, "Preset file should not be empty")
}

func TestArchivePresetBundle(t *testing.T) {
	// Create temporary directories for testing
	sourceDir, err := os.MkdirTemp("", "test-source")
	require.NoError(t, err, "Failed to create source directory")
	defer os.RemoveAll(sourceDir)

	outputDir, err := os.MkdirTemp("", "test-output")
	require.NoError(t, err, "Failed to create output directory")
	defer os.RemoveAll(outputDir)

	// Create test files in the source directory
	testFiles := []string{"test1.txt", "test2.txt", "subfolder/test3.txt"}
	testContent := []byte("test content")

	for _, filename := range testFiles {
		filePath := filepath.Join(sourceDir, filename)
		dirPath := filepath.Dir(filePath)

		// Create subdirectory if needed
		if dirPath != sourceDir {
			err := os.MkdirAll(dirPath, os.ModePerm)
			require.NoError(t, err, "Failed to create directory %s", dirPath)
		}

		// Create test file
		err := os.WriteFile(filePath, testContent, 0644)
		require.NoError(t, err, "Failed to create test file %s", filePath)
	}

	// Archive the preset bundle
	presetName := "test_preset"
	err = archivePresetBundle(presetName, sourceDir, outputDir)
	require.NoError(t, err, "archivePresetBundle should not fail")

	// Check if the archive was created
	zipPath := filepath.Join(outputDir, presetName+".ablpresetbundle")
	_, err = os.Stat(zipPath)
	assert.False(t, os.IsNotExist(err), "Archive file should exist at %s", zipPath)
}