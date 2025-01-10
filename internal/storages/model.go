package storages

type CurrencyRequest struct {
	FromCurrency string
	ToCurrency   string
}

type ExchangeRateResponse struct {
	FromCurrency string
	ToCurrency   string
	Rate         float32
}
type ExchangeRatesResponse struct {
	Rates map[string]float32
}
