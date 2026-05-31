package ablmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDrumRackDevicePresetWithSamples(t *testing.T) {
	assertEmptyDrumRackPreset(t)

	preset := NewDrumRackDevicePresetWithSamples(sampleAudioFiles())

	assertDrumRackPresetWithSamples(t, preset)
}

func assertEmptyDrumRackPreset(t *testing.T) {
	t.Helper()

	emptyPreset := NewDrumRackDevicePresetWithSamples([]AudioFile{})

	assert.Equal(t, DevicePresetSchema, emptyPreset.Schema, "Schema should be set correctly")
	assert.Len(t, emptyPreset.Chains, 1, "Should create 1 chain")
}

func sampleAudioFiles() []AudioFile {
	testPath1 := "sample1.wav"
	testPath2 := "sample2.wav"
	return []AudioFile{
		{
			FilePath: &testPath1,
			Duration: 1000.0,
		},
		{
			FilePath: &testPath2,
			Duration: 2000.0,
		},
	}
}

func assertDrumRackPresetWithSamples(t *testing.T, preset *DevicePreset) {
	t.Helper()
	assert.Equal(t, DevicePresetSchema, preset.Schema, "Schema should be set correctly")
	assert.Equal(t, "instrumentRack", preset.Kind, "Kind should be instrumentRack")
	assert.Len(t, preset.Chains, 1, "Should create 1 chain")
}
