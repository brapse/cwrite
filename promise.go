package main

import (
	"sync/atomic"
)

type Promise struct {
	ready int32
}

func NewPromise() *Promise {
	return &Promise{
		ready: 0,
	}
}

func (p *Promise) Resolve() {
	if !p.Get() {
		atomic.StoreInt32(&p.ready, 1)
	}
}

func (p *Promise) Get() bool {
	return atomic.LoadInt32(&p.ready) == 1
}

func (p *Promise) Wait() bool {
	for !p.Get() {
	}

	return true
}
