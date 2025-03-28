package ablmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDrumRackDevicePresetWithSamples verifies that creating drum rack presets works correctly
// with both empty samples and provided sample data.
func TestNewDrumRackDevicePresetWithSamples(t *testing.T) {
	emptyPreset := NewDrumRackDevicePresetWithSamples([]AudioFile{})

	assert.Equal(t, DevicePresetSchema, emptyPreset.Schema, "Schema should be set correctly")
	assert.Len(t, emptyPreset.Chains, 1, "Should create 1 chain")

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

	assert.Equal(t, DevicePresetSchema, preset.Schema, "Schema should be set correctly")
	assert.Equal(t, "instrumentRack", preset.Kind, "Kind should be instrumentRack")
	assert.Len(t, preset.Chains, 1, "Should create 1 chain")
}
