package ablmodels

type SaturatorParameters struct {
	BaseDrive           float64 `json:"BaseDrive"`
	BassShaperThreshold float64 `json:"BassShaperThreshold"`
	ColorDepth          float64 `json:"ColorDepth"`
	ColorFrequency      float64 `json:"ColorFrequency"`
	ColorOn             bool    `json:"ColorOn"`
	ColorWidth          float64 `json:"ColorWidth"`
	DryWet              float64 `json:"DryWet"`
	Enabled             bool    `json:"Enabled"`
	Oversampling        bool    `json:"Oversampling"`
	PostClip            string  `json:"PostClip"`
	PostDrive           float64 `json:"PostDrive"`
	PreDcFilter         bool    `json:"PreDcFilter"`
	PreDrive            float64 `json:"PreDrive"`
	Type                string  `json:"Type"`
	WsCurve             float64 `json:"WsCurve"`
	WsDamp              float64 `json:"WsDamp"`
	WsDepth             float64 `json:"WsDepth"`
	WsDrive             float64 `json:"WsDrive"`
	WsLin               float64 `json:"WsLin"`
	WsPeriod            float64 `json:"WsPeriod"`
}

func DefaultSaturatorParameters() SaturatorParameters {
	return SaturatorParameters{
		BaseDrive:           -20.25,
		BassShaperThreshold: -50.0,
		ColorDepth:          0.0,
		ColorFrequency:      999.9998779296876,
		ColorOn:             true,
		ColorWidth:          0.30000001192092896,
		DryWet:              0.2936508059501648,
		Enabled:             true,
		Oversampling:        true,
		PostClip:            "off",
		PostDrive:           -23.714284896850582,
		PreDcFilter:         false,
		PreDrive:            20.571426391601563,
		Type:                "Soft Sine",
		WsCurve:             0.05000000074505806,
		WsDamp:              0.0,
		WsDepth:             0.0,
		WsDrive:             1.0,
		WsLin:               0.5,
		WsPeriod:            0.0,
	}
}

const SaturatorDeviceKind = "saturator"

func NewSaturator() *Device {
	return NewDevice(SaturatorDeviceKind).WithParameters(DefaultSaturatorParameters())
}
