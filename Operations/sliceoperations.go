// sliceoperations
package projecteuler

func SumOfSlice(s []uint64) uint64 {
	var sum uint64 = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func ProductOfSlice(s []uint64) uint64 {
	var product uint64 = 1
	for _, v := range s {
		product *= v
	}
	return product
}

func productOfFloatSlice(s []float64) float64 {
	var product float64 = 1
	for _, v := range s {
		product *= v
	}
	return product
}

func CreateUintSlice(max uint64) []uint64 {
	slice := make([]uint64, 0, 10)
	for i := 1; uint64(i) <= max; i++ {
		slice = append(slice, uint64(i))
	}
	return slice
}

func areSlicesEqual(first []uint64, second []uint64) bool {
	if len(first) != len(second) {
		return false
	}
	for i, _ := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func smallerOrEqual(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//Swap two elements identified by their index in slice
func SwapIndexInSlice(s []int, index1 int, index2 int) {
	tmp := s[index1]
	s[index1] = s[index2]
	s[index2] = tmp
}
