package l1_test

import (
	"log"
	"testing"
)

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
