package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWriteAudioFileSlices verifies that audio files are correctly sliced
// into smaller WAV files with the expected naming pattern and metadata.
func TestWriteAudioFileSlices(t *testing.T) {
	inputDir, err := os.MkdirTemp("", "test-input")
	require.NoError(t, err, "Failed to create input directory")
	defer os.RemoveAll(inputDir)

	outputDir, err := os.MkdirTemp("", "test-output")
	require.NoError(t, err, "Failed to create output directory")
	defer os.RemoveAll(outputDir)

	inputFilePath := filepath.Join(inputDir, "test.wav")
	createTestWAVFile(t, inputFilePath, 44100)

	filenamePrefix := "test_prefix"
	numberOfSlices := 4

	audioFiles, err := writeAudioFileSlices(inputFilePath, outputDir, numberOfSlices, filenamePrefix)
	require.NoError(t, err, "writeAudioFileSlices should not fail")
	require.NotNil(t, audioFiles, "writeAudioFileSlices should not return nil audioFiles")
	assert.Len(t, *audioFiles, numberOfSlices, "Should create correct number of audio files")

	for i := 0; i < numberOfSlices; i++ {
		partNum := i + 1
		sliceFilename := filepath.Join(outputDir, filenamePrefix+"_part_"+fmt.Sprintf("%d", partNum)+".wav")

		_, err = os.Stat(sliceFilename)
		assert.False(t, os.IsNotExist(err), "Slice file should exist at %s", sliceFilename)

		require.NotNil(t, (*audioFiles)[i].FilePath, "Audio file %d should not have nil FilePath", i)

		expectedPath := "Samples/" + filenamePrefix + "_part_" + fmt.Sprintf("%d", partNum) + ".wav"
		assert.Equal(t, expectedPath, *(*audioFiles)[i].FilePath, "FilePath should match expected pattern")
		assert.Greater(t, (*audioFiles)[i].Duration, 0.0, "Duration should be positive")
	}
}

// TestWriteAudioFileSlicesWithNonExistentFile verifies that the function correctly handles
// errors when the input file doesn't exist.
func TestWriteAudioFileSlicesWithNonExistentFile(t *testing.T) {
	outputDir, err := os.MkdirTemp("", "test-output")
	require.NoError(t, err, "Failed to create output directory")
	defer os.RemoveAll(outputDir)

	// Use a non-existent file path
	nonExistentFilePath := "/path/to/nonexistent/file.wav"

	audioFiles, err := writeAudioFileSlices(nonExistentFilePath, outputDir, 4, "test_prefix")

	// Verify that the function returns an error
	assert.Error(t, err, "writeAudioFileSlices should fail with non-existent file")
	assert.Nil(t, audioFiles, "audioFiles should be nil when an error occurs")
	assert.Contains(t, err.Error(), "could not open source file", "Error message should indicate the file couldn't be opened")
}
