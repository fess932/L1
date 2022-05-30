package l1_test

import (
	"log"
	"testing"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func strSet(input []string) (set map[string]struct{}) {
	set = make(map[string]struct{})
	for _, v := range input {
		set[v] = struct{}{}
	}

	return
}

func Test_strSet(t *testing.T) {
	log.Println(strSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}
