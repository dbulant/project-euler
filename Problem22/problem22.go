// problem22
package projecteuler

import (
	"bufio"
	"os"
	"strings"
)

/*Using names.txt (file is included in this folder), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order.
 Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

const Double_Quote string = "\""
const Apostrophe string = ","

var AlphabetPosition map[rune]uint64 = make(map[rune]uint64)

func init() {
	AlphabetPosition['A'] = 1
	AlphabetPosition['B'] = 2
	AlphabetPosition['C'] = 3
	AlphabetPosition['D'] = 4
	AlphabetPosition['E'] = 5
	AlphabetPosition['F'] = 6
	AlphabetPosition['G'] = 7
	AlphabetPosition['H'] = 8
	AlphabetPosition['I'] = 9
	AlphabetPosition['J'] = 10
	AlphabetPosition['K'] = 11
	AlphabetPosition['L'] = 12
	AlphabetPosition['M'] = 13
	AlphabetPosition['N'] = 14
	AlphabetPosition['O'] = 15
	AlphabetPosition['P'] = 16
	AlphabetPosition['Q'] = 17
	AlphabetPosition['R'] = 18
	AlphabetPosition['S'] = 19
	AlphabetPosition['T'] = 20
	AlphabetPosition['U'] = 21
	AlphabetPosition['V'] = 22
	AlphabetPosition['W'] = 23
	AlphabetPosition['X'] = 24
	AlphabetPosition['Y'] = 25
	AlphabetPosition['Z'] = 26
}

func Problem22() uint64 {
	//entire content of file goes into memory
	file, err := os.Open("p022_names.txt")
	checkError(err)
	defer file.Close()

	buf := bufio.NewReader(file)
	names := make([]string, 0)
	for {
		name, err := buf.ReadString(',')
		name = strings.Replace(name, Double_Quote, "", 2)
		name = strings.Replace(name, Apostrophe, "", 1)
		names = append(names, name)
		if err != nil {
			break
		}
	}
	sortSliceLexiographically(names)
	sum := computeSumOfScores(names)
	return sum
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func computeAlphabeticScore(name string) uint64 {
	var sum uint64 = 0
	for _, n := range name {
		v := AlphabetPosition[n]
		sum += v
	}
	return sum
}

func computeSumOfScores(names []string) uint64 {
	var sum uint64 = 0
	for i, n := range names {
		sc := computeAlphabeticScore(n)
		namescore := (uint64(i + 1)) * sc
		sum += uint64(namescore)
	}
	return sum
}

func sortSliceLexiographically(names []string) {
	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			//if true, names[i] is bigger than names[j], move it to the right
			if c := strings.Compare(names[i], names[j]); c == 1 {
				tmp := names[i]
				names[i] = names[j]
				names[j] = tmp
			}
		}
	}
}
