package numbers

import (
	"errors"
	"strings"
)

func TextToNumber(s string) (int64, error) {
	s = normalize(s)
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0, errors.New("texto vazio")
	}

	unitsInv := invertMap(units)
	tensInv := invertMap(tens)

	hundredsInv := map[string]int{
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

	thousandsInv := map[string]int64{
		"mil":         1_000,
		"milhao":      1_000_000,
		"milhoes":     1_000_000,
		"bilhao":      1_000_000_000,
		"bilhoes":     1_000_000_000,
		"trilhao":     1_000_000_000_000,
		"trilhoes":    1_000_000_000_000,
		"quatrilhao":  1_000_000_000_000_000,
		"quatrilhoes": 1_000_000_000_000_000,
	}

	var total, partial int64

	for _, w := range words {
		if val, ok := unitsInv[w]; ok {
			partial += int64(val)
		} else if val, ok := tensInv[w]; ok {
			partial += int64(val)
		} else if val, ok := hundredsInv[w]; ok {
			partial += int64(val)
		} else if mult, ok := thousandsInv[w]; ok {
			if partial == 0 {
				partial = 1
			}
			total += partial * mult
			partial = 0
		} else if w == "e" {
			// ignora o 'e'
		} else {
			return 0, errors.New("palavra não reconhecida: " + w)
		}
	}

	total += partial
	return total, nil
}
