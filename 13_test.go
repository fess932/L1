package l1_test

import (
	"fmt"
	"testing"
)

// Поменять местами два числа без создания временной переменной.

func swap(i, j int) (int, int) {
	return j, i
}

func Test_swap(t *testing.T) {
	i, j := 1, 2
	i, j = j, i
	fmt.Println(i, j)
	fmt.Println(swap(1, 2))
}
