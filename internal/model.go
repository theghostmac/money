package internal

type Amount = int64

type Currency struct {
	Code     string // Currency code (e.g., USD, EUR)
	Name     string // Currency name (e.g., United States Dollar, Euro)
	Symbol   string // Currency symbol (e.g., $, €)
	Decimals int    // Number of decimal places for the currency (e.g., 2 for USD, 2 for EUR)
}

type Currencies map[string]*Currency

type Money struct {
	Amount   Amount
	Currency *Currency
}
