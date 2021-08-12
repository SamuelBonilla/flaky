package domain

type House struct {
	Id        int32  `json:"id"`
	Address   string `json:"address"`
	Homeowner string `json:"homeowner"`
	Price     int64  `json:"price"`
	PhotoURL  string `json:"photoURL"`
}

type HomeVision struct {
	Houses []House `json:"houses"`
}

