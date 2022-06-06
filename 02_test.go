package l1_test

import (
	"os"
	"strconv"
	"sync"
	"testing"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func PrintSquares(n []int) {
	var wg sync.WaitGroup

	wg.Add(len(n))

	for _, v := range n {
		go func(d int) {
			os.Stdout.WriteString(strconv.Itoa(d*d) + "\n")
			wg.Done()
		}(v)
	}

	wg.Wait()
}

func Test_PrintSquares(t *testing.T) {
	PrintSquares([]int{1, 2, 3, 4, 5})
}
