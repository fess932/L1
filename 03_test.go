package l1_test

import (
	"testing"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func SquaresSum(n []int) int {
	var sum int

	out := make(chan int, 1)
	calc := func(d int, out chan<- int) {
		out <- d * d
	}

	for _, v := range n {
		go calc(v, out)
	}

	for i := 0; i < len(n); i++ {
		sum += <-out
	}

	return sum
}

func Test_SquaresSum(t *testing.T) {
	if s := []int{2, 4, 6, 8, 10}; SquaresSum(s) != 220 {
		t.Error("Expected 220, got ", SquaresSum(s))
	}
}
