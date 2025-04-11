package internal

import (
        "archive/zip"
        ablmodels2 "github.com/alexfedosov/move-tool/internal/ablmodels"
        "io"
        "os"
        "path/filepath"
        "strings"
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
        audioFile := []ablmodels2.AudioFile{
                {
                        FilePath: &filePath,
                        Duration: 1000.0,
                },
        }
        preset := ablmodels2.NewDrumRackDevicePresetWithSamples(audioFile)

        err = writePresetFile(preset, testDir)
        require.NoError(t, err, "writePresetFile should not fail")

        presetPath := filepath.Join(testDir, "Preset.ablpreset")
        _, err = os.Stat(presetPath)
        assert.False(t, os.IsNotExist(err), "Preset file should exist at %s", presetPath)

        content, err := os.ReadFile(presetPath)
        require.NoError(t, err, "Should be able to read preset file")

        assert.Greater(t, len(content), 0, "Preset file should not be empty")
}

// TestWritePresetFileWithInvalidDirectory verifies that the function correctly handles
// errors when the output directory doesn't exist.
func TestWritePresetFileWithInvalidDirectory(t *testing.T) {
        // Use a non-existent directory
        nonExistentDir := "/path/to/nonexistent/directory"
        
        filePath := "TestPath"
        audioFile := []ablmodels2.AudioFile{
                {
                        FilePath: &filePath,
                        Duration: 1000.0,
                },
        }
        preset := ablmodels2.NewDrumRackDevicePresetWithSamples(audioFile)
        
        err := writePresetFile(preset, nonExistentDir)
        
        // Verify that the function returns an error
        assert.Error(t, err, "writePresetFile should fail with non-existent directory")
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
        
        // Verify the contents of the archive
        extractDir, err := os.MkdirTemp("", "test-extract")
        require.NoError(t, err, "Failed to create extraction directory")
        defer os.RemoveAll(extractDir)
        
        // Open the zip file
        reader, err := zip.OpenReader(zipPath)
        require.NoError(t, err, "Failed to open zip file")
        defer reader.Close()
        
        // Check that all expected files are in the archive
        var foundFiles []string
        for _, file := range reader.File {
                foundFiles = append(foundFiles, file.Name)
                
                // Extract and verify content of each file
                rc, err := file.Open()
                require.NoError(t, err, "Failed to open file in archive")
                
                content, err := io.ReadAll(rc)
                require.NoError(t, err, "Failed to read file content")
                rc.Close()
                
                // Verify content for non-directory entries
                if !strings.HasSuffix(file.Name, "/") {
                        assert.Equal(t, testContent, content, "File content should match for %s", file.Name)
                }
        }
        
        // Verify all expected files are in the archive
        for _, expectedFile := range testFiles {
                assert.Contains(t, foundFiles, expectedFile, "Archive should contain %s", expectedFile)
        }
}

// TestArchivePresetBundleWithInvalidOutputDir verifies that the function correctly handles
// errors when the output directory doesn't exist.
func TestArchivePresetBundleWithInvalidOutputDir(t *testing.T) {
        sourceDir, err := os.MkdirTemp("", "test-source")
        require.NoError(t, err, "Failed to create source directory")
        defer os.RemoveAll(sourceDir)
        
        // Create a test file in the source directory
        testFilePath := filepath.Join(sourceDir, "test.txt")
        err = os.WriteFile(testFilePath, []byte("test content"), 0644)
        require.NoError(t, err, "Failed to create test file")
        
        // Use a non-existent output directory
        nonExistentDir := "/path/to/nonexistent/directory"
        
        err = archivePresetBundle("test_preset", sourceDir, nonExistentDir)
        
        // Verify that the function returns an error
        assert.Error(t, err, "archivePresetBundle should fail with non-existent output directory")
}
