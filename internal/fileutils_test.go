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

func TestCreateFolderIfNotExist(t *testing.T) {
	testDir, err := os.MkdirTemp("", "test-move-tool")
	require.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(testDir)

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

func TestWritePresetFile(t *testing.T) {
	testDir := mustCreateTempDir(t, "test-move-tool")
	preset := testDrumRackPreset()

	err := writePresetFile(preset, testDir)
	require.NoError(t, err, "writePresetFile should not fail")

	assertPresetFileWritten(t, testDir)
}

func testDrumRackPreset() *ablmodels2.DevicePreset {
	filePath := "TestPath"
	audioFile := []ablmodels2.AudioFile{
		{
			FilePath: &filePath,
			Duration: 1000.0,
		},
	}
	return ablmodels2.NewDrumRackDevicePresetWithSamples(audioFile)
}

func assertPresetFileWritten(t *testing.T, testDir string) {
	t.Helper()

	presetPath := filepath.Join(testDir, "Preset.ablpreset")
	_, err := os.Stat(presetPath)
	assert.False(t, os.IsNotExist(err), "Preset file should exist at %s", presetPath)

	content, err := os.ReadFile(presetPath)
	require.NoError(t, err, "Should be able to read preset file")

	assert.Greater(t, len(content), 0, "Preset file should not be empty")
}

func TestWritePresetFileWithInvalidDirectory(t *testing.T) {
	nonExistentDir := "/path/to/nonexistent/directory"
	preset := testDrumRackPreset()

	err := writePresetFile(preset, nonExistentDir)

	assert.Error(t, err, "writePresetFile should fail with non-existent directory")
}

func TestArchivePresetBundle(t *testing.T) {
	fixture := newPresetBundleFixture(t)
	presetName := "test_preset"

	reader := fixture.archivePresetBundle(t, presetName)
	defer reader.Close()

	assertArchiveFiles(t, reader, fixture.TestFiles, fixture.TestContent)
}

type presetBundleFixture struct {
	SourceDir   string
	OutputDir   string
	TestFiles   []string
	TestContent []byte
}

func newPresetBundleFixture(t *testing.T) presetBundleFixture {
	t.Helper()

	sourceDir := mustCreateTempDir(t, "test-source")
	outputDir := mustCreateTempDir(t, "test-output")
	testFiles := []string{"test1.txt", "test2.txt", "subfolder/test3.txt"}
	testContent := []byte("test content")
	writeTestFiles(t, sourceDir, testFiles, testContent)

	return presetBundleFixture{
		SourceDir:   sourceDir,
		OutputDir:   outputDir,
		TestFiles:   testFiles,
		TestContent: testContent,
	}
}

func (f presetBundleFixture) archivePresetBundle(t *testing.T, presetName string) *zip.ReadCloser {
	t.Helper()

	err := archivePresetBundle(presetName, f.SourceDir, f.OutputDir)
	require.NoError(t, err, "archivePresetBundle should not fail")

	zipPath := filepath.Join(f.OutputDir, presetName+".ablpresetbundle")
	_, err = os.Stat(zipPath)
	assert.False(t, os.IsNotExist(err), "Archive file should exist at %s", zipPath)

	reader, err := zip.OpenReader(zipPath)
	require.NoError(t, err, "Failed to open zip file")
	return reader
}

func writeTestFiles(t *testing.T, sourceDir string, testFiles []string, testContent []byte) {
	t.Helper()

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
}

func assertArchiveFiles(t *testing.T, reader *zip.ReadCloser, testFiles []string, testContent []byte) {
	t.Helper()

	var foundFiles []string
	for _, file := range reader.File {
		foundFiles = append(foundFiles, file.Name)
		assertArchiveFileContent(t, file, testContent)
	}

	assertExpectedArchiveFiles(t, foundFiles, testFiles)
}

func assertArchiveFileContent(t *testing.T, file *zip.File, testContent []byte) {
	t.Helper()

	if strings.HasSuffix(file.Name, "/") {
		return
	}

	rc, err := file.Open()
	require.NoError(t, err, "Failed to open file in archive")
	defer rc.Close()

	content, err := io.ReadAll(rc)
	require.NoError(t, err, "Failed to read file content")

	assert.Equal(t, testContent, content, "File content should match for %s", file.Name)
}

func assertExpectedArchiveFiles(t *testing.T, foundFiles []string, testFiles []string) {
	t.Helper()

	for _, expectedFile := range testFiles {
		assert.Contains(t, foundFiles, expectedFile, "Archive should contain %s", expectedFile)
	}
}

func TestArchivePresetBundleWithInvalidOutputDir(t *testing.T) {
	sourceDir := sourceDirWithTestFile(t)

	err := archivePresetBundle("test_preset", sourceDir, "/path/to/nonexistent/directory")

	assert.Error(t, err, "archivePresetBundle should fail with non-existent output directory")
}

func sourceDirWithTestFile(t *testing.T) string {
	t.Helper()

	sourceDir := mustCreateTempDir(t, "test-source")
	testFilePath := filepath.Join(sourceDir, "test.txt")
	err := os.WriteFile(testFilePath, []byte("test content"), 0644)
	require.NoError(t, err, "Failed to create test file")

	return sourceDir
}
