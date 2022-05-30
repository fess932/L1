package l1_test

import (
	"log"
	"sort"
	"testing"
)

// Реализовать бинарный поиск встроенными методами языка.

func bsearch(s []int) {
	sort.Ints(s)
}

func Test_bsearch(t *testing.T) {
	s := []int{5, 4, 3, 10, 2, -1, 0, 999}
	bsearch(s)
	log.Println(s)
}
