package model

//Stock é uma ação
type Stock struct {
	Ticker string  `json:"01. symbol"`
	Price  float32 `json:"05. price"`
}
