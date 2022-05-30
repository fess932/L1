package l1_test

import (
	"fmt"
	"testing"
)

// Реализовать паттерн «адаптер» на любом примере.

type Namer interface {
	Name() string // returing name
}

func Greater(n Namer) {
	fmt.Println("Hello", n.Name())
}

type Person struct {
	login string
}

func (p Person) Login() string {
	return p.login
}

type Pet struct {
	nickname string
}

func (p Pet) Nick() string {
	return p.nickname
}

type NamePersonAdapter struct {
	Person
}

func (n NamePersonAdapter) Name() string {
	return n.Login()
}

func Test_adapter(t *testing.T) {
	p := Person{login: "Vasya"}

	Greater(NamePersonAdapter{p})
}
