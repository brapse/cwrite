package main

import (
	"math/rand"
	"sync"
)

type LockedStore struct {
	store []string
	lock  *sync.Mutex
	rand  *rand.Rand
	pos   int
	ready *Promise
}

func NewLockedStore(size int) *LockedStore {
	return &LockedStore{
		store: make([]string, size),
		lock:  &sync.Mutex{},
		rand:  rand.New(rand.NewSource(1)),
		pos:   0,
		ready: NewPromise(),
	}
}

func (s *LockedStore) Put(input string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.store[s.pos%cap(s.store)] = input
	s.pos++
	s.ready.Resolve()
}

func (s *LockedStore) Get() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.ready.Wait()

	return s.store[s.rand.Intn(len(s.store))]
}
