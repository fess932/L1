package l1_test

import (
	"testing"
	"time"
)

// Реализовать собственную функцию sleep.
// ???
func sleep(d time.Duration) {
	<-time.After(d)
}

func Test_sleep(t *testing.T) {

	t1 := time.Now()
	sleep(1 * time.Second)
	t.Log("sleep time:", time.Since(t1))
}
