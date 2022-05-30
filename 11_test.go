package l1_test

import (
	"log"
	"testing"
)

// Реализовать пересечение двух неупорядоченных множеств.

// intersection of unordered sets
func intersection(a, b []int) (result []int) {
	// O(2n)
	m := make(map[int]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

func Test_wtf(t *testing.T) {
	// 1, 2, 3, 4
	// 3, 2, 5, 6
	// -> 2, 3
	log.Println(intersection([]int{1, 2, 3, 4}, []int{3, 2, 5, 6}))
}
