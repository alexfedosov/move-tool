package internal

import (
	"fmt"
	"github.com/alexfedosov/move-tool/internal/ablmodels"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteAudioFileSlices(t *testing.T) {
	fixture := newAudioSlicesFixture(t)

	audioFiles, err := writeAudioFileSlices(fixture.InputFilePath, fixture.OutputDir, fixture.NumberOfSlices, fixture.FilenamePrefix)
	require.NoError(t, err, "writeAudioFileSlices should not fail")

	assertAudioSlices(t, fixture, audioFiles)
}

type audioSlicesFixture struct {
	InputFilePath  string
	OutputDir      string
	FilenamePrefix string
	NumberOfSlices int
}

func newAudioSlicesFixture(t *testing.T) audioSlicesFixture {
	t.Helper()

	inputDir := mustCreateTempDir(t, "test-input")
	outputDir := mustCreateTempDir(t, "test-output")
	inputFilePath := filepath.Join(inputDir, "test.wav")
	createTestWAVFile(t, inputFilePath, 44100)

	return audioSlicesFixture{
		InputFilePath:  inputFilePath,
		OutputDir:      outputDir,
		FilenamePrefix: "test_prefix",
		NumberOfSlices: 4,
	}
}

func assertAudioSlices(t *testing.T, fixture audioSlicesFixture, audioFiles *[]ablmodels.AudioFile) {
	t.Helper()

	require.NotNil(t, audioFiles, "writeAudioFileSlices should not return nil audioFiles")
	assert.Len(t, *audioFiles, fixture.NumberOfSlices, "Should create correct number of audio files")

	for i := 0; i < fixture.NumberOfSlices; i++ {
		assertAudioSlice(t, fixture, *audioFiles, i)
	}
}

func assertAudioSlice(t *testing.T, fixture audioSlicesFixture, audioFiles []ablmodels.AudioFile, index int) {
	t.Helper()

	partNum := index + 1
	sliceFilename := filepath.Join(fixture.OutputDir, fixture.FilenamePrefix+"_part_"+fmt.Sprintf("%d", partNum)+".wav")

	_, err := os.Stat(sliceFilename)
	assert.False(t, os.IsNotExist(err), "Slice file should exist at %s", sliceFilename)

	require.NotNil(t, audioFiles[index].FilePath, "Audio file %d should not have nil FilePath", index)

	expectedPath := "Samples/" + fixture.FilenamePrefix + "_part_" + fmt.Sprintf("%d", partNum) + ".wav"
	assert.Equal(t, expectedPath, *audioFiles[index].FilePath, "FilePath should match expected pattern")
	assert.Greater(t, audioFiles[index].Duration, 0.0, "Duration should be positive")
}

func TestWriteAudioFileSlicesWithNonExistentFile(t *testing.T) {
	outputDir := mustCreateTempDir(t, "test-output")
	nonExistentFilePath := "/path/to/nonexistent/file.wav"

	audioFiles, err := writeAudioFileSlices(nonExistentFilePath, outputDir, 4, "test_prefix")

	assert.Error(t, err, "writeAudioFileSlices should fail with non-existent file")
	assert.Nil(t, audioFiles, "audioFiles should be nil when an error occurs")
	assert.Contains(t, err.Error(), "could not open source file", "Error message should indicate the file couldn't be opened")
}
