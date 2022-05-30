package l1_test

import (
	"log"
	"math/rand"
	"reflect"
	"testing"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// quicksort
func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := rand.Intn(len(arr)) //nolint:gosec
	left := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	for i := 0; i < len(arr); i++ {
		if i == pivot {
			continue
		}

		if arr[i] < arr[pivot] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quicksort(left), arr[pivot]), quicksort(right)...)
}

func Test_quicksort(t *testing.T) {
	s := []int{5, 4, 3, 10, 2, -1, 0, 23, 999}
	s = quicksort(s)
	log.Println(s)

	sorted := []int{-1, 0, 2, 3, 4, 5, 10, 23, 999}
	if reflect.DeepEqual(s, sorted) {
		t.Log("OK")
	} else {
		t.Error("FAIL")
	}
}
