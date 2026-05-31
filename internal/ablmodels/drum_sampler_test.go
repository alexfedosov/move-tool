package ablmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDrumSampler(t *testing.T) {
	sampler := NewDrumSampler()

	assert.Equal(t, DrumSamplerDeviceKind, sampler.Kind, "Kind should be drumCell")
	assert.Nil(t, sampler.DeviceData.SampleURI, "SampleURI should be nil initially")
	assert.Nil(t, sampler.Parameters, "Parameters should be nil initially")
}

func TestDrumSamplerWithSample(t *testing.T) {
	sampler := NewDrumSampler()
	audioFile := sampleAudioFile()

	result := sampler.WithSample(audioFile)

	assertSamplerWithSample(t, sampler, audioFile, result)
}

func sampleAudioFile() AudioFile {
	filePath := "test/sample.wav"
	return AudioFile{
		FilePath: &filePath,
		Duration: 2000.0,
	}
}

func assertSamplerWithSample(t *testing.T, sampler *DrumSampler, audioFile AudioFile, result *DrumSampler) {
	t.Helper()
	assert.NotNil(t, sampler.DeviceData.SampleURI, "SampleURI should not be nil after setting sample")
	assert.Equal(t, *audioFile.FilePath, *sampler.DeviceData.SampleURI, "SampleURI should be set to the file path")
	assert.NotNil(t, sampler.Parameters, "Parameters should be set")
	assert.Same(t, sampler, result, "WithSample should return the same sampler instance")
}
