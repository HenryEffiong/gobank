package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	NGN = "NGN"
)

// returns true if input currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, NGN:
		return true
	}
	return false
}
