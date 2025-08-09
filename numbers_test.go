package numbers

import (
	"testing"
)

func TestTextToNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"zero", 0},
		{"um", 1},
		{"dez", 10},
		{"vinte e cinco", 25},
		{"cem", 100},
		{"cento e um", 101},
		{"mil", 1000},
		{"dois mil e vinte", 2020},
		{"quinhentos e vinte e três mil e onze", 523011},
	}
	for _, tt := range tests {
		num, err := TextToNumber(tt.input)
		if err != nil {
			t.Errorf("TextToNumber(%q) returned error: %v", tt.input, err)
		}
		if num != tt.expected {
			t.Errorf("TextToNumber(%q) = %d; want %d", tt.input, num, tt.expected)
		}
	}
}

func TestNumberToText(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "zero"},
		{1, "um"},
		{10, "dez"},
		{25, "vinte e cinco"},
		{100, "cem"},
		{101, "cento e um"},
		{1000, "mil"},
		{2020, "dois mil e vinte"},
		{523011, "quinhentos e vinte e três mil e onze"},
	}
	for _, tt := range tests {
		text, err := NumberToText(tt.input)
		if err != nil {
			t.Errorf("NumberToText(%d) returned error: %v", tt.input, err)
		}
		if text != tt.expected {
			t.Errorf("NumberToText(%d) = %q; want %q", tt.input, text, tt.expected)
		}
	}
}

func TestNumberToTextErrors(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		wantErr string
	}{
		{
			name:    "negative number",
			input:   -1,
			wantErr: "Cannot write negative numbers",
		},
		{
			name:    "negative large number",
			input:   -1000,
			wantErr: "Cannot write negative numbers",
		},
		{
			name:    "number exceeds maximum",
			input:   1000000000000000000, // exceeds maxNumber (999999999999999999)
			wantErr: "Number exceeds maximum allowed value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NumberToText(tt.input)
			if err == nil {
				t.Errorf("NumberToText(%d) expected error, but got none", tt.input)
				return
			}
			if err.Error() != tt.wantErr {
				t.Errorf("NumberToText(%d) error = %q; want %q", tt.input, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestTextToNumberErrors(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr string
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: "empty text",
		},
		{
			name:    "only spaces",
			input:   "   ",
			wantErr: "empty text",
		},
		{
			name:    "unrecognized word",
			input:   "banana",
			wantErr: "unrecognized word: banana",
		},
		{
			name:    "mixed valid and invalid words",
			input:   "dez bananas",
			wantErr: "unrecognized word: bananas",
		},
		{
			name:    "number with invalid word",
			input:   "vinte e banana cinco",
			wantErr: "unrecognized word: banana",
		},
		{
			name:    "english number words",
			input:   "twenty five",
			wantErr: "unrecognized word: twenty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TextToNumber(tt.input)
			if err == nil {
				t.Errorf("TextToNumber(%q) expected error, but got none", tt.input)
				return
			}
			if err.Error() != tt.wantErr {
				t.Errorf("TextToNumber(%q) error = %q; want %q", tt.input, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("maximum allowed number", func(t *testing.T) {
		maxNum := int64(999999999999999999)
		text, err := NumberToText(maxNum)
		if err != nil {
			t.Errorf("NumberToText(%d) unexpected error: %v", maxNum, err)
		}
		if text == "" {
			t.Errorf("NumberToText(%d) returned empty string", maxNum)
		}
	})

	t.Run("zero boundary", func(t *testing.T) {
		text, err := NumberToText(0)
		if err != nil {
			t.Errorf("NumberToText(0) unexpected error: %v", err)
		}
		if text != "zero" {
			t.Errorf("NumberToText(0) = %q; want %q", text, "zero")
		}
	})

	t.Run("text with accents", func(t *testing.T) {
		num, err := TextToNumber("três")
		if err != nil {
			t.Errorf("TextToNumber('três') unexpected error: %v", err)
		}
		if num != 3 {
			t.Errorf("TextToNumber('três') = %d; want %d", num, 3)
		}
	})
}

func TestTextToDecimal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected DecimalResult
	}{
		{
			name:  "monetary value with reais and centavos",
			input: "cinquenta reais e vinte e cinco centavos",
			expected: DecimalResult{
				Integer:    50,
				Fractional: 25,
				Decimals:   2,
			},
		},
		{
			name:  "decimal with vírgula",
			input: "cinquenta e dois vírgula vinte e cinco",
			expected: DecimalResult{
				Integer:    52,
				Fractional: 25,
				Decimals:   2,
			},
		},
		{
			name:  "decimal with ponto",
			input: "três ponto quatorze",
			expected: DecimalResult{
				Integer:    3,
				Fractional: 14,
				Decimals:   2,
			},
		},
		{
			name:  "only reais",
			input: "cem reais",
			expected: DecimalResult{
				Integer:    100,
				Fractional: 0,
				Decimals:   2,
			},
		},
		{
			name:  "only centavos",
			input: "cinquenta centavos",
			expected: DecimalResult{
				Integer:    0,
				Fractional: 50,
				Decimals:   2,
			},
		},
		{
			name:  "integer only",
			input: "quarenta e dois",
			expected: DecimalResult{
				Integer:    42,
				Fractional: 0,
				Decimals:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TextToDecimal(tt.input)
			if err != nil {
				t.Errorf("TextToDecimal(%q) returned error: %v", tt.input, err)
				return
			}

			if result.Integer != tt.expected.Integer {
				t.Errorf("TextToDecimal(%q) Integer = %d; want %d", tt.input, result.Integer, tt.expected.Integer)
			}
			if result.Fractional != tt.expected.Fractional {
				t.Errorf("TextToDecimal(%q) Fractional = %d; want %d", tt.input, result.Fractional, tt.expected.Fractional)
			}
			if result.Decimals != tt.expected.Decimals {
				t.Errorf("TextToDecimal(%q) Decimals = %d; want %d", tt.input, result.Decimals, tt.expected.Decimals)
			}
		})
	}
}

func TestDecimalToText(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		monetary bool
		expected string
	}{
		{
			name:     "monetary value",
			value:    50.25,
			monetary: true,
			expected: "cinquenta reais e vinte e cinco centavos",
		},
		{
			name:     "one real",
			value:    1.00,
			monetary: true,
			expected: "um real",
		},
		{
			name:     "only centavos",
			value:    0.99,
			monetary: true,
			expected: "noventa e nove centavos",
		},
		{
			name:     "one centavo",
			value:    0.01,
			monetary: true,
			expected: "um centavo",
		},
		{
			name:     "decimal with vírgula",
			value:    52.25,
			monetary: false,
			expected: "cinquenta e dois vírgula vinte e cinco",
		},
		{
			name:     "decimal starting with zero",
			value:    0.75,
			monetary: false,
			expected: "zero vírgula setenta e cinco",
		},
		{
			name:     "integer only",
			value:    42.0,
			monetary: false,
			expected: "quarenta e dois",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecimalToText(tt.value, tt.monetary)
			if err != nil {
				t.Errorf("DecimalToText(%f, %t) returned error: %v", tt.value, tt.monetary, err)
				return
			}

			if result != tt.expected {
				t.Errorf("DecimalToText(%f, %t) = %q; want %q", tt.value, tt.monetary, result, tt.expected)
			}
		})
	}
}

func TestDecimalResult(t *testing.T) {
	t.Run("ToFloat64", func(t *testing.T) {
		tests := []struct {
			name     string
			decimal  DecimalResult
			expected float64
		}{
			{
				name: "monetary value",
				decimal: DecimalResult{
					Integer:    50,
					Fractional: 25,
					Decimals:   2,
				},
				expected: 50.25,
			},
			{
				name: "integer only",
				decimal: DecimalResult{
					Integer:    42,
					Fractional: 0,
					Decimals:   0,
				},
				expected: 42.0,
			},
			{
				name: "three decimal places",
				decimal: DecimalResult{
					Integer:    3,
					Fractional: 145,
					Decimals:   3,
				},
				expected: 3.145,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.decimal.ToFloat64()
				if result != tt.expected {
					t.Errorf("ToFloat64() = %f; want %f", result, tt.expected)
				}
			})
		}
	})

	t.Run("ToString", func(t *testing.T) {
		tests := []struct {
			name     string
			decimal  DecimalResult
			expected string
		}{
			{
				name: "monetary value",
				decimal: DecimalResult{
					Integer:    50,
					Fractional: 25,
					Decimals:   2,
				},
				expected: "50.25",
			},
			{
				name: "integer only",
				decimal: DecimalResult{
					Integer:    42,
					Fractional: 0,
					Decimals:   0,
				},
				expected: "42",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.decimal.ToString()
				if result != tt.expected {
					t.Errorf("ToString() = %q; want %q", result, tt.expected)
				}
			})
		}
	})
}
