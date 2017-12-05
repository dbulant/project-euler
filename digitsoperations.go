// digitsoperations
package projecteuler

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
