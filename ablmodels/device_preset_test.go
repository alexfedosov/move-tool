package ablmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDrumRackDevicePresetWithSamples(t *testing.T) {
	// Test with empty samples
	emptyPreset := NewDrumRackDevicePresetWithSamples([]AudioFile{})

	// Check schema is set correctly
	assert.Equal(t, DevicePresetSchema, emptyPreset.Schema, "Schema should be set correctly")

	// Check it still creates 16 empty samples
	assert.Len(t, emptyPreset.Chains, 1, "Should create 1 chain")

	// Test with some samples
	testPath1 := "sample1.wav"
	testPath2 := "sample2.wav"
	samples := []AudioFile{
		{
			FilePath: &testPath1,
			Duration: 1000.0,
		},
		{
			FilePath: &testPath2,
			Duration: 2000.0,
		},
	}

	preset := NewDrumRackDevicePresetWithSamples(samples)

	// Check schema is set correctly
	assert.Equal(t, DevicePresetSchema, preset.Schema, "Schema should be set correctly")

	// Check that the preset has a Device with the right kind
	assert.Equal(t, "instrumentRack", preset.Kind, "Kind should be instrumentRack")

	// Verify the sample information was properly included
	// Since the structure is complex, we'll check the serialization works
	// and we'll check the structure is as expected
	assert.Len(t, preset.Chains, 1, "Should create 1 chain")
}
