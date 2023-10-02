# Currency Library for Go

The Currency library is a Go package for working with monetary 
values and currencies. It provides functionality for 
managing currencies, performing arithmetic operations with 
monetary values, and formatting currency amounts. Currently, it supports the:
- USD
- NGN
- EUR

# To-do
- [ ] Add Currency Conversion logic
- [ ] Fetch and update exchange rates in real-time
- [ ] Handle Decimal Arithmetic a bit better

## Installation

To use this library in your Go project, you can install it using `go get`:

```shell
go get github.com/theghostmac/money
```

## Usage

### Creating and Formatting Money

```go
package main

import (
	"fmt"
	"github.com/theghostmac/money/pkg/currency"
)

func main() {
	// Create a new Money instance with 100 USD
	money := currency.New(100, "USD")

	// Format the money with currency symbol and two decimal places
	formatter := currency.NewFormatter("USD", "$", 2)
	formattedAmount := formatter.FormatWithSymbol(money.Amount)
	fmt.Println(formattedAmount) // Output: $1.00
}
```

### Adding a Custom Currency

```go
package main

import (
	"fmt"
	"github.com/theghostmac/money/pkg/currency"
)

func main() {
	// Add a custom currency to the available currencies
	currency.AddCurrency("CAD", "Canadian Dollar", "$", 2)

	// Create a new Money instance with the custom currency
	money := currency.New(50, "CAD")

	// Format the money with currency symbol and two decimal places
	formatter := currency.NewFormatter("CAD", "$", 2)
	formattedAmount := formatter.FormatWithSymbol(money.Amount)
	fmt.Println(formattedAmount) // Output: $0.50
}
```

### Currency Conversion (not implemented in this example)

```go
package main

import (
	"fmt"
	"github.com/theghostmac/money/pkg/currency"
)

func main() {
	// Create money in USD
	usdMoney := currency.New(100, "USD")

	// Convert USD to EUR using exchange rates
	// (Exchange rate data and conversion logic not shown in this example)
	eurMoney := usdMoney.ConvertTo("EUR")

	// Format the converted money with currency symbol and two decimal places
	formatter := currency.NewFormatter("EUR", "€", 2)
	formattedAmount := formatter.FormatWithSymbol(eurMoney.Amount)
	fmt.Println(formattedAmount) // Output: €0.88
}
```

## Documentation

For detailed documentation and examples, please read comments in the code for now.

## Contributing

I welcome contributions from the community. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
