package main

import (
	"testing"
)

func BenchmarkUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestRun(NewUnsafeStore(sampleSize))
	}
}

func BenchmarkLocked(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestRun(NewLockedStore(sampleSize))
	}
}

func BenchmarkLockSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestRun(NewLockSwapStore(sampleSize))
	}
}

func BenchmarkSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestRun(NewSwapStore(sampleSize))
	}
}
