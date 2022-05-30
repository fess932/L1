package l1_test

import (
	"errors"
	"log"
	"strings"
	"testing"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

// TODO: глобальная переменая, паника если строка меньше 100 байт

// как надо

func someFunc() (string, error) {
	strlen := 100

	v := createHugeString(1 << 10)
	if len(v) < strlen {
		return "", errors.New("string is too short")
	}

	return v[:strlen], nil
}

func createHugeString(count int) string {
	var s strings.Builder

	for i := 0; i < count; i++ {
		s.WriteString("a")
	}

	return s.String()
}

func Test_someFunc(t *testing.T) {
	log.Println(someFunc())
}
