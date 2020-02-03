package bases

import (
	"bases/src/errors"
	"math"
	"strconv"
	"strings"
)

func ConvertBases(number string, actualBase, destinationBase int) (string, error){

	//Validate symbols
	err := ValidateSymbols(number, actualBase)
	if err != nil {
		return "", err
	}

	//Convert the number to base 10
	decimal := ToDecimal(number, actualBase)

	var quotient, base float64
	base = float64(destinationBase)
	quotient = decimal

	var residues []int

	//Get the first residue
	residue := int(quotient) % int(base)
	residues = append(residues, residue)

	//Get all the residues while the quotient be
	//equals or more than the actual base.
	for quotient >= base {
		quotient = quotient/base
		residue := int(quotient) % int(base)
		residues = append(residues, residue)
	}

	//Apply reverse to array of residues
	reverse := reverseArray(residues)

	//add sub index of the resultant number
	return reverse + SubIndex(destinationBase), nil
}

func ToDecimal( number string, base int ) float64 {
	number = reverse(number)
	array := strings.Split(number, "")

	var result int

	for i, v := range array{
		intValue := NumericValue(v)
		pow := intValue * int(math.Pow( float64(base), float64(i)))
		result += pow
	}

	return float64(result)
}

func ValidateSymbols(number string, actualBase int) error {
	symbols := []string {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A",
		"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
		"X","Y","Z",
	}

	validSymbols := symbols[0:actualBase]

	mapSymbols := make(map[string]int, len(validSymbols))
	for _, v := range validSymbols {
		mapSymbols[v] = 0
	}

	number = strings.ToUpper(number)

	for _, v := range number {
		_, ok := mapSymbols[string(v)]
		if !ok {
			return &errors.BasesError{
				Number: number,
				Base: actualBase,
				ValidSymbols: validSymbols,
			}
		}
	}

	return nil
}

// https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseArray(a []int) string {
	max := len(a)
	reverseString := ""

	for i := max - 1; i >= 0; i-- {
		var actualValue string
		if a[i] > 9 {
			actualValue = AlphabeticValue(a[i])
		} else {
			actualValue = strconv.Itoa(a[i])
		}
		reverseString = reverseString + actualValue
	}

	return reverseString
}