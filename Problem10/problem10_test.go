// problem10_test.go
package projecteuler

import (
	"testing"
)

func TestGeneralProblem10(t *testing.T) {
	sum := Problem10()
	if sum != 142913828922 {
		t.Errorf("expected result is %d, got instead %d \n", 142913828922, sum)
	}
}
