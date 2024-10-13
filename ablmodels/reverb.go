package ablmodels

type ReverbParameters struct {
	AllPassGain          float64 `json:"AllPassGain"`
	AllPassSize          float64 `json:"AllPassSize"`
	BandFreq             float64 `json:"BandFreq"`
	BandHighOn           bool    `json:"BandHighOn"`
	BandLowOn            bool    `json:"BandLowOn"`
	BandWidth            float64 `json:"BandWidth"`
	ChorusOn             bool    `json:"ChorusOn"`
	CutOn                bool    `json:"CutOn"`
	DecayTime            float64 `json:"DecayTime"`
	DiffuseDelay         float64 `json:"DiffuseDelay"`
	EarlyReflectModDepth float64 `json:"EarlyReflectModDepth"`
	EarlyReflectModFreq  float64 `json:"EarlyReflectModFreq"`
	Enabled              bool    `json:"Enabled"`
	FlatOn               bool    `json:"FlatOn"`
	FreezeOn             bool    `json:"FreezeOn"`
	HighFilterType       string  `json:"HighFilterType"`
	MixDiffuse           float64 `json:"MixDiffuse"`
	MixDirect            float64 `json:"MixDirect"`
	MixReflect           float64 `json:"MixReflect"`
	PreDelay             float64 `json:"PreDelay"`
	RoomSize             float64 `json:"RoomSize"`
	RoomType             string  `json:"RoomType"`
	ShelfHiFreq          float64 `json:"ShelfHiFreq"`
	ShelfHiGain          float64 `json:"ShelfHiGain"`
	ShelfHighOn          bool    `json:"ShelfHighOn"`
	ShelfLoFreq          float64 `json:"ShelfLoFreq"`
	ShelfLoGain          float64 `json:"ShelfLoGain"`
	ShelfLowOn           bool    `json:"ShelfLowOn"`
	SizeModDepth         float64 `json:"SizeModDepth"`
	SizeModFreq          float64 `json:"SizeModFreq"`
	SizeSmoothing        string  `json:"SizeSmoothing"`
	SpinOn               bool    `json:"SpinOn"`
	StereoSeparation     float64 `json:"StereoSeparation"`
}

func DefaultReverbParameters() ReverbParameters {
	return ReverbParameters{
		AllPassGain:          0.6000000238418579,
		AllPassSize:          0.4000000059604645,
		BandFreq:             829.999755859375,
		BandHighOn:           false,
		BandLowOn:            true,
		BandWidth:            5.849999904632568,
		ChorusOn:             true,
		CutOn:                true,
		DecayTime:            1200.0001220703125,
		DiffuseDelay:         0.5,
		EarlyReflectModDepth: 17.5,
		EarlyReflectModFreq:  0.29770001769065857,
		Enabled:              true,
		FlatOn:               true,
		FreezeOn:             false,
		HighFilterType:       "Shelf",
		MixDiffuse:           1.0,
		MixDirect:            0.550000011920929,
		MixReflect:           1.0,
		PreDelay:             2.5,
		RoomSize:             99.99999237060548,
		RoomType:             "SuperEco",
		ShelfHiFreq:          4500.00146484375,
		ShelfHiGain:          0.699999988079071,
		ShelfHighOn:          true,
		ShelfLoFreq:          90.0,
		ShelfLoGain:          1.0,
		ShelfLowOn:           true,
		SizeModDepth:         0.019999999552965164,
		SizeModFreq:          0.020000001415610313,
		SizeSmoothing:        "Fast",
		SpinOn:               true,
		StereoSeparation:     100.0,
	}
}

const ReverbDeviceKind = "reverb"

func NewReverb() *Device {
	return NewDevice(ReverbDeviceKind).WithParameters(DefaultReverbParameters())
}
