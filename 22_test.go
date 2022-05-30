package l1_test

import (
	"fmt"
	"log"
	"math/big"
	"testing"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

func bigCalc(a, b, op string) (string, error) {
	na, ok := big.NewInt(0).SetString(a, 10)
	if !ok {
		return "", fmt.Errorf("ошибка преобразования строки в число: %s", a)
	}

	nb, ok := big.NewInt(0).SetString(b, 10)
	if !ok {
		return "", fmt.Errorf("ошибка преобразования строки в число: %s", b)
	}

	switch op {
	case "+":
		return na.Add(na, nb).String(), nil
	case "-":
		return na.Sub(na, nb).String(), nil
	case "*":
		return na.Mul(na, nb).String(), nil
	case "/":
		return na.Div(na, nb).String(), nil
	default:
		return "", fmt.Errorf("неизвестная операция: %s", op)
	}
}

func Test_bigCalc(t *testing.T) {
	log.Println(bigCalc("12312312312312315453458987908723452345", "12312312311231231235234523452345", "*"))
}
