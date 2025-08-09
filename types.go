package numbers

type word struct {
	singular string
	plural   string
}

var units = map[int]string{
	1:  "um",
	2:  "dois",
	3:  "três",
	4:  "quatro",
	5:  "cinco",
	6:  "seis",
	7:  "sete",
	8:  "oito",
	9:  "nove",
	10: "dez",
	11: "onze",
	12: "doze",
	13: "treze",
	14: "quatorze",
	15: "quinze",
	16: "dezesseis",
	17: "dezessete",
	18: "dezoito",
	19: "dezenove",
}

var tens = map[int]string{
	2: "vinte",
	3: "trinta",
	4: "quarenta",
	5: "cinquenta",
	6: "sessenta",
	7: "setenta",
	8: "oitenta",
	9: "noventa",
}

var hundreds = map[int]word{
	1: {"cem", "cento"},
	2: {"duzentos", "duzentos"},
	3: {"trezentos", "trezentos"},
	4: {"quatrocentos", "quatrocentos"},
	5: {"quinhentos", "quinhentos"},
	6: {"seiscentos", "seiscentos"},
	7: {"setecentos", "setecentos"},
	8: {"oitocentos", "oitocentos"},
	9: {"novecentos", "novecentos"},
}

var thousands = []word{
	{"", ""},
	{"mil", "mil"},
	{"milhão", "milhões"},
	{"bilhão", "bilhões"},
	{"trilhão", "trilhões"},
	{"quatrilhão", "quatrilhões"},
}
