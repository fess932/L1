package l1_test

import (
	"fmt"
	"testing"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func pipeNumbers(num []int) {
	in := make(chan int)
	out := make(chan int)

	go func(out chan<- int) {
		for _, v := range num {
			in <- v
		}

		// закрываем канал чтобы все следующая горутина в пайплане закончила чтение
		close(out)
	}(in)

	go func(in <-chan int, out chan<- int) {
		for v := range in {
			out <- v * 2
		}

		close(out)
	}(in, out)

	for v := range out {
		fmt.Println(v)
	}
}

func Test_pipeNumbers(t *testing.T) {
	pipeNumbers([]int{1, 2, 3, 4, 5})
}
