package l1_test

import (
	"log"
	"sort"
	"testing"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

func counter(s []int) {
	sort.Ints(s)
}

func Test_counter(t *testing.T) {
	s := []int{5, 4, 3, 10, 2, -1, 0, 999}
	counter(s)
	log.Println(s)
}
