package app

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteAudioFileSlices(t *testing.T) {
	// Create temporary directories for testing
	inputDir, err := os.MkdirTemp("", "test-input")
	require.NoError(t, err, "Failed to create input directory")
	defer os.RemoveAll(inputDir)

	outputDir, err := os.MkdirTemp("", "test-output")
	require.NoError(t, err, "Failed to create output directory")
	defer os.RemoveAll(outputDir)

	// Create a simple WAV file for testing
	inputFilePath := filepath.Join(inputDir, "test.wav")
	createTestWAVFile(t, inputFilePath, 44100)

	// Test with 4 slices
	filenamePrefix := "test_prefix"
	numberOfSlices := 4

	audioFiles, err := writeAudioFileSlices(inputFilePath, outputDir, numberOfSlices, filenamePrefix)
	require.NoError(t, err, "writeAudioFileSlices should not fail")
	require.NotNil(t, audioFiles, "writeAudioFileSlices should not return nil audioFiles")
	assert.Len(t, *audioFiles, numberOfSlices, "Should create correct number of audio files")

	// Check if all the slice files were created
	for i := 0; i < numberOfSlices; i++ {
		partNum := i + 1
		sliceFilename := filepath.Join(outputDir, filenamePrefix+"_part_"+fmt.Sprintf("%d", partNum)+".wav")

		// Check if the file was created on disk
		_, err = os.Stat(sliceFilename)
		assert.False(t, os.IsNotExist(err), "Slice file should exist at %s", sliceFilename)

		require.NotNil(t, (*audioFiles)[i].FilePath, "Audio file %d should not have nil FilePath", i)

		expectedPath := "Samples/" + filenamePrefix + "_part_" + fmt.Sprintf("%d", partNum) + ".wav"
		assert.Equal(t, expectedPath, *(*audioFiles)[i].FilePath, "FilePath should match expected pattern")
		assert.Greater(t, (*audioFiles)[i].Duration, 0.0, "Duration should be positive")
	}
}
