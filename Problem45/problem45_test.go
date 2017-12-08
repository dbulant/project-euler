// problem45_test
package projecteuler

import (
	"testing"
)

func TestProblem45(t *testing.T) {

	r := FindNextNumber()
	t.Errorf("next number is %d \n", r)

	/*
		tn := computeTriangleNumber(285)
		pn := computePentagonalNumber(165)
		hn := computeHexagonalNumber(143)
		t.Errorf("numbers are %d %d %d \n", tn, pn, hn)
	*/
}
