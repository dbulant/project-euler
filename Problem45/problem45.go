// Problem45
package projecteuler

/*
It can be verified that T285 = P165 = H143 = 40755.
Find the next triangle number that is also pentagonal and hexagonal.
*/
func FindNextNumber() uint64 {
	var fm map[int]func(uint64) uint64 = make(map[int]func(uint64) uint64)

	fm[0] = computeTriangleNumber
	fm[1] = computePentagonalNumber
	fm[2] = computeHexagonalNumber

	/*
		var startingT uint64 = 285
		var startingP uint64 = 165
		var startingH uint64 = 143
	*/

	var startingT uint64 = 285
	var startingP uint64 = 165
	var startingH uint64 = 143

	var result uint64 = 0
	notFound := true
	for i := startingT + 1; notFound; i++ {
		tn := computeTriangleNumber(i)
		for j := startingP; j < i; j++ {
			pn := computePentagonalNumber(j)
			if tn != pn {
				continue
			}
			for k := startingH; k < j; k++ {
				hn := computeHexagonalNumber(k)
				if tn == pn && pn == hn {
					notFound = false
					result = tn
					break
				}

			}
		}
	}

	return result
}

//Triangle number	 	T(n-position)=n(n+1)/2
func computeTriangleNumber(n uint64) uint64 {
	return (n * (n + 1)) / 2
}

//Pentagonal number	 	Pn=n(3n−1)/2
func computePentagonalNumber(n uint64) uint64 {
	return (n * (3*n - 1)) / 2
}

//Hexagonal	number 	H(n-position)=n(2n−1)
func computeHexagonalNumber(n uint64) uint64 {
	return n * (2*n - 1)
}

func computeNumber(f func(uint64) uint64) {
}
