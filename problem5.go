// problem5
package projecteuler

//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func Problem5() uint64 {
	var max uint64 = 20
	n := GeneralProblem5(max)
	return n
}

//Parameter is topDivider, i.e all positive integers from 1 to max
func GeneralProblem5(max uint64) uint64 {
	s := createUintSlice(max)
	maxDivider := productOfSlice(s)
	var i uint64
	for i = 1; i <= maxDivider; i++ {
		if isNumberDivisibleBySlice(i, s) {
			return i
		}
	}
	return maxDivider
}

func productOfSlice(s []uint64) uint64 {
	var product uint64 = 1
	for _, v := range s {
		product *= v
	}
	return product
}

func isNumberDivisibleBySlice(n uint64, dividers []uint64) bool {
	for _, v := range dividers {
		if n%v != 0 {
			return false
		}
	}
	return true
}

func createUintSlice(max uint64) []uint64 {
	slice := make([]uint64, 0, 10)
	for i := 1; uint64(i) <= max; i++ {
		slice = append(slice, uint64(i))
	}
	return slice
}
