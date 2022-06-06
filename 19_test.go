package l1_test

import (
	"fmt"
	"strings"
	"testing"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

// более универсальный вариант
func strReverse2(s string) string {
	var b strings.Builder

	strs := strings.Split(s, "")

	for i := len(strs) - 1; i > -1; i-- {
		b.WriteString(strs[i])
	}

	return b.String()
}

func Test_strReverse(t *testing.T) {
	fmt.Println(strReverse2("abc"))
	fmt.Println(strReverse2("фыва"))
}
