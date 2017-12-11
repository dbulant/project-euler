// problem4_test.go
package projecteuler

import (
	"testing"
)

func TestGeneralProblem4(t *testing.T) {
	n := GeneralProblem4(2)
	var expected uint64 = 9009
	if n != expected {
		t.Errorf("expected result is %d, got instead %d \n", expected, n)
	}
}
