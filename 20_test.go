package l1_test

import (
	"fmt"
	"strings"
	"testing"
)

// Разработать программу, которая переворачивает слова в строке.
//	Пример: «snow dog sun — sun dog snow».

func wordReverse(s string) string {
	var result strings.Builder

	split := strings.Split(s, " ")
	for i := len(split) - 1; i >= 0; i-- {
		result.WriteString(split[i])
		result.WriteString(" ")
	}

	return result.String()
}

func Test_wordReverse(t *testing.T) {
	fmt.Println(wordReverse("snow dog sun"))
	fmt.Println(wordReverse("снег собака солнце"))
}
