// problem7_test.go
package projecteuler

import (
	"testing"
)

func TestGeneralProblem7(t *testing.T) {
	pn := GeneralProblem7(10001)
	if pn != 104743 {
		t.Errorf("expected result is %d, got instead %d \n", 104743, pn)
	}
}
