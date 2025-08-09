package numbers

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const maxNumber int64 = 999999999999999999
const thousandSize int = 3

func NumberToText(n int64) (string, error) {

	err := checkMaxNumber(n)
	if err != "" {
		return "", errors.New(err)
	}

	err = checkNegativeNumber(n)
	if err != "" {
		return "", errors.New(err)
	}

	if n == 0 {
		return "zero", nil
	}

	if n == 1000 {
		return "mil", nil
	}

	result := ""
	numberStr := strconv.FormatInt(n, 10)
	length := len(numberStr)

	groups := int(math.Ceil(float64(length) / float64(thousandSize)))

	numberStr = leftPad(numberStr, "0", groups*3)

	for i := 0; i < groups; i++ {
		groupValue, _ := strconv.ParseInt(numberStr[i*3:3+i*3], 10, 0)

		if groupValue == 0 {
			continue
		}

		hundred := int(math.Floor(float64(groupValue / 100)))
		ten := int(math.Floor(float64((groupValue - int64(hundred*100)) / 10)))
		unit := int(groupValue - int64(hundred*100) - int64(ten*10))

		if hundred > 0 {
			if ten == 0 && unit == 0 && result != "" {
				result += " e "
			} else if result != "" {
				result += ", "
			}
			if hundred == 1 {
				if ten > 0 || unit > 0 {
					result += hundreds[hundred].plural
				} else {
					result += hundreds[hundred].singular
				}
			} else {
				result += hundreds[hundred].singular
			}
		}

		if ten > 0 {
			if result != "" {
				result += " e "
			}

			if ten == 1 {
				ten = 10 + unit
				unit = 0
				result += units[ten]
			} else {
				result += tens[ten]
			}
		}

		if unit > 0 {
			if result != "" {
				result += " e "
			}
			result += units[unit]
		}

		if groupValue > 1 {
			result += " " + thousands[groups-i-1].plural
		} else {
			result += " " + thousands[groups-i-1].singular
		}
	}

	// At the end, remove extra spaces and commas
	result = strings.TrimSpace(result)
	result = strings.TrimSuffix(result, ",")
	result = strings.TrimSuffix(result, " ")

	return result, nil
}

func DecimalToText(value float64, monetary bool) (string, error) {
	if value < 0 {
		return "", errors.New("Cannot write negative decimal numbers")
	}

	integerPart := int64(value)
	fractionalPart := value - float64(integerPart)

	integerText, err := NumberToText(integerPart)
	if err != nil {
		return "", err
	}

	if fractionalPart == 0 {
		if monetary {
			if integerPart == 1 {
				return integerText + " real", nil
			}
			return integerText + " reais", nil
		}
		return integerText, nil
	}

	if monetary {
		centavos := int64(fractionalPart*100 + 0.5)

		if centavos == 0 {
			if integerPart == 1 {
				return integerText + " real", nil
			}
			return integerText + " reais", nil
		}

		centavosText, err := NumberToText(centavos)
		if err != nil {
			return "", err
		}

		var reaisText string
		if integerPart == 0 {
			reaisText = ""
		} else if integerPart == 1 {
			reaisText = integerText + " real e "
		} else {
			reaisText = integerText + " reais e "
		}

		var centavosUnit string
		if centavos == 1 {
			centavosUnit = " centavo"
		} else {
			centavosUnit = " centavos"
		}

		return reaisText + centavosText + centavosUnit, nil
	} else {
		fractionalStr := fmt.Sprintf("%.10f", fractionalPart)[2:] // Remove "0."
		fractionalStr = strings.TrimRight(fractionalStr, "0")     // Remove trailing zeros

		if len(fractionalStr) == 0 {
			return integerText, nil
		}

		fractionalInt := int64(0)
		for _, digit := range fractionalStr {
			if digit >= '0' && digit <= '9' {
				fractionalInt = fractionalInt*10 + int64(digit-'0')
			}
		}

		fractionalText, err := NumberToText(fractionalInt)
		if err != nil {
			return "", err
		}

		if integerPart == 0 {
			return "zero vírgula " + fractionalText, nil
		}

		return integerText + " vírgula " + fractionalText, nil
	}
}

func checkMaxNumber(n int64) string {
	if n > maxNumber {
		return "Number exceeds maximum allowed value"
	}
	return ""
}

func checkNegativeNumber(n int64) string {
	if n < 0 {
		return "Cannot write negative numbers"
	}
	return ""
}
