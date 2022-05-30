package l1_test

import (
	"log"
	"testing"
)

// Реализовать бинарный поиск встроенными методами языка.

func bsearch(s []int, x int) int {
	for i := 0; ; i++ {
		if len(s) < 2 {
			return -1
		}

		a := len(s) / 2
		if s[a] == x {
			return i
		}

		if s[a] > x {
			s = s[:a]
		} else {
			s = s[a+1:]
		}
	}
}

func Test_bsearch(t *testing.T) {
	s := []int{-4, -1, 0, 0, 5, 7, 9, 12, 15, 15, 17, 20, 22}
	log.Println(bsearch(s, 12))
}
