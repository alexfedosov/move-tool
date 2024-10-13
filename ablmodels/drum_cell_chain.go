package ablmodels

type DrumCellChain struct {
	Chain
	DrumZoneSettings *DrumZoneSettings `json:"drumZoneSettings"`
}

type DrumZoneSettings struct {
	ReceivingNote int         `json:"receivingNote"`
	SendingNote   int         `json:"sendingNote"`
	ChokeGroup    interface{} `json:"chokeGroup"`
}

func NewDrumCellChain(padIndex int, sample AudioFile) *DrumCellChain {
	chain := NewChain().WithDevice(NewDrumSampler().WithSample(sample))
	chain.Mixer = *NewMixer().WithDefaultSend()
	return &DrumCellChain{
		*chain,
		&DrumZoneSettings{
			SendingNote:   60,
			ReceivingNote: 36 + padIndex,
		},
	}
}
