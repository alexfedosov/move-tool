package ablmodels

const DrumSamplerDeviceKind = "drumCell"

type DrumSampler struct {
	Device
	DeviceData DeviceData `json:"deviceData"`
}

type DeviceData struct {
	SampleURI *string `json:"sampleUri"`
}

func NewDrumSampler() *DrumSampler {
	return &DrumSampler{
		*NewDevice(DrumSamplerDeviceKind),
		DeviceData{nil},
	}
}

func (s *DrumSampler) WithSample(file AudioFile) *DrumSampler {
	s.DeviceData.SampleURI = file.FilePath
	s.Parameters = DefaultDrumSamplerParameters().WithVoiceEnvelopeHold(file.Duration).WithGateMode()
	return s
}
