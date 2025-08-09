package numbers

import (
	"errors"
	"math"
	"strconv"
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

	return result, nil
}

func checkMaxNumber(n int64) string {
	if n > maxNumber {
		return "Número informado maior que o máximo permitido"
	}
	return ""
}

func checkNegativeNumber(n int64) string {
	if n < 0 {
		return "Não é possível escrever números negativos"
	}
	return ""
}
