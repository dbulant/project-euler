// project14_test
package projecteuler

import "testing"

func TestGeneralProblem14(t *testing.T) {
	r, chain := GeneralProblem14(1000000)
	if len(chain) == 0 {
		t.Errorf("number with longest chain is %d \n", r)
		t.Errorf("chain is %+v \n", chain)
	}
}
