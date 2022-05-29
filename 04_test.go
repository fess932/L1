package l1_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"
)

func runService() {
	ch := make(chan string)
	randStr := func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, length)
		rand.Read(b)
		return fmt.Sprintf("%x", b)[:length]
	}

	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			ch <- randStr(10)
		}
	}()

	worker := func(ch chan string) {
		for {
			fmt.Println(<-ch)
		}
	}

	for i := 0; i < 10; i++ {
		go worker(ch)
	}

	ctx := context.Background()
	ctx, c := signal.NotifyContext(ctx, os.Interrupt)
	defer c()

	select {
	case <-ctx.Done():
		fmt.Println("Done")
	}
}

func Test_worker(t *testing.T) {
	runService()
}
