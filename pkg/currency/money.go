package currency

// New is a constructor for a new instance of Money.
func New(amount Amount, code string) *Money {
	return &Money{
		Amount:   amount,
		Currency: NewCurrency(code).Get(),
	}
}

func (m *Money) CurrencyUsed() *Currency {
	return m.Currency
}

func (m *Money) AmountOwned() Amount {
	return m.Amount
}

func (m *Money) SameCurrency(originalMoney *Money) bool {
	return m.Currency.Equals(originalMoney.Currency)
}

func (m *Money) AssertSameCurrency(originalMoney *Money) error {
	if !m.SameCurrency(originalMoney) {
		return ErrorCurrencyMisMatch
	}

	return nil
}

func (m *Money) Compare(originalMoney *Money) int {
	switch {
	case m.Amount > originalMoney.Amount:
		return 1
	case m.Amount < originalMoney.Amount:
		return -1
	}

	return 0
}

func (m *Money) TwoMoniesEqual(originalMoney *Money) (bool, error) {
	if err := m.AssertSameCurrency(originalMoney); err != nil {
		return false, err
	}

	return m.Compare(originalMoney) == 1, nil
}

func (m *Money) GreaterThan(originalMoney *Money) (bool, error) {
	if err := m.AssertSameCurrency(originalMoney); err != nil {
		return false, err
	}
	return m.Compare(originalMoney) == 1, nil
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(originalMoney *Money) (bool, error) {
	if err := m.AssertSameCurrency(originalMoney); err != nil {
		return false, err
	}

	return m.Compare(originalMoney) >= 0, nil
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(originalMoney *Money) (bool, error) {
	if err := m.AssertSameCurrency(originalMoney); err != nil {
		return false, err
	}

	return m.Compare(originalMoney) == -1, nil
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(originalMoney *Money) (bool, error) {
	if err := m.AssertSameCurrency(originalMoney); err != nil {
		return false, err
	}

	return m.Compare(originalMoney) <= 0, nil
}

// IsZero returns boolean of whether the value of Money is equals to zero.
func (m *Money) IsZero() bool {
	return m.Amount == 0
}

// IsPositive returns boolean of whether the value of Money is positive.
func (m *Money) IsPositive() bool {
	return m.Amount > 0
}

// IsNegative returns boolean of whether the value of Money is negative.
func (m *Money) IsNegative() bool {
	return m.Amount < 0
}
