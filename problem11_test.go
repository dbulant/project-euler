// problem11_test
package projecteuler

import (
	"testing"
)

func TestGridToMatrix(t *testing.T) {
	m := gridToMatrix(GRID)
	maxValueFromLeftToRight(m, 4)
	//max := maxDiagonally(m, 4)
	//t.Errorf("max is %d \n", max)

	/*
		m := gridToMatrix(GRID)
		t.Errorf("matrix is: %+v \n", m)
	*/
}
