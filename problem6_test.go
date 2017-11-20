// problem6_test
package projecteuler

import "testing"

func TestSumOfSquares(t *testing.T) {
	var n uint64 = 10
	r := sumOfSquares(n)
	if r != 385 {
		t.Errorf("expected result is %d, got instead %d \n", 385, r)
	}
}

func TestSquareOfSum(t *testing.T) {
	var n uint64 = 10
	r := squareOfSum(n)
	if r != 3025 {
		t.Errorf("expected result is %d, got instead %d \n", 3025, r)
	}
}

func TestGeneralProblem6(t *testing.T) {
	diff := GeneralProblem6(10)
	if diff != 2640 {
		t.Errorf("expected result is %d, got instead %d \n", 2640, diff)
	}
}
