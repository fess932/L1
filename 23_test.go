package l1_test

import (
	"fmt"
	"log"
	"testing"
)

// Удалить i-ый элемент из слайса.

func sliceDelete(slice []int, i int) ([]int, error) {
	if i < 0 || i >= len(slice) {
		return nil, fmt.Errorf("неверный индекс")
	}

	slice = append(slice[:i], slice[i+1:]...)

	return slice, nil
}

func Test_sliceDelete(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice, err := sliceDelete(s, 2)
	if err != nil {
		t.Errorf("ошибка: %s", err)

		return
	}

	log.Println(slice)

	slice, err = sliceDelete(s, 2)
	if err != nil {
		t.Errorf("ошибка: %s", err)

		return
	}

	log.Println(slice)
}
