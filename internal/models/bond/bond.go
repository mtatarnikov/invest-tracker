package bond

type Bond struct {
	Ticker          string  `json:"ticker"`
	Price           float64 `json:"price"`
	ProfitPercent   float64 `json:"profit_percent"`    // Доходность к погашению
	DateExpire      string  `json:"date_expire"`       // Дата погашения облигации
	CouponValue     float64 `json:"coupon_value"`      // Величина купона
	PaymentsPerYear int8    `json:"payments_per_year"` // Количество выплат в год
}
