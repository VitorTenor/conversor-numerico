package numbers

import (
	"errors"
	"fmt"
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

		if word == "e" {
			i++
			continue
		}

		if val, ok := unitsInv[word]; ok {
			currentSum += val
			i++
			continue
		}

		if val, ok := tensInv[word]; ok {
			currentSum += val
			i++
			continue
		}

		if val, ok := hundredsInv[word]; ok {
			currentSum += val
			i++
			continue
		}

		if mult, ok := multipliers[word]; ok {

			if currentSum == 0 {
				currentSum = 1
			}

			currentSum *= mult

			if mult == 1000 && i < len(words)-1 {
				result += currentSum
				currentSum = 0
			} else {

				result += currentSum
				currentSum = 0
			}
			i++
			continue
		}

		return 0, errors.New("unrecognized word: " + word)
	}

	result += currentSum

	return result, nil
}

type DecimalResult struct {
	Integer    int64
	Fractional int64
	Decimals   int
}

func TextToDecimal(s string) (DecimalResult, error) {
	s = normalize(s)
	words := strings.Fields(s)
	if len(words) == 0 {
		return DecimalResult{}, errors.New("empty text")
	}

	if isMonetaryExpression(words) {
		return parseMonetaryExpression(words)
	}

	if hasDecimalSeparator(words) {
		return parseDecimalExpression(words)
	}

	intValue, err := TextToNumber(s)
	if err != nil {
		return DecimalResult{}, err
	}

	return DecimalResult{Integer: intValue, Fractional: 0, Decimals: 0}, nil
}

func isMonetaryExpression(words []string) bool {
	for _, word := range words {
		if currencyWords[word] {
			return true
		}
	}
	return false
}

func hasDecimalSeparator(words []string) bool {
	for _, word := range words {
		if word == "virgula" || word == "vírgula" || word == "ponto" {
			return true
		}
	}
	return false
}

func parseMonetaryExpression(words []string) (DecimalResult, error) {
	var realsPart []string
	var centavosPart []string
	var inCentavos bool

	for i, word := range words {
		if word == "real" || word == "reais" {
			continue
		} else if word == "centavo" || word == "centavos" {
			inCentavos = false
			continue
		} else if (word == "real" || word == "reais") && i < len(words)-1 {
			if i+1 < len(words) && words[i+1] == "e" {
				inCentavos = true
				continue
			}
		} else if inCentavos {
			if word != "e" {
				centavosPart = append(centavosPart, word)
			}
		} else {
			if word != "e" || len(realsPart) == 0 {
				realsPart = append(realsPart, word)
			}
		}
	}

	realsEndIndex := -1
	centavosStartIndex := -1

	for i, word := range words {
		if word == "real" || word == "reais" {
			realsEndIndex = i
		}
		if (word == "centavo" || word == "centavos") && centavosStartIndex == -1 {
			for j := realsEndIndex + 1; j < i; j++ {
				if words[j] != "e" {
					centavosStartIndex = j
					break
				}
			}
			break
		}
	}

	var reaisValue int64 = 0
	var err error
	if realsEndIndex > 0 {
		reaisPart := strings.Join(words[:realsEndIndex], " ")
		reaisValue, err = TextToNumber(reaisPart)
		if err != nil {
			return DecimalResult{}, err
		}
	}

	var centavosValue int64 = 0
	if centavosStartIndex != -1 {
		centavosEnd := len(words)
		for i := centavosStartIndex; i < len(words); i++ {
			if words[i] == "centavo" || words[i] == "centavos" {
				centavosEnd = i
				break
			}
		}
		if centavosStartIndex < centavosEnd {
			centavosPart := strings.Join(words[centavosStartIndex:centavosEnd], " ")
			centavosValue, err = TextToNumber(centavosPart)
			if err != nil {
				return DecimalResult{}, err
			}
		}
	}

	return DecimalResult{
		Integer:    reaisValue,
		Fractional: centavosValue,
		Decimals:   2,
	}, nil
}

func parseDecimalExpression(words []string) (DecimalResult, error) {
	separatorIndex := -1

	for i, word := range words {
		if word == "virgula" || word == "vírgula" || word == "ponto" {
			separatorIndex = i
			break
		}
	}

	if separatorIndex == -1 {
		return DecimalResult{}, errors.New("decimal separator not found")
	}

	var integerValue int64 = 0
	var err error
	if separatorIndex > 0 {
		integerPart := strings.Join(words[:separatorIndex], " ")
		integerValue, err = TextToNumber(integerPart)
		if err != nil {
			return DecimalResult{}, err
		}
	}

	var fractionalValue int64 = 0
	decimals := 0
	if separatorIndex < len(words)-1 {
		fractionalPart := strings.Join(words[separatorIndex+1:], " ")
		fractionalValue, err = TextToNumber(fractionalPart)
		if err != nil {
			return DecimalResult{}, err
		}

		temp := fractionalValue
		if temp == 0 {
			decimals = 1
		} else {
			for temp > 0 {
				decimals++
				temp /= 10
			}
		}
	}

	return DecimalResult{
		Integer:    integerValue,
		Fractional: fractionalValue,
		Decimals:   decimals,
	}, nil
}

func (d DecimalResult) ToFloat64() float64 {
	if d.Decimals == 0 {
		return float64(d.Integer)
	}

	divisor := 1.0
	for i := 0; i < d.Decimals; i++ {
		divisor *= 10
	}

	return float64(d.Integer) + float64(d.Fractional)/divisor
}

func (d DecimalResult) ToString() string {
	if d.Decimals == 0 {
		return fmt.Sprintf("%d", d.Integer)
	}

	format := fmt.Sprintf("%%.%df", d.Decimals)
	return fmt.Sprintf(format, d.ToFloat64())
}
