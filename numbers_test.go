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
