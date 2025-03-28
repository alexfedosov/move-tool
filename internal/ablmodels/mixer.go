package ablmodels

type Mixer struct {
	Pan       float64 `json:"pan"`
	SoloCue   bool    `json:"solo-cue"`
	SpeakerOn bool    `json:"speakerOn"`
	Volume    float64 `json:"volume"`
	Sends     []Send  `json:"sends"`
}

type Send struct {
	IsEnabled bool    `json:"isEnabled"`
	Amount    float64 `json:"amount"`
}

func NewMixer() *Mixer {
	return &Mixer{
		Pan:       0,
		SoloCue:   false,
		SpeakerOn: true,
		Volume:    0,
		Sends:     make([]Send, 0),
	}
}

func (m *Mixer) WithDefaultSend() *Mixer {
	m.Sends = append(m.Sends, Send{
		IsEnabled: true,
		Amount:    0,
	})
	return m
}
