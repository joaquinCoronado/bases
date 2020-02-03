package bases

import (
	"strconv"
	"strings"
)

var bases = map[int]string {
	1:  "₁",
	2:  "₂",
	3:  "₃",
	4:  "₄",
	5:  "₅",
	6:  "₆",
	7:  "₇",
	8:  "₈",
	9:  "₉",
	10: "₁₀",
	11: "₁₁",
	12: "₁₂",
	13: "₁₃",
	14: "₁₄",
	15: "₁₅",
	16: "₁₆",
	17: "₁₇",
	18: "₁₈",
	19: "₁₉",
	20: "₂₀",
	21: "₂₁",
	22: "₂₂",
	23: "₂₃",
	24: "₂₄",
	25: "₂₅",
	26: "₂₆",
	27: "₂₇",
	28: "₂₈",
	29: "₂₉",
	30: "₃₀",
}

func SubIndex(base int) string{
	return bases[base]
}

var alphabetic = map[int]string{
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
	16: "G",
	17: "H",
	18: "I",
	19: "J",
	20: "K",
	21: "L",
	22: "M",
	23: "N",
	24: "O",
	25: "P",
	26: "Q",
	27: "R",
	28: "S",
	29: "T",
	30: "U",
	31: "V",
	32: "W",
	33: "X",
	34: "Y",
	35: "Z",
}

func AlphabeticValue(number int) string {
	return alphabetic[number]
}

var numeric = map[string]int{
	 "A": 10,
	 "B": 11,
	 "C": 12,
	 "D": 13,
	 "E": 14,
	 "F": 15,
	 "G": 16,
	 "H": 17,
	 "I": 18,
	 "J": 19,
	 "K": 20,
	 "L": 21,
	 "M": 22,
	 "N": 23,
	 "O": 24,
	 "P": 25,
	 "Q": 26,
	 "R": 27,
	 "S": 28,
	 "T": 29,
	 "U": 30,
	 "V": 31,
	 "W": 32,
	 "X": 33,
	 "Y": 34,
	 "Z": 35,
}

func NumericValue(value string) int {
	value = strings.ToUpper(value)
	response := numeric[value]

	if response == 0 {
		test, _ := strconv.Atoi(value)
		return test
	}
	return response
}
