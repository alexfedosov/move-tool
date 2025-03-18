package ablmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDrumSampler verifies that a new DrumSampler is created with the correct initial values
// including kind, empty sample URI, and nil parameters.
func TestNewDrumSampler(t *testing.T) {
	sampler := NewDrumSampler()

	assert.Equal(t, DrumSamplerDeviceKind, sampler.Kind, "Kind should be drumCell")
	assert.Nil(t, sampler.DeviceData.SampleURI, "SampleURI should be nil initially")
	assert.Nil(t, sampler.Parameters, "Parameters should be nil initially")
}

// TestDrumSamplerWithSample verifies that adding a sample to a DrumSampler works correctly
// by setting the SampleURI and initializing parameters based on sample duration.
func TestDrumSamplerWithSample(t *testing.T) {
	sampler := NewDrumSampler()
	filePath := "test/sample.wav"
	duration := 2000.0

	audioFile := AudioFile{
		FilePath: &filePath,
		Duration: duration,
	}

	result := sampler.WithSample(audioFile)

	assert.NotNil(t, sampler.DeviceData.SampleURI, "SampleURI should not be nil after setting sample")
	assert.Equal(t, filePath, *sampler.DeviceData.SampleURI, "SampleURI should be set to the file path")
	assert.NotNil(t, sampler.Parameters, "Parameters should be set")
	assert.Same(t, sampler, result, "WithSample should return the same sampler instance")
}
