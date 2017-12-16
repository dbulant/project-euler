// problem18
package projecteuler

/*
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
*/

type index int
type row []uint64

func Problem18() {
	var pyramid []row = make([]row, 0, 15)
	var i index = 0
	for ; i < 15; i++ {
		switch i {
		case 0:
			createRowAndAppendToPyramid(i, &pyramid, 75)
		case 1:
			createRowAndAppendToPyramid(i, &pyramid, 95, 64)
		case 2:
			createRowAndAppendToPyramid(i, &pyramid, 17, 47, 82)
		case 3:
			createRowAndAppendToPyramid(i, &pyramid, 18, 35, 87, 10)
		case 4:
			createRowAndAppendToPyramid(i, &pyramid, 20, 4, 82, 47, 65)
		case 5:
			createRowAndAppendToPyramid(i, &pyramid, 19, 1, 23, 75, 3, 34)
		case 6:
			createRowAndAppendToPyramid(i, &pyramid, 88, 2, 77, 73, 7, 63, 67)
		case 7:
			createRowAndAppendToPyramid(i, &pyramid, 99, 65, 4, 28, 6, 16, 70, 92)
		case 8:
			createRowAndAppendToPyramid(i, &pyramid, 41, 41, 26, 56, 83, 40, 80, 70, 33)
		case 9:
			createRowAndAppendToPyramid(i, &pyramid, 41, 48, 72, 33, 47, 32, 37, 16, 94, 29)
		case 10:
			createRowAndAppendToPyramid(i, &pyramid, 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14)
		case 11:
			createRowAndAppendToPyramid(i, &pyramid, 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57)
		case 12:
			createRowAndAppendToPyramid(i, &pyramid, 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48)
		case 13:
			createRowAndAppendToPyramid(i, &pyramid, 63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31)
		case 14:
			createRowAndAppendToPyramid(i, &pyramid, 4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23)
		}
	}
}

func createRowAndAppendToPyramid(i index, pyramid *[]row, numbers ...uint64) {
	r := make(row, 0, i)
	r = append(r, numbers...)
	*pyramid = append(*pyramid, r)
}

func traverseAndCompute(p *[]row) uint64 {
	var sum uint64 = 0
	currentMaxRightColumnPosition := 1
	currentMaxLeftColumnPosition := 0
	currentMax := (*p)[0][0]
	sum += currentMax
	for i := 1; i < len(*p); i++ {
		leftValue := (*p)[i][currentMaxLeftColumnPosition]
		rightValue := (*p)[i][currentMaxRightColumnPosition]
		if leftValue > rightValue {
			currentMax = leftValue
			if currentMaxLeftColumnPosition != 0 {
				currentMaxLeftColumnPosition -= 1
			}
		} else {
			currentMax = rightValue
			if currentMaxRightColumnPosition != len((*p)[i]) {
				currentMaxRightColumnPosition += 1
			}
		}
		sum += currentMax
	}
	return sum
}
