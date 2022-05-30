package l1_test

import (
	"testing"
)

// Реализовать паттерн «адаптер» на любом примере.

func adapter(s string) string {
	return ""
}

func Test_adapter(t *testing.T) {
	adapter("abc")
}
