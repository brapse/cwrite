package main

import (
	"math/rand"
	"sync"
	//"fmt"
)

type LockSwapStore struct {
	sync.RWMutex
	readable []string
	writable []string
	rand     *rand.Rand
	pos      int
	size     int
	ready    *Promise
}

func NewLockSwapStore(size int) *LockSwapStore {
	return &LockSwapStore{
		readable: make([]string, size),
		writable: make([]string, size),
		rand:     rand.New(rand.NewSource(1)),
		pos:      0,
		size:     size,
		ready:    NewPromise(),
	}
}

func (s *LockSwapStore) Put(input string) {
	s.writable[s.pos] = input

	if s.pos == s.size-1 {
		s.Lock()
		defer s.Unlock()

		s.pos = 0
		tmp := s.readable
		s.readable = s.writable
		s.writable = tmp

		s.ready.Resolve()
	} else {
		s.pos++
	}
}

func (s *LockSwapStore) Get() string {
	s.ready.Wait()
	s.RLock()
	defer s.RUnlock()

	//fmt.Println("waiting on promise")
	return s.readable[s.rand.Intn(len(s.readable))]
}
