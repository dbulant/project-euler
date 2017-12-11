// problem1_test.go
package projecteuler

import (
	"testing"
)

func TestGeneralProblem1(t *testing.T) {
	s := GeneralProblem1(10, []int{3, 5})
	if s != 23 {
		t.Errorf("Expected result is: %d, got: %d \n", 23, s)
	}
}
