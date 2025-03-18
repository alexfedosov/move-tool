package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSanitizePresetName verifies that preset names are properly sanitized
// by replacing non-lowercase-letters and non-underscores with underscores.
func TestSanitizePresetName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "lowercase_letters_only",
			input:    "abcdefghijklmnopqrstuvwxyz",
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:     "underscore_preserved",
			input:    "drum_sample_kit",
			expected: "drum_sample_kit",
		},
		{
			name:     "uppercase_letters_converted",
			input:    "DrumKit",
			expected: "_rum_it",
		},
		{
			name:     "numbers_converted",
			input:    "kit123",
			expected: "kit___",
		},
		{
			name:     "special_chars_converted",
			input:    "kit@#$%^&",
			expected: "kit______",
		},
		{
			name:     "mixed_content",
			input:    "Drum-Kit_2023!",
			expected: "_rum__it______",
		},
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizePresetName(tt.input)
			assert.Equal(t, tt.expected, result, "sanitizePresetName(%q) should return %q", tt.input, tt.expected)
		})
	}
}

// TestSliceSampleIntoDrumRackWithCustomPresetName verifies that slicing with a custom preset name works
// by creating temporary files and checking if the preset bundle is generated correctly.
func TestSliceSampleIntoDrumRackWithCustomPresetName(t *testing.T) {
	inputDir, err := os.MkdirTemp("", "test-input")
	require.NoError(t, err, "Failed to create input directory")
	defer os.RemoveAll(inputDir)

	outputDir, err := os.MkdirTemp("", "test-output")
	require.NoError(t, err, "Failed to create output directory")
	defer os.RemoveAll(outputDir)

	inputFilePath := filepath.Join(inputDir, "test.wav")
	createTestWAVFile(t, inputFilePath, 44100)

	customPresetName := "my_custom_preset"

	err = SliceSampleIntoDrumRack(inputFilePath, outputDir, 4, customPresetName)
	require.NoError(t, err, "SliceSampleIntoDrumRack should not fail")

	bundlePath := filepath.Join(outputDir, customPresetName+".ablpresetbundle")
	_, err = os.Stat(bundlePath)
	assert.False(t, os.IsNotExist(err), "Preset bundle should exist at %s", bundlePath)
}

// createTestWAVFile generates a minimal valid WAV file for testing
// with the specified sample rate and 8 bytes of silent sample data.
func createTestWAVFile(t *testing.T, filePath string, sampleRate int) {
	header := []byte{
		'R', 'I', 'F', 'F', // ChunkID
		52, 0, 0, 0, // ChunkSize (36 + SubChunk2Size)
		'W', 'A', 'V', 'E', // Format
		'f', 'm', 't', ' ', // Subchunk1ID
		16, 0, 0, 0, // Subchunk1Size
		1, 0, // AudioFormat (1 = PCM)
		1, 0, // NumChannels
		byte(sampleRate & 0xff), byte((sampleRate >> 8) & 0xff), byte((sampleRate >> 16) & 0xff), byte((sampleRate >> 24) & 0xff), // SampleRate
		byte(sampleRate & 0xff), byte((sampleRate >> 8) & 0xff), byte((sampleRate >> 16) & 0xff), byte((sampleRate >> 24) & 0xff), // ByteRate (SampleRate * NumChannels * BitsPerSample/8)
		2, 0, // BlockAlign (NumChannels * BitsPerSample/8)
		16, 0, // BitsPerSample
		'd', 'a', 't', 'a', // Subchunk2ID
		8, 0, 0, 0, // Subchunk2Size (NumSamples * NumChannels * BitsPerSample/8)
		0, 0, 0, 0, 0, 0, 0, 0, // Sample data (8 bytes of silence)
	}

	err := os.WriteFile(filePath, header, 0644)
	require.NoError(t, err, "Failed to create test WAV file")
}
