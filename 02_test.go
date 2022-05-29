package l1_test

import (
	"os"
	"strconv"
	"sync"
	"testing"
)

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
