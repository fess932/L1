package l1_test

import (
	"log"
	"testing"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание(композиция) методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
}

func (h *Human) Say() {
	log.Println("Hello, my name is", h.Name)
}

type Action struct {
	Human
}

func Test_test(t *testing.T) {
	h := Human{Name: "Tom"}
	a := Action{h}
	a.Say()
}
