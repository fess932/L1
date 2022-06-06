package l1_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func parallel() {
	randStr := func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, length)
		rand.Read(b)

		return fmt.Sprintf("%x", b)[:length]
	}

	ch := make(chan string)

	go func() { // read
		for {
			fmt.Println(<-ch)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go func() { // write
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done")
			default:
				ch <- randStr(10)
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 10)
	cancel()
}

func Test_paralel(t *testing.T) {
	parallel()
}
