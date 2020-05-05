package model

import "database/sql"

//Cotacao tipo que registra uma cotação de uma ação em um determinado dia
type Cotacao struct {
	Ticker sql.NullString  `json:"01. symbol" db:"ticker"`
	Price  sql.NullFloat64 `json:"05. price" db:"price"`
	Dia    sql.NullTime    `json:"07. latest trading day" db:"day"`
}
