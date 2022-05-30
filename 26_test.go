package l1_test

import (
	"strings"
	"testing"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func unique(s string) bool {
	s = strings.ToLower(s)
	entries := map[string]struct{}{}

	for _, v := range s {
		if _, ok := entries[string(v)]; ok {
			return false
		}

		entries[string(v)] = struct{}{}
	}

	return true
}

func Test_unique(t *testing.T) {
	t.Log(unique("abcd"))
	t.Log(unique("aabc"))
	t.Log(unique("Aabc"))
	t.Log(unique("Фвф"))
	t.Log(unique("фыва"))
}
