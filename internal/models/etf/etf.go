package etf

type Etf struct {
	Ticker  string  `json:"ticker"`
	Price   float64 `json:"price"`    //Цена пая
	FundTax float64 `json:"fund_tax"` //Расходы фонда
}
