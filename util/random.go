package util

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer between two values.
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max - min + 1)
}

// RandomString returns a random string of a given length.
func RandomString(length int) string {
	var sb strings.Builder
	n := len(chars)

	for i := 0; i < length; i++ {
		sb.WriteByte(chars[rand.Intn(n)])
	}

	return sb.String()
}

// RandomOwner returns a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random currency.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random currency.
func RandomCurrency() string {
	var currencies = []string{"USD", "EUR", "GBP", "AUD", "CAD", "CHF", "CNY", "DKK", "HKD", "JPY", "KRW", "NZD", "PLN", "RUB", "SEK", "SGD", "THB", "TRY", "ZAR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// RandomEmail returns a random email.
func RandomEmail() string {
	return RandomString(5) + "@" + RandomString(10) + ".com"
}
