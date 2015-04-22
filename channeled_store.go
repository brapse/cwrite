package main

import (
  "math/rand"
  "sync"
  "fmt"
)

type ChanneledStore struct {
  sync.RWMutex
  readable []string
  writable []string
  rand *rand.Rand
  pos int
  size int
  ready bool
  readyc chan bool
}

func NewChanneledStore(size int) *ChanneledStore {
  return &ChanneledStore{
    readable: make([]string, size),
    writable: make([]string, size),
    rand: rand.New(rand.NewSource(1)),
    pos: 0,
    size: size,
    ready: false,
    readyc: make(chan bool, 1),
  }
}

func (s *ChanneledStore) Put (input string) {
  s.writable[s.pos] = input

  if s.pos == s.size-1 {
    s.Lock()
    defer s.Unlock()

    s.pos = 0
    tmp := s.readable
    s.readable = s.writable
    s.writable = tmp

    if !s.ready {
      fmt.Println("Unlocking!")
      s.ready = true
      s.readyc <- true
    }
  } else {
    s.pos++
  }
}

func (s *ChanneledStore) Get() string {
  s.RLock()
  defer s.RUnlock()

  if !s.ready {
    fmt.Println("WAiting on chan")
    <-s.readyc
  }

  return s.readable[s.rand.Intn(len(s.readable))]
}
