package unit

import (
	"fmt"
	"github.com/theghostmac/money/internal/currency"
	"testing"
)

func TestAddCurrency(t *testing.T) {
	// Ensure that adding a new currency works as expected
	newCurrency, err := currency.AddCurrency("GHS", "Ghanaian Cedi", "₵", 2)
	if err != nil {
		t.Errorf("Error adding currency: %v", err)
	}

	fmt.Println(newCurrency)

	// Ensure that newly added currency exists in AllCurrencies map.
	if currency.AllCurrencies["GHS"] == nil {
		t.Errorf("Currency 'GHS' not found in AllCurrencies after adding")
	}
}

func TestAddCurrencyDuplicate(t *testing.T) {
	// Attempt to add a currency with an existing code, should return an error
	duplicateCurrency, err := currency.AddCurrency("USD", "Duplicate Dollar", "$", 2)
	if err == nil {
		t.Errorf("Expected error when adding duplicate currency, got nil")
	}
	fmt.Println(duplicateCurrency)
}

func TestGetCurrencyByCode(t *testing.T) {
	// Test getting an existing currency by code
	currencyToGet := currency.GetCurrencyByCode("USD")
	if currencyToGet == nil {
		t.Errorf("Expected to find USD currency, got nil")
	}

	// Test getting a non-existing currency by code
	currencyToGet = currency.GetCurrencyByCode("XYZ")
	if currencyToGet != nil {
		t.Errorf("Expeced to not find XYZ currency, got %v", currencyToGet)
	}
}
