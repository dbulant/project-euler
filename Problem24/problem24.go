// problem24
package projecteuler

import (
	. "github.com/dbulant/project-euler/Operations"
)

/*
A permutation is an ordered arrangement of objects.
For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically,
we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

func Problem24() []int {
	permuted := make([]int, 0, 9)
	permuted = append(permuted, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := make([]int, 0, 9)
	index := 0
	permute(&result, permuted, 0, 9, &index)
	return result
}

func permute(result *[]int, permuted []int, i int, n int, index *int) {
	if i == n {
		if (*index) == 1000000 {
			*result = append((*result), permuted...)
		}
		(*index) += 1
		return
	} else {
		for j := i; j <= n; j++ {
			SwapIndexInSlice(permuted, i, j)
			permute(result, permuted, i+1, n, index)
			SwapIndexInSlice(permuted, i, j)
		}
	}
}
