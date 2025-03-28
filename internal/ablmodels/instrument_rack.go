package ablmodels

type InstrumentRackParameters struct {
	Enabled bool    `json:"Enabled"`
	Macro0  float64 `json:"Macro0"`
	Macro1  float64 `json:"Macro1"`
	Macro2  float64 `json:"Macro2"`
	Macro3  float64 `json:"Macro3"`
	Macro4  float64 `json:"Macro4"`
	Macro5  float64 `json:"Macro5"`
	Macro6  float64 `json:"Macro6"`
	Macro7  float64 `json:"Macro7"`
}

func DefaultInstrumentRackParameters() *InstrumentRackParameters {
	return &InstrumentRackParameters{
		Enabled: true,
		Macro0:  0,
		Macro1:  0,
		Macro2:  0,
		Macro3:  0,
		Macro4:  0,
		Macro5:  0,
		Macro6:  0,
		Macro7:  0,
	}
}

func NewInstrumentRack() *Device {
	return NewDevice("instrumentRack").WithParameters(DefaultInstrumentRackParameters())
}
