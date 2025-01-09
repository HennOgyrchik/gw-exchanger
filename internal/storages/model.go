package storages

type CurrencyRequest struct {
	fromCurrency string
	toCurrency   string
}

type ExchangeRateResponse struct {
	fromCurrency string
	toCurrency   string
	rate         float64
}
type ExchangeRatesResponse struct {
	rates map[string]float64
}
