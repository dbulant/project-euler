// problem16_test.go
package projecteuler

import (
	"testing"
)

func TestGeneralProblem16(t *testing.T) {
	r := GeneralProblem16(15)
	if r != 26 {
		t.Errorf("expected result is %d, got instead %d \n", 26, r)
	}

	r = GeneralProblem16(1000)
	if r != 89 {
		t.Errorf("expected result is %d, got instead %d \n", 89, r)
	}
}
