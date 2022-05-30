package l1_test

import (
	"log"
	"sort"
	"testing"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(s []int) {
	sort.Ints(s)
}

func Test_quicksort(t *testing.T) {
	s := []int{5, 4, 3, 10, 2, -1, 0, 999}
	quicksort(s)
	log.Println(s)
}
