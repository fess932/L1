package l1_test

import (
	"log"
	"testing"
)

// Реализовать бинарный поиск встроенными методами языка.

// return index of value or -1 if not found
func bsearch(s []int, value int) int {
	left, right := 0, len(s)-1
	for left <= right {
		mid := (left + right) / 2
		if s[mid] == value {
			return mid
		}

		if s[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func Test_bsearch(t *testing.T) {
	s := []int{-4, -1, 0, 0, 5, 7, 9, 12, 15, 15, 17, 20, 22}
	log.Println(bsearch(s, 12))
	log.Println(bsearch(s, 20))
	log.Println(bsearch(s, 99))
}
