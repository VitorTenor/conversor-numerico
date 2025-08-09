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
