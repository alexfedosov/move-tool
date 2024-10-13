package ablmodels

const DrumRackDeviceKind = "drumRack"

type DrumRack = Device

func NewDrumRack() *DrumRack {
	return NewDevice(DrumRackDeviceKind).WithParameters(DefaultInstrumentRackParameters()).AddReturnChain(NewChain().WithDevice(NewReverb()))
}

func (d *DrumRack) AddSample(sample AudioFile) *DrumRack {
	return d.AddChain(NewDrumCellChain(len(d.Chains), sample))
}
