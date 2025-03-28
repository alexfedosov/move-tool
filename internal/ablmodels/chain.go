package ablmodels

type Chain struct {
	Name    string        `json:"name"`
	Color   int           `json:"color"`
	Devices []interface{} `json:"devices"`
	Mixer   Mixer         `json:"mixer"`
}

func NewChain() *Chain {
	return &Chain{
		Name:    "",
		Color:   2,
		Devices: make([]interface{}, 0),
		Mixer:   *NewMixer(),
	}
}

func (c *Chain) WithDevice(device interface{}) *Chain {
	c.Devices = append(c.Devices, device)
	return c
}
