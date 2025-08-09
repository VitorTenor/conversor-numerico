package numbers

import (
	"strings"
	"unicode"
)

func leftPad(str, pad string, length int) string {
	for {
		if len(str) < length {
			str = pad + str
		} else {
			return str
		}
	}
}

func normalize(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		switch r {
		case 'á', 'à', 'â', 'ã':
			r = 'a'
		case 'é', 'è', 'ê':
			r = 'e'
		case 'í', 'ì', 'î':
			r = 'i'
		case 'ó', 'ò', 'ô', 'õ':
			r = 'o'
		case 'ú', 'ù', 'û':
			r = 'u'
		case 'ç':
			r = 'c'
		}
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func invertMap(m map[int]string) map[string]int {
	out := make(map[string]int)
	for k, v := range m {
		out[normalize(v)] = k
	}
	return out
}
