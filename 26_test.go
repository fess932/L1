package l1_test

import (
	"testing"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func unique(s string) string {
	return ""
}

func Test_unique(t *testing.T) {
	unique("abc")
}
