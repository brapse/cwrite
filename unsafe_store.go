package main

import (
	"math/rand"
)

type UnsafeStore struct {
	store []string
	rand  *rand.Rand
	pos   int
}

func NewUnsafeStore(size int) *UnsafeStore {
	return &UnsafeStore{
		store: make([]string, size),
		rand:  rand.New(rand.NewSource(1)),
		pos:   0,
	}
}

func (s *UnsafeStore) Put(input string) {
	s.store[s.pos%cap(s.store)] = input
	s.pos++
}

func (s *UnsafeStore) Get() string {
	return s.store[s.rand.Intn(len(s.store))]
}
