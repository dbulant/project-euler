// problem18_test.go
package projecteuler

import (
	"testing"
)

/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

	3
   7 4
  2 4 6
 8 5 9 3
*/

func TestCreateRowAndAppendToPyramid(t *testing.T) {
	rowsPyramid := createTestingPyramid()

	if rowsPyramid[0][0] != 3 {
		t.Errorf("pyramid is %+v \n", rowsPyramid)
	}

	if rowsPyramid[1][0] != 7 && rowsPyramid[1][1] != 4 {
		t.Errorf("pyramid is %+v \n", rowsPyramid)
	}

	if rowsPyramid[2][0] != 2 && rowsPyramid[2][1] != 4 && rowsPyramid[2][2] != 6 {
		t.Errorf("pyramid is %+v \n", rowsPyramid)
	}

	if rowsPyramid[3][0] != 8 && rowsPyramid[3][1] != 5 && rowsPyramid[3][2] != 9 && rowsPyramid[3][3] != 3 {
		t.Errorf("pyramid is %+v \n", rowsPyramid)
	}
}

func createTestingPyramid() []row {
	var pyramid []row = make([]row, 0, 4)
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			createRowAndAppendToPyramid(index(i), &pyramid, 3)
		case 1:
			createRowAndAppendToPyramid(index(i), &pyramid, 7, 4)
		case 2:
			createRowAndAppendToPyramid(index(i), &pyramid, 2, 4, 6)
		case 3:
			createRowAndAppendToPyramid(index(i), &pyramid, 8, 5, 9, 3)
		}
	}
	return pyramid
}

func TestTraverseAndCompute(t *testing.T) {
	rowsPyramid := createTestingPyramid()
	p := &rowsPyramid
	sum := traverseAndCompute(p)
	var expected uint64 = 23
	if sum != expected {
		t.Errorf("sum is %d, expected %d \n", sum, expected)
	}
}

func TestTraversePyramid(t *testing.T) {
	rowsPyramid := createTestingPyramid()
	p := &rowsPyramid
	var sum uint64 = 0
	currentMaxRightColumnPosition := 1
	currentMaxLeftColumnPosition := 0
	currentMax := (*p)[0][0]
	sum += currentMax
	for i := 1; i < len(*p); i++ {
		//find elements under current max +1 to right
		leftValue := (*p)[i][currentMaxLeftColumnPosition]
		rightValue := (*p)[i][currentMaxRightColumnPosition]
		if leftValue > rightValue {
			currentMax = leftValue
			//t.Errorf("max value is now is %d ", currentMax)
			if currentMaxLeftColumnPosition != 0 {
				currentMaxLeftColumnPosition -= 1
			}
		} else {
			currentMax = rightValue
			//t.Errorf("max value is %d ", currentMax)
			if currentMaxRightColumnPosition != len((*p)[i]) {
				currentMaxRightColumnPosition += 1
			}
		}
		switch i {
		case 1:
			if currentMax != 7 {
				t.Errorf("expected max %d, got %d instead \n", 7, currentMax)
			}
		case 2:
			if currentMax != 4 {
				t.Errorf("expected max %d, got %d instead \n", 4, currentMax)
			}
		case 3:
			if currentMax != 9 {
				t.Errorf("expected max %d, got %d instead \n", 9, currentMax)
			}
		}
	}
}
