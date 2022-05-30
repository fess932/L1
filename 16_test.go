package l1_test

import (
	"log"
	"testing"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// quicksort
func quicksort(arr []int) {
	//if len(arr) < 2 {
	//	return
	//}
	//
	//// левая и правая границы
	////left, right := 1, len(arr)-1
	//
	//// точка разделения
	//
	//// [3, 5, 4, 2]
	//// 3 pivot
	//// left [2]
	//// right [5, 4]
	//
	//left, right := []int{}, []int{}
	//pivot := arr[0]
	//
	//for i, v := range arr {
	//	if i == 0 {
	//		continue
	//	}
	//
	//	if v > arr[pivot] {
	//		right--
	//	} else {
	//		arr[left], arr[i] = arr[i], arr[left]
	//		left++
	//	}
	//}
	//
	//quicksort(arr[left:])
	//quicksort(arr[:left])
}

func Test_quicksort(t *testing.T) {
	s := []int{5, 4, 3, 10, 2, -1, 0, 23, 999}
	quicksort(s)
	log.Println(s)
}
