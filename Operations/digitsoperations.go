// digitsoperations
package projecteuler

func IsNumberPalindrome(n uint64) bool {
	d := GetDigitsFromNumber(n)
	for i, j := 0, len(d); i < len(d)/2; i++ {
		if d[i] != d[j-i-1] {
			return false
		}
	}
	return true
}

func IsBinaryNumberPalindrome(n uint64) bool {
	var digits []uint64
	if n%2 == 0 {
		digits = evenUint64ToBinary(n)
	} else {
		digits = oddUint64ToBinary(n)
	}
	for i, j := 0, len(digits); i < len(digits)/2; i++ {
		if digits[i] != digits[j-i-1] {
			return false
		}
	}
	return true
}

func GetDigitsFromNumber(n uint64) []uint64 {
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
