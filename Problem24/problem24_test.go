// problem24_test.go
package projecteuler

import (
	"testing"

	. "github.com/dbulant/project-euler/Operations"
)

func TestProblem24(t *testing.T) {
	r := Problem24()
	t.Errorf("result is %v \n", r)
}

func TestPermutation(t *testing.T) {
	t.Errorf("test permutation called \n")
	permuted := make([]int, 0, 3)
	permuted = append(permuted, 0, 1, 2)
	result := make([][]int, 6)
	index := 0
	permuteAndSave(&result, permuted, 0, 2, &index)
	t.Errorf("slice of results is %+v \n", result)
}

func bruteForcePermutation(t *testing.T) {
	max := 3
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			for k := 0; k < max; k++ {
				if i != j && j != k && i != k {
					s := make([]int, 0, 3)
					s = append(s, i, j, k)
					t.Errorf("permutation is %v \n", s)
				}
			}
		}
	}
}

//UGLY!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func permuteAndSave(result *[][]int, permuted []int, i int, n int, index *int) {
	if i == n {
		(*result)[*index] = append((*result)[*index], permuted...)
		(*index) += 1
		return
	} else {
		for j := i; j <= n; j++ {
			SwapIndexInSlice(permuted, i, j)
			permuteAndSave(result, permuted, i+1, n, index)
			SwapIndexInSlice(permuted, i, j)
		}
	}
}
