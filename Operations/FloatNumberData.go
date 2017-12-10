// FloatNumberData
package projecteuler

import (
	"strconv"
	"strings"
)

type FloatNumberData struct {
	decimalPart          uint64
	floatingPart         []uint64
	exponent             uint64
	stringRepresentation string
}

//TODO Locale dependent function
//Parses float in example form 9,33262154439441e+157so  you can use . instead of ,
//Converts float to string, then back, returns decimalParts digits i.e digits before .
//Fractional part digits i.e digits after .
//Exponent i.e number after e+ or E+
//Ugly function
func NewFloatNumberData(n float64) (FloatNumberData, error) {
	var data FloatNumberData
	s := getFloatString(n)
	data.stringRepresentation = s
	decPart := strings.FieldsFunc(s, fieldFloatSeparator)
	number, err := strconv.ParseUint(decPart[0], 10, 64)
	if err != nil {
		return data, err
	}
	data.decimalPart = number

	floatingPointPart := strings.FieldsFunc(decPart[1], fieldExponentSeparator)
	number, err = strconv.ParseUint(floatingPointPart[0], 10, 64)
	if err != nil {
		return data, err
	}
	digits := GetDigitsFromNumber(number)
	data.floatingPart = digits

	//this is ugly, + or - is right after 'e' or 'E', FieldFunc does not return it,
	//because
	exponentPart := strings.FieldsFunc(floatingPointPart[1], fieldSignSeparator)
	exponent, err := strconv.ParseUint(exponentPart[0], 10, 64)
	if err != nil {
		return data, err
	}
	data.exponent = exponent

	return data, nil
}

//Compare two FloatNumberData, significant means how many integers from floating point part
//are compared.s
func (d *FloatNumberData) isEqual(data *FloatNumberData, significant int) bool {
	b := significant
	fp1 := len(d.floatingPart)
	fp2 := len(data.floatingPart)
	if significant > fp1 || significant > fp2 {
		l := smallerOrEqual(fp1, fp2)
		b = l
	}

	if d.decimalPart == data.decimalPart && areSlicesEqual(data.floatingPart[:b], data.floatingPart[:b]) && d.exponent == data.exponent {
		return true
	}
	return false
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
