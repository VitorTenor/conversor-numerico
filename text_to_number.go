package numbers

import (
	"errors"
	"strings"
)

func TextToNumber(s string) (int64, error) {
	s = normalize(s)
	if s == "zero" {
		return 0, nil
	}

	words := strings.Fields(s)
	if len(words) == 0 {
		return 0, errors.New("empty text")
	}

	// Maps for conversion
	unitsInv := map[string]int64{
		"um":        1,
		"dois":      2,
		"três":      3,
		"tres":      3,
		"quatro":    4,
		"cinco":     5,
		"seis":      6,
		"sete":      7,
		"oito":      8,
		"nove":      9,
		"dez":       10,
		"onze":      11,
		"doze":      12,
		"treze":     13,
		"quatorze":  14,
		"quinze":    15,
		"dezesseis": 16,
		"dezessete": 17,
		"dezoito":   18,
		"dezenove":  19,
	}

	tensInv := map[string]int64{
		"vinte":     20,
		"trinta":    30,
		"quarenta":  40,
		"cinquenta": 50,
		"sessenta":  60,
		"setenta":   70,
		"oitenta":   80,
		"noventa":   90,
	}

	hundredsInv := map[string]int64{
		"cem":          100,
		"cento":        100,
		"duzentos":     200,
		"trezentos":    300,
		"quatrocentos": 400,
		"quinhentos":   500,
		"seiscentos":   600,
		"setecentos":   700,
		"oitocentos":   800,
		"novecentos":   900,
	}

	multipliers := map[string]int64{
		"mil":         1000,
		"milhao":      1000000,
		"milhões":     1000000,
		"milhoes":     1000000,
		"bilhao":      1000000000,
		"bilhões":     1000000000,
		"bilhoes":     1000000000,
		"trilhao":     1000000000000,
		"trilhões":    1000000000000,
		"trilhoes":    1000000000000,
		"quatrilhao":  1000000000000000,
		"quatrilhões": 1000000000000000,
		"quatrilhoes": 1000000000000000,
	}

	var result int64 = 0
	var currentSum int64 = 0

	i := 0
	for i < len(words) {
		word := words[i]

		// Ignore conjunctions
		if word == "e" {
			i++
			continue
		}

		// Check if it's a unit
		if val, ok := unitsInv[word]; ok {
			currentSum += val
			i++
			continue
		}

		// Check if it's a ten
		if val, ok := tensInv[word]; ok {
			currentSum += val
			i++
			continue
		}

		// Check if it's a hundred
		if val, ok := hundredsInv[word]; ok {
			currentSum += val
			i++
			continue
		}

		// Check if it's a multiplier
		if mult, ok := multipliers[word]; ok {
			// If no accumulated value, consider as 1
			if currentSum == 0 {
				currentSum = 1
			}

			// Multiply the accumulated value
			currentSum *= mult

			// If it's "mil" and not the end of phrase, store partial result
			if mult == 1000 && i < len(words)-1 {
				result += currentSum
				currentSum = 0
			} else {
				// For other multipliers or end of phrase, add to result
				result += currentSum
				currentSum = 0
			}
			i++
			continue
		}

		return 0, errors.New("unrecognized word: " + word)
	}

	// Add any remaining value
	result += currentSum

	return result, nil
}
