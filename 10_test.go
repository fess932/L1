package l1_test

import (
	"fmt"
	"testing"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

// temperatureMerge 2. без сортировки
func temperatureMerge(nums []float64) (result []map[int][]float64) {
	// храним имя температуры и адрес в массиве temps
	temps := map[int]int{}

	for _, v := range nums {
		key := (int(v) / 10) * 10

		if resIdx, ok := temps[key]; ok {
			result[resIdx][key] = append(result[resIdx][key], v)

			continue
		}

		result = append(result, map[int][]float64{key: {v}})
		temps[key] = len(result) - 1
	}

	return result
}

func Test_temperatureMerge(t *testing.T) {
	fmt.Printf("%v\n", temperatureMerge([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}))
}
