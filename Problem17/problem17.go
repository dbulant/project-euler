// problem17
package projecteuler

import . "github.com/dbulant/project-euler/Operations"

//If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//NOTE: Do not count spaces or hyphens.
//For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
func Problem17() {
	GeneralProblem17(1000)
}

func GeneralProblem17(max uint64) uint64 {
	initMap()
	letters := make([]string, 0, max)
	for i := uint64(1); i <= max; i++ {
		letters = append(letters, NumberToLetters(i)...)
	}
	tc := uint64(0)
	for _, w := range letters {
		tc += CountWithoutHyphensSpaces(w)
	}
	return tc
}

var NumberStringMap map[uint64]string

func initMap() {
	NumberStringMap = make(map[uint64]string)
	NumberStringMap[1] = "one"
	NumberStringMap[2] = "two"
	NumberStringMap[3] = "three"
	NumberStringMap[4] = "four"
	NumberStringMap[5] = "five"
	NumberStringMap[6] = "six"
	NumberStringMap[7] = "seven"
	NumberStringMap[8] = "eight"
	NumberStringMap[9] = "nine"
	NumberStringMap[10] = "ten"
	NumberStringMap[0] = "zero"

	NumberStringMap[20] = "twenty"
	NumberStringMap[30] = "thirty"
	NumberStringMap[40] = "forty"
	NumberStringMap[50] = "fifty"
	NumberStringMap[60] = "sixty"
	NumberStringMap[70] = "seventy"
	NumberStringMap[80] = "eighty"
	NumberStringMap[90] = "ninety"

	NumberStringMap[100] = "one hundred"
	NumberStringMap[200] = "two hundred"
	NumberStringMap[300] = "three hundred"
	NumberStringMap[400] = "four hundred"
	NumberStringMap[500] = "five hundred"
	NumberStringMap[600] = "six hundred"
	NumberStringMap[700] = "seven hundred"
	NumberStringMap[800] = "eight hundred"
	NumberStringMap[900] = "nine hundred"
	NumberStringMap[1000] = "one thousand"
}

//Works only till 1000.
func NumberToLetters(number uint64) []string {
	digits := GetDigitsFromNumber(number)
	lend := len(digits)
	letters := make([]string, 0, 1)
	if lend == 1 {
		n := digits[0]
		letters = append(letters, NumberStringMap[n])
	} else if lend == 2 {
		n2 := 10 * digits[0]
		n1 := digits[1]
		s := NumberStringMap[n2] + " and " + NumberStringMap[n1]
		letters = append(letters, s)
	} else if lend == 3 {
		n3 := 100 * digits[0]
		n2 := 10 * digits[1]
		n1 := digits[2]
		s := NumberStringMap[n3] + " and " + NumberStringMap[n2] + "-" + NumberStringMap[n1]
		letters = append(letters, s)
	}
	return letters
}

//Counts number of letters, spaces and hyphens are excluded
func CountWithoutHyphensSpaces(letters string) uint64 {
	count := uint64(0)
	for _, l := range letters {
		if l != ' ' && l != '-' {
			count++
		}
	}
	return count
}
