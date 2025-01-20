package tests

import (
	"github.com/gouef/currency"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCurrency(t *testing.T) {
	tests := []struct {
		code     string
		expected *currency.Currency
		hasError bool
	}{
		{
			code: "USD",
			expected: &currency.Currency{
				Code:       "USD",
				Name:       "United States Dollar",
				Symbol:     "$",
				MinorUnits: 2,
			},
			hasError: false,
		},
		{
			code: "EUR",
			expected: &currency.Currency{
				Code:       "EUR",
				Name:       "Euro",
				Symbol:     "€",
				MinorUnits: 2,
			},
			hasError: false,
		},
		{
			code:     "XYZ", // neexistující měna
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {
			result, err := currency.FindCurrency(test.code)
			if test.hasError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestValidateCurrency(t *testing.T) {
	tests := []struct {
		code     string
		expected bool
	}{
		{code: "USD", expected: true},
		{code: "EUR", expected: true},
		{code: "XYZ", expected: false}, // neexistující měna
	}

	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {
			result := currency.ValidateCurrency(test.code)
			assert.Equal(t, test.expected, result)
		})
	}
}
