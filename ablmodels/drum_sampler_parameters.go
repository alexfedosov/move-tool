package ablmodels

type DrumCellParameters struct {
	Effect_EightBitFilterDecay    float64 `json:"Effect_EightBitFilterDecay"`
	Effect_EightBitResamplingRate float64 `json:"Effect_EightBitResamplingRate"`
	Effect_FmAmount               float64 `json:"Effect_FmAmount"`
	Effect_FmFrequency            float64 `json:"Effect_FmFrequency"`
	Effect_LoopLength             float64 `json:"Effect_LoopLength"`
	Effect_LoopOffset             float64 `json:"Effect_LoopOffset"`
	Effect_NoiseAmount            float64 `json:"Effect_NoiseAmount"`
	Effect_NoiseFrequency         float64 `json:"Effect_NoiseFrequency"`
	Effect_On                     bool    `json:"Effect_On"`
	Effect_PitchEnvelopeAmount    float64 `json:"Effect_PitchEnvelopeAmount"`
	Effect_PitchEnvelopeDecay     float64 `json:"Effect_PitchEnvelopeDecay"`
	Effect_PunchAmount            float64 `json:"Effect_PunchAmount"`
	Effect_PunchTime              float64 `json:"Effect_PunchTime"`
	Effect_RingModAmount          float64 `json:"Effect_RingModAmount"`
	Effect_RingModFrequency       float64 `json:"Effect_RingModFrequency"`
	Effect_StretchFactor          float64 `json:"Effect_StretchFactor"`
	Effect_StretchGrainSize       float64 `json:"Effect_StretchGrainSize"`
	Effect_SubOscAmount           float64 `json:"Effect_SubOscAmount"`
	Effect_SubOscFrequency        float64 `json:"Effect_SubOscFrequency"`
	Effect_Type                   string  `json:"Effect_Type"`
	Enabled                       bool    `json:"Enabled"`
	NotePitchBend                 bool    `json:"NotePitchBend"`
	Pan                           float64 `json:"Pan"`
	Voice_Detune                  float64 `json:"Voice_Detune"`
	Voice_Envelope_Attack         float64 `json:"Voice_Envelope_Attack"`
	Voice_Envelope_Decay          float64 `json:"Voice_Envelope_Decay"`
	Voice_Envelope_Hold           float64 `json:"Voice_Envelope_Hold"`
	Voice_Envelope_Mode           string  `json:"Voice_Envelope_Mode"`
	Voice_Filter_Frequency        float64 `json:"Voice_Filter_Frequency"`
	Voice_Filter_On               bool    `json:"Voice_Filter_On"`
	Voice_Filter_PeakGain         float64 `json:"Voice_Filter_PeakGain"`
	Voice_Filter_Resonance        float64 `json:"Voice_Filter_Resonance"`
	Voice_Filter_Type             string  `json:"Voice_Filter_Type"`
	Voice_Gain                    float64 `json:"Voice_Gain"`
	Voice_ModulationAmount        float64 `json:"Voice_ModulationAmount"`
	Voice_ModulationSource        string  `json:"Voice_ModulationSource"`
	Voice_ModulationTarget        string  `json:"Voice_ModulationTarget"`
	Voice_PlaybackLength          float64 `json:"Voice_PlaybackLength"`
	Voice_PlaybackStart           float64 `json:"Voice_PlaybackStart"`
	Voice_Transpose               int     `json:"Voice_Transpose"`
	Voice_VelocityToVolume        float64 `json:"Voice_VelocityToVolume"`
	Volume                        float64 `json:"Volume"`
}

func DefaultDrumSamplerParameters() *DrumCellParameters {
	p := &DrumCellParameters{
		Effect_EightBitFilterDecay:    5.0,
		Effect_EightBitResamplingRate: 14080.0,
		Effect_FmAmount:               0.0,
		Effect_FmFrequency:            999.9998779296876,
		Effect_LoopLength:             0.30000001192092896,
		Effect_LoopOffset:             0.019999997690320015,
		Effect_NoiseAmount:            0.0,
		Effect_NoiseFrequency:         10000.0009765625,
		Effect_On:                     true,
		Effect_PitchEnvelopeAmount:    0.0,
		Effect_PitchEnvelopeDecay:     0.29999998211860657,
		Effect_PunchAmount:            0.0,
		Effect_PunchTime:              0.12015999853610992,
		Effect_RingModAmount:          0.0,
		Effect_RingModFrequency:       999.9998168945313,
		Effect_StretchFactor:          1.0,
		Effect_StretchGrainSize:       0.09999999403953552,
		Effect_SubOscAmount:           0.0,
		Effect_SubOscFrequency:        59.99999237060547,
		Effect_Type:                   "Stretch",
		Enabled:                       true,
		NotePitchBend:                 true,
		Pan:                           0.0,
		Voice_Detune:                  0.0,
		Voice_Envelope_Attack:         0.00009999999747378752,
		Voice_Envelope_Decay:          0.3,
		Voice_Envelope_Hold:           1,
		Voice_Envelope_Mode:           "A-S-R",
		Voice_Filter_Frequency:        21999.990234375,
		Voice_Filter_On:               true,
		Voice_Filter_PeakGain:         1.0,
		Voice_Filter_Resonance:        0.0,
		Voice_Filter_Type:             "Lowpass",
		Voice_Gain:                    1.0,
		Voice_ModulationAmount:        0.0,
		Voice_ModulationSource:        "Velocity",
		Voice_ModulationTarget:        "Filter",
		Voice_PlaybackLength:          1.0,
		Voice_PlaybackStart:           0.0,
		Voice_Transpose:               0,
		Voice_VelocityToVolume:        0.3499999940395355,
		Volume:                        0,
	}
	return p.WithTriggerMode()
}

func (p *DrumCellParameters) WithGateMode() *DrumCellParameters {
	p.Voice_Envelope_Mode = "A-S-R"
	return p
}

func (p *DrumCellParameters) WithTriggerMode() *DrumCellParameters {
	p.Voice_Envelope_Mode = "A-H-D"
	return p
}

func (p *DrumCellParameters) WithVoiceEnvelopeHold(value float64) *DrumCellParameters {
	p.Voice_Envelope_Hold = value
	return p
}
