package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"unsafe"
)

type SwapStore struct {
	sync.RWMutex
	readable unsafe.Pointer
	writable unsafe.Pointer
	rand     *rand.Rand
	pos      int
	size     int
	readyc   chan struct{}
}

func NewSwapStore(size int) *SwapStore {
	a := make([]string, size)
	b := make([]string, size)
	return &SwapStore{
		readable: unsafe.Pointer(&a),
		writable: unsafe.Pointer(&b),
		rand:     rand.New(rand.NewSource(1)),
		pos:      0,
		size:     size,
		readyc:   make(chan struct{}),
	}
}

func (s *SwapStore) Put(input string) {
	wrt := *(*[]string)(s.writable)
	wrt[s.pos] = input

	if s.pos == s.size-1 {
		s.pos = 0

		// swap in writable
		s.writable = atomic.SwapPointer(&s.readable, s.writable)

		// Start
		if s.readyc != nil {
			close(s.readyc)
			s.readyc = nil
		}
	} else {
		s.pos++
	}
}

func (s *SwapStore) Get() string {
	if s.readyc != nil {
		<-s.readyc
	}
	r := *(*[]string)(s.readable)
	return r[s.rand.Intn(len(r))]
}
