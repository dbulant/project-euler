// problem22_test
package projecteuler

import (
	"testing"
)

func TestProblem22(t *testing.T) {

	sum := Problem22()
	if sum != 871198282 {
		t.Errorf("sum is: %d, expected %d \n", sum, 871198282)
	}
}
