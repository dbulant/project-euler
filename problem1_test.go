// problem1_test
package main

import "testing"

func TestGeneralProblem1(t *testing.T) {
	s := GeneralProblem1(10, []int{3, 5})
	if s != 23 {
		t.Error("Expected result is: %d, got: %d \n", 23, s)
	}
}
