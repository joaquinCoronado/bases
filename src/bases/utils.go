package bases

import (
	"math"
	"strconv"
	"strings"
)

func ConvertBases(number string, actualBase, destinationBase int) string{

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
	return reverse + SubIndex(destinationBase)
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