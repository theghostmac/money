package currency

type Amount = int64

// Currency represents a currency
type Currency struct {
	Code     string // Currency code (e.g., USD, EUR)
	Name     string // Currency name (e.g., United States Dollar, Euro)
	Symbol   string // Currency symbol (e.g., $, €)
	Decimals int    // Number of decimal places for the currency (e.g., 2 for USD, 2 for EUR)
}

// Currencies is a map of Currency pointers, with the currency code as the key
type Currencies map[string]*Currency

type Money struct {
	Amount   Amount
	Currency *Currency
}

var AllCurrencies = Currencies{
	"USD": &Currency{
		Code:     "USD",
		Name:     "United States Dollar",
		Symbol:   "$",
		Decimals: 2,
	},
	"EUR": &Currency{
		Code:     "EUR",
		Name:     "Euro",
		Symbol:   "€",
		Decimals: 2,
	},
	"NGN": &Currency{
		Code:     "NGN",
		Name:     "Nigerian Naira",
		Symbol:   "₦",
		Decimals: 0,
	},
}

type Formatter struct {
	Name    string
	Symbol  string
	Decimal int
}
