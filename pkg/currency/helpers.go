package currency

import (
	"encoding/json"
	"errors"
	"fmt"
)

func NewFormatter(name string, symbol string, decimal int) *Formatter {
	return &Formatter{
		Name:    name,
		Symbol:  symbol,
		Decimal: decimal,
	}
}

func (f *Formatter) Format(amount Amount) string {
	amountStr := fmt.Sprintf("%d", amount)
	decimalIndex := len(amountStr) - f.Decimal
	if decimalIndex < 0 {
		decimalIndex = 0
	}
	formattedAmount := amountStr[:decimalIndex] + "." + amountStr[decimalIndex:]
	return f.Symbol + formattedAmount
}

func (f *Formatter) FormatWithSymbol(amount Amount) string {
	// Include the currency symbol in the formatted amount
	formattedAmount := f.Format(amount)
	return formattedAmount
}

func (f *Formatter) ToMajorUnits(amount Amount) Amount {
	// Convert the amount from minor units (e.g., cents) to major units (e.g., dollars)
	return amount / Amount(10^f.Decimal)
}

func (f *Formatter) Abs(amount Amount) Amount {
	if amount < 0 {
		return -amount
	}

	return amount
}

var (
	UnmarshallJSON = defaultUnMarshallJSON
	MarshallJSON   = defaultMarshallJSON

	ErrorCurrencyMisMatch = errors.New("Currencies do not match")

	InvalidJSONUnmarshall = errors.New("Invalid JSON marshal")
)

func defaultUnMarshallJSON(m *Money, b []byte) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	var amount Amount
	if amountRaw, ok := data["amount"]; ok {
		amount, ok = amountRaw.(Amount)
		if !ok {
			return InvalidJSONUnmarshall
		}
	}

	var currency string
	if currencyRaw, ok := data["currency"]; ok {
		currency, ok = currencyRaw.(string)
		if !ok {
			return InvalidJSONUnmarshall
		}
	}

	var ref *Money
	if amount == 0 && currency == "" {
		ref = &Money{}
	} else {
		ref = New(int64(amount), currency)
	}

	*m = *ref
	return nil
}

func defaultMarshallJSON(m *Money) ([]byte, error) {
	if m == nil || m.Currency == nil {
		return nil, errors.New("Money or Currency is nil")
	}

	data := struct {
		Amount   Amount `json:"amount"`
		Currency string `json:"currency"`
	}{
		Amount:   m.Amount,
		Currency: m.Currency.Code,
	}

	return json.Marshal(data)
}
