package currency

import "strings"

// Add adds a new Currency to the Currencies map
func (c Currencies) Add(currency *Currency) Currencies {
	c[currency.Code] = currency
	return c
}

// CurrencyByCode returns a Currency pointer for the given currency code
func (c Currencies) CurrencyByCode(code string) *Currency {
	getCurrency, ok := c[code]
	if !ok {
		return nil
	}
	return getCurrency
}

// Formatter returns a Formatter for the Currency
func (c *Currency) Formatter() *Formatter {
	return &Formatter{
		Name:    c.Name,
		Symbol:  c.Symbol,
		Decimal: c.Decimals,
	}
}

// AddCurrency adds a new Currency and updates AllCurrencies
func AddCurrency(code, name, symbol string, decimal int) *Currency {
	newCurrency := Currency{
		Code:     code,
		Name:     name,
		Symbol:   symbol,
		Decimals: decimal,
	}

	AllCurrencies.Add(&newCurrency)
	return &newCurrency
}

// NewCurrency returns a new Currency for the given code
func NewCurrency(code string) *Currency {
	return &Currency{
		Code: strings.ToUpper(code),
	}
}

// GetCurrencyByCode returns a Currency for the given code from AllCurrencies
func GetCurrencyByCode(code string) *Currency {
	return AllCurrencies.CurrencyByCode(code)
}
