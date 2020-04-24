package model

type Cotacao struct {
	Ticker string  `json:"01. symbol"`
	Price  float32 `json:"05. price"`
	Dia    int32   `json:"05. price"`
}
