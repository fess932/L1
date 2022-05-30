package l1_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type counter int64

func (c *counter) add() {
	atomic.AddInt64((*int64)(c), 1)
}

func Test_counter(t *testing.T) {
	c := counter(0)

	for i := 0; i < 1000; i++ {
		go c.add()
	}

	time.Sleep(time.Second)
	fmt.Println(c)
}
