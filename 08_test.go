package l1_test

import (
	"strconv"
	"testing"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(n int, idx uint, val bool) int {
	// прибавляем или вычитаем степень двойки
	if val {
		return n + 1<<idx
	}

	return n - 1<<idx
}

func Test_number(t *testing.T) {
	tcs := []struct {
		idx         uint
		n, expected int
		value       bool
	}{
		{
			n:        8,
			idx:      0,
			value:    true,
			expected: 9,
		},
		{
			n:        8,
			idx:      1,
			value:    true,
			expected: 10,
		},
		{
			n:        9,
			idx:      3,
			value:    false,
			expected: 1,
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := setBit(tc.n, tc.idx, tc.value)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
