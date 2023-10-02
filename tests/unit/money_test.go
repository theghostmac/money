package unit

import (
	"github.com/theghostmac/money/pkg/currency"
	"testing"
)

func TestMoneyCurrencyUser(t *testing.T) {
	// Create a Money instance
	m := currency.New(100, "USD")

	// Check if CurrencyUsed returns the correct currency
	currency := m.CurrencyUsed()
	if currency == nil || currency.Code != "USD" {
		t.Errorf("CurrencyUsed() returned an incorrect currency")
	}
}

func TestMoneyAmountUsed(t *testing.T) {
	// Create a Money instance
	m := currency.New(100, "USD")

	// Check if AmountOwned returns the correct amount
	amount := m.AmountOwned()
	if amount != 100 {
		t.Errorf("AmountOwned() returned an incorrect amount")
	}
}

func TestMoneySameCurrency(t *testing.T) {
	// Create two Money instances with the same currency
	m1 := currency.New(100, "USD")
	m2 := currency.New(200, "USD")

	// Check if SameCurrency returns true for the same currency
	sameCurrency := m1.SameCurrency(m2)
	if !sameCurrency {
		t.Errorf("SameCurrency() should return true for the same currency")
	}

	// Create two Money instances with different currencies
	m3 := currency.New(100, "USD")
	m4 := currency.New(200, "EUR")

	// Check if SameCurrency returns false for different currencies
	sameCurrency = m3.SameCurrency(m4)
	if sameCurrency {
		t.Errorf("SameCurrency() should return false for different currencies")
	}
}
