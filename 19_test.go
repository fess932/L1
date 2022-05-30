package l1_test

import (
	"fmt"
	"strings"
	"testing"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func strReverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteString(string(s[i]))
	}

	return result.String()
}

func Test_strReverse(t *testing.T) {
	fmt.Println(strReverse("abc"))
}
