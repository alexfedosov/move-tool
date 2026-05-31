package ablmodels

type ReverbParameters struct {
	ReverbRoomParameters
	ReverbRoomModeParameters
	ReverbBandParameters
	ReverbReflectionParameters
	ReverbMixParameters
	ReverbShelfParameters
	ReverbModulationParameters
}

type ReverbRoomParameters struct {
	DecayTime float64 `json:"DecayTime"`
	Enabled   bool    `json:"Enabled"`
	FlatOn    bool    `json:"FlatOn"`
	FreezeOn  bool    `json:"FreezeOn"`
	PreDelay  float64 `json:"PreDelay"`
	RoomSize  float64 `json:"RoomSize"`
}

type ReverbRoomModeParameters struct {
	HighFilterType   string  `json:"HighFilterType"`
	RoomType         string  `json:"RoomType"`
	StereoSeparation float64 `json:"StereoSeparation"`
}

type ReverbBandParameters struct {
	BandFreq   float64 `json:"BandFreq"`
	BandHighOn bool    `json:"BandHighOn"`
	BandLowOn  bool    `json:"BandLowOn"`
	BandWidth  float64 `json:"BandWidth"`
	CutOn      bool    `json:"CutOn"`
}

type ReverbReflectionParameters struct {
	AllPassGain          float64 `json:"AllPassGain"`
	AllPassSize          float64 `json:"AllPassSize"`
	DiffuseDelay         float64 `json:"DiffuseDelay"`
	EarlyReflectModDepth float64 `json:"EarlyReflectModDepth"`
	EarlyReflectModFreq  float64 `json:"EarlyReflectModFreq"`
}

type ReverbMixParameters struct {
	MixDiffuse float64 `json:"MixDiffuse"`
	MixDirect  float64 `json:"MixDirect"`
	MixReflect float64 `json:"MixReflect"`
}

type ReverbShelfParameters struct {
	ShelfHiFreq float64 `json:"ShelfHiFreq"`
	ShelfHiGain float64 `json:"ShelfHiGain"`
	ShelfHighOn bool    `json:"ShelfHighOn"`
	ShelfLoFreq float64 `json:"ShelfLoFreq"`
	ShelfLoGain float64 `json:"ShelfLoGain"`
	ShelfLowOn  bool    `json:"ShelfLowOn"`
}

type ReverbModulationParameters struct {
	ChorusOn      bool    `json:"ChorusOn"`
	SizeModDepth  float64 `json:"SizeModDepth"`
	SizeModFreq   float64 `json:"SizeModFreq"`
	SizeSmoothing string  `json:"SizeSmoothing"`
	SpinOn        bool    `json:"SpinOn"`
}

func DefaultReverbParameters() ReverbParameters {
	return ReverbParameters{
		ReverbRoomParameters: ReverbRoomParameters{
			DecayTime: 1200.0001220703125,
			Enabled:   true,
			FlatOn:    true,
			FreezeOn:  false,
			PreDelay:  2.5,
			RoomSize:  99.99999237060548,
		},
		ReverbRoomModeParameters: ReverbRoomModeParameters{
			HighFilterType:   "Shelf",
			RoomType:         "SuperEco",
			StereoSeparation: 100.0,
		},
		ReverbBandParameters: ReverbBandParameters{
			BandFreq:   829.999755859375,
			BandHighOn: false,
			BandLowOn:  true,
			BandWidth:  5.849999904632568,
			CutOn:      true,
		},
		ReverbReflectionParameters: ReverbReflectionParameters{
			AllPassGain:          0.6000000238418579,
			AllPassSize:          0.4000000059604645,
			DiffuseDelay:         0.5,
			EarlyReflectModDepth: 17.5,
			EarlyReflectModFreq:  0.29770001769065857,
		},
		ReverbMixParameters: ReverbMixParameters{
			MixDiffuse: 1.0,
			MixDirect:  0.550000011920929,
			MixReflect: 1.0,
		},
		ReverbShelfParameters: ReverbShelfParameters{
			ShelfHiFreq: 4500.00146484375,
			ShelfHiGain: 0.699999988079071,
			ShelfHighOn: true,
			ShelfLoFreq: 90.0,
			ShelfLoGain: 1.0,
			ShelfLowOn:  true,
		},
		ReverbModulationParameters: ReverbModulationParameters{
			ChorusOn:      true,
			SizeModDepth:  0.019999999552965164,
			SizeModFreq:   0.020000001415610313,
			SizeSmoothing: "Fast",
			SpinOn:        true,
		},
	}
}

const ReverbDeviceKind = "reverb"

func NewReverb() *Device {
	return NewDevice(ReverbDeviceKind).WithParameters(DefaultReverbParameters())
}
