package ablmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDrumSampler(t *testing.T) {
	sampler := NewDrumSampler()

	// Check the device kind is set correctly
	assert.Equal(t, DrumSamplerDeviceKind, sampler.Kind, "Kind should be drumCell")

	// Check the sample URI is initially nil
	assert.Nil(t, sampler.DeviceData.SampleURI, "SampleURI should be nil initially")

	// Check parameters are initially nil
	assert.Nil(t, sampler.Parameters, "Parameters should be nil initially")
}

func TestDrumSamplerWithSample(t *testing.T) {
	sampler := NewDrumSampler()
	filePath := "test/sample.wav"
	duration := 2000.0

	audioFile := AudioFile{
		FilePath: &filePath,
		Duration: duration,
	}

	result := sampler.WithSample(audioFile)

	// Check that the sample URI was set correctly
	assert.NotNil(t, sampler.DeviceData.SampleURI, "SampleURI should not be nil after setting sample")
	assert.Equal(t, filePath, *sampler.DeviceData.SampleURI, "SampleURI should be set to the file path")

	// Check that parameters were set
	assert.NotNil(t, sampler.Parameters, "Parameters should be set")

	// Check that the returned sampler is the same instance
	assert.Same(t, sampler, result, "WithSample should return the same sampler instance")
}