package ablmodels

type Device struct {
	PresetURI    interface{}   `json:"presetUri"`
	Kind         string        `json:"kind"`
	Name         string        `json:"name"`
	Parameters   interface{}   `json:"parameters"`
	Chains       []interface{} `json:"chains,omitempty"`
	ReturnChains []interface{} `json:"returnChains,omitempty"`
}

func NewDevice(kind string) *Device {
	return &Device{
		PresetURI:    nil,
		Kind:         kind,
		Name:         "",
		Parameters:   nil,
		Chains:       make([]interface{}, 0),
		ReturnChains: make([]interface{}, 0),
	}
}

func (d *Device) AddChain(chain interface{}) *Device {
	d.Chains = append(d.Chains, chain)
	return d
}

func (d *Device) AddReturnChain(chain interface{}) *Device {
	d.ReturnChains = append(d.ReturnChains, chain)
	return d
}

func (d *Device) WithParameters(parameters interface{}) *Device {
	d.Parameters = parameters
	return d
}
