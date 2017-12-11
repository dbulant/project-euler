// pandigital
package projecteuler

var DECIMAL_DIGITS []uint64 = []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}

func isDecimalDigitBaseUniqueInSlice(digs []uint64) bool {
	for _, D := range DECIMAL_DIGITS {
		if !isNumberUniqueInSlice(D, digs) {
			return false
		}
	}
	return true
}

func isNumberUniqueInSlice(number uint64, digs []uint64) bool {
	occurance := 0
	for _, d := range digs {
		if d == number {
			occurance++
		}
	}
	if occurance != 1 {
		return false
	}
	return true
}
