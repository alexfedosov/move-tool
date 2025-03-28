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
