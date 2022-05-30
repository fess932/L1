package l1_test

import (
	"fmt"
	"testing"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func swap(i, j int) (int, int) {
	return j, i
}

func Test_swap(t *testing.T) {
	i, j := 1, 2
	i, j = j, i
	fmt.Println(i, j)
	fmt.Println(swap(1, 2))
}
