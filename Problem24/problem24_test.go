// problem24_test.go
package projecteuler

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	//a, b, c := 1, 2, 3
	t.Errorf("test permutation called \n")
	permuted := make([]int, 0, 3)
	permuted = append(permuted, 0, 1, 2)
	result := make([][]int, 6)
	index := 0
	permute(&result, permuted, 0, 2, &index)
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
func permute(result *[][]int, permuted []int, i int, n int, index *int) {
	if i == n {
		(*result)[*index] = append((*result)[*index], permuted...)
		(*index) += 1
		return
	} else {
		for j := i; j <= n; j++ {
			swap(permuted, i, j)
			permute(result, permuted, i+1, n, index)
			swap(permuted, i, j)
		}
	}
}

func swap(permuted []int, index1 int, index2 int) {
	tmp := permuted[index1]
	permuted[index1] = permuted[index2]
	permuted[index2] = tmp
}
