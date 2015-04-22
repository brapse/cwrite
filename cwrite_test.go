package main

import (
  "testing"
)

func BenchmarkUnsafe(b *testing.B) {
  for i := 0; i < b.N; i++ {
    unsafeRun()
  }
}

func BenchmarkLocked(b *testing.B) {
  for i := 0; i < b.N; i++ {
    lockedRun()
  }
}

func BenchmarkChanneled(b *testing.B) {
  for i := 0; i < b.N; i++ {
    chanRun()
  }
}

func BenchmarkSwap(b *testing.B) {
  for i := 0; i < b.N; i++ {
    swapRun()
  }
}
