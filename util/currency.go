package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	JPY = "JPY"
	CAD = "CAD"
	AUD = "AUD"
)

func isSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, JPY, CAD, AUD:
		return true
	}
	return false
}
