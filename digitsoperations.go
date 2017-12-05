// digitsoperations
package projecteuler

import "strconv"
import "strings"

func isNumberPalindrome(n uint64) bool {
	d := getDigitsFromNumber(n)
	for i, j := 0, len(d); i < len(d)/2; i++ {
		if d[i] != d[j-i-1] {
			return false
		}
	}
	return true
}

func getDigitsFromNumber(n uint64) []uint64 {
	digits := make([]uint64, 0, 10)
	var base uint64 = 10
	for n > 0 {
		r := n % base
		n -= r
		n /= base
		rs := make([]uint64, 0, 1)
		rs = append(rs, r)
		digits = append(rs, digits...)
		rs = nil
	}
	return digits
}

type FloatNumberData struct {
	decimalPartDigits  []uint64
	floatingPartDigits []uint64
	exponent           uint64
}

//TODO Locale dependent function
//Parses float in example form 9,33262154439441e+157 zou can use . instead of ,
//Converts float to string, then back, returns decimalParts digits i.e digits before .
//Fractional part digits i.e digits after .
//Exponent i.e number after e+ or E+
func getDigitsFromFloatNumber(n float64) (FloatNumberData, error) {
	var data FloatNumberData
	s := strconv.FormatFloat(n, 'e', -1, 64)
	sub1 := strings.FieldsFunc(s, fieldFloatSeparator)
	number, err := strconv.ParseUint(sub1[0], 10, 64)
	if err != nil {
		return data, err
	}

	digits := getDigitsFromNumber(number)
	data.decimalPartDigits = digits

	sub2 := strings.FieldsFunc(sub1[1], fieldExponentSeparator)
	number, err = strconv.ParseUint(sub2[0], 10, 64)
	if err != nil {
		return data, err
	}
	digits = getDigitsFromNumber(number)
	data.floatingPartDigits = digits

	//this is ugly, + or - is right after 'e' or 'E'
	sub3 := strings.FieldsFunc(sub2[1], fieldSignSeparator)
	exponent, err := strconv.ParseUint(sub3[0], 10, 64)
	if err != nil {
		return data, err
	}
	data.exponent = exponent

	return data, nil
}

func fieldFloatSeparator(r rune) bool {
	if r == ',' || r == '.' {
		return true
	}
	return false
}

func fieldExponentSeparator(r rune) bool {
	if r == 'e' || r == 'E' {
		return true
	}
	return false
}

func fieldSignSeparator(r rune) bool {
	if r == '+' || r == '-' {
		return true
	}
	return false
}

func getFloatString(n float64) string {
	s := strconv.FormatFloat(n, 'e', -1, 64)
	return s
}
