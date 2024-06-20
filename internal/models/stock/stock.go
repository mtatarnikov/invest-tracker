package stock

type Stock struct {
	Ticker           string  `json:"ticker"`
	Price            float64 `json:"price"`
	Pe               float64 `json:"pe"`
	Ps               float64 `json:"ps"`
	MarketCap        float64 `json:"market_cap"`
	AvgDividend5Year float64 `json:"avg_dividend_5_year"` // Сред. див. доходность за 5 лет
}
