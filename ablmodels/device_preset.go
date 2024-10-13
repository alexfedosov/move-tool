package ablmodels

type DevicePreset struct {
	Schema string `json:"$schema"`
	Device
}

const DevicePresetSchema = "http://tech.ableton.com/schema/song/1.4.4/devicePreset.json"

func NewDrumRackDevicePresetWithSamples(samples []AudioFile) *DevicePreset {
	drumRack := NewDrumRack()
	for i := 0; i < 16; i++ {
		if i < len(samples) {
			drumRack.AddSample(samples[i])
		} else {
			drumRack.AddSample(AudioFile{
				FilePath: nil,
				Duration: 0,
			})
		}
	}
	return &DevicePreset{
		DevicePresetSchema,
		*NewInstrumentRack().AddChain(NewChain().WithDevice(drumRack).WithDevice(NewSaturator())),
	}
}
