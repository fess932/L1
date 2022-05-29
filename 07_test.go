package l1_test

import (
	"log"
	"sync"
	"testing"
)

type syncmap struct {
	sync.Mutex
	m map[int]int
}

func (s *syncmap) add(key, value int) {
	s.Lock()
	s.m[key] = value
	s.Unlock()
}

func (s *syncmap) get(key int) int {
	s.Lock()
	defer s.Unlock()

	return s.m[key] // zero value if not found
}

func Test_syncmap(t *testing.T) {
	a := syncmap{m: make(map[int]int)}

	a.add(1, 2)
	log.Println(a.get(1))
	log.Println(a.get(2))
}
