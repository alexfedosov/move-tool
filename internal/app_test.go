package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		{
			name:     "all_uppercase_letters",
			input:    "ABCDEFG",
			expected: "_______",
		},
		{
			name:     "mixed_case_with_valid_chars",
			input:    "aBcDeFg_123",
			expected: "a_c_e_g____",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizePresetName(tt.input)
			assert.Equal(t, tt.expected, result, "sanitizePresetName(%q) should return %q", tt.input, tt.expected)
		})
	}
}

func TestSliceSampleIntoDrumRackWithCustomPresetName(t *testing.T) {
	fixture := newSliceDrumRackFixture(t)
	customPresetName := "my_custom_preset"

	err := SliceSampleIntoDrumRack(fixture.InputFilePath, fixture.OutputDir, 4, customPresetName)
	require.NoError(t, err, "SliceSampleIntoDrumRack should not fail")

	fixture.assertPresetBundleExists(t, customPresetName)
}

type sliceDrumRackFixture struct {
	InputFilePath string
	OutputDir     string
}

func newSliceDrumRackFixture(t *testing.T) sliceDrumRackFixture {
	t.Helper()

	inputDir := mustCreateTempDir(t, "test-input")
	outputDir := mustCreateTempDir(t, "test-output")
	inputFilePath := filepath.Join(inputDir, "test.wav")
	createTestWAVFile(t, inputFilePath, 44100)

	return sliceDrumRackFixture{
		InputFilePath: inputFilePath,
		OutputDir:     outputDir,
	}
}

func mustCreateTempDir(t *testing.T, pattern string) string {
	t.Helper()

	dir, err := os.MkdirTemp("", pattern)
	require.NoError(t, err, "Failed to create temp directory")
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	return dir
}

func (f sliceDrumRackFixture) assertPresetBundleExists(t *testing.T, customPresetName string) {
	t.Helper()

	bundlePath := filepath.Join(f.OutputDir, customPresetName+".ablpresetbundle")
	_, err := os.Stat(bundlePath)
	assert.False(t, os.IsNotExist(err), "Preset bundle should exist at %s", bundlePath)
}

func createTestWAVFile(t *testing.T, filePath string, sampleRate int) {
	header := []byte{
		'R', 'I', 'F', 'F',
		52, 0, 0, 0,
		'W', 'A', 'V', 'E',
		'f', 'm', 't', ' ',
		16, 0, 0, 0,
		1, 0,
		1, 0,
		byte(sampleRate & 0xff), byte((sampleRate >> 8) & 0xff), byte((sampleRate >> 16) & 0xff), byte((sampleRate >> 24) & 0xff),
		byte(sampleRate & 0xff), byte((sampleRate >> 8) & 0xff), byte((sampleRate >> 16) & 0xff), byte((sampleRate >> 24) & 0xff),
		2, 0,
		16, 0,
		'd', 'a', 't', 'a',
		8, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	err := os.WriteFile(filePath, header, 0644)
	require.NoError(t, err, "Failed to create test WAV file")
}
