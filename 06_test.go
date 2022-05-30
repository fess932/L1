package l1_test

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

func closechan(in <-chan interface{}) {
	for range in {
	}
	fmt.Println("close chan")
}

func cancelCtx(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("cancelCtx")
}

func chanWaitForRead(in <-chan interface{}) {
	<-in
	fmt.Println("chanWaitForRead")
}

func chanWaitForWriteStruct(in chan<- struct{}) {
	in <- struct{}{}
	fmt.Println("chanWaitForWriteStruct")
}

func Test_gorutinesstop(t *testing.T) {
	log.Println("gorutines before: ", runtime.NumGoroutine())

	//1 close chan from writer
	ch := make(chan interface{})
	go closechan(ch)
	sleep(time.Millisecond)
	close(ch)

	//2 cancel context
	ctx, cancel := context.WithCancel(context.Background())
	go cancelCtx(ctx)
	sleep(time.Millisecond)
	cancel()
	sleep(time.Millisecond)

	//3 read from chan
	ch1 := make(chan interface{})
	go chanWaitForRead(ch1)
	sleep(time.Millisecond)
	ch1 <- struct{}{}
	sleep(time.Millisecond)

	//4 write to chan
	ch2 := make(chan struct{})
	go chanWaitForWriteStruct(ch2)
	sleep(time.Millisecond)
	<-ch2
	sleep(time.Millisecond)

	log.Println("gorutines after: ", runtime.NumGoroutine())
}
