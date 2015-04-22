package main

import (
  "fmt"
  "sync"
)

var testString = "likewtfomgbbq"
var steps = 100000
var sampleSize = 10

func unsafeRun() {
  ed := make([]string, 1)
  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    for i := 0; i < steps; i++ {
      ed[0] = testString
    }
    wg.Done()
  }()

  go func() {
    for i := 0; i < steps; i++ {
      _ = ed[0]
      //if current != testString {
      //  fmt.Printf("misread!: %s\n", current)
      //}
    }
    wg.Done()
  }()

  wg.Wait()
}

func lockedRun() {
  ed := make([]string, 1)
  var wg sync.WaitGroup
  var mutex = &sync.Mutex{}
  wg.Add(2)

  go func() {
    for i := 0; i < steps; i++ {
      mutex.Lock()
      ed[0] = testString
      mutex.Unlock()
    }
    wg.Done()
  }()

  go func() {
    for i := 0; i < steps; i++ {
      mutex.Lock()
      current := ed[0]
      mutex.Unlock()
      if current != testString {
        //fmt.Printf("misread!: %s\n", current)
      }
    }
    wg.Done()
  }()

  wg.Wait()
}

func chanRun() {
  store := NewChanneledStore(sampleSize)
  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    for i := 0; i < steps; i++ {
      //fmt.Println("write")
      store.Put(testString)
    }
    wg.Done()
  }()

  go func() {
    for i := 0; i < steps; i++ {
      //fmt.Println("read")
      current := store.Get()
      if current != testString {
        fmt.Printf("misread!: %s\n", current)
      }
    }
    wg.Done()
  }()

  wg.Wait()
}

func swapRun() {
  store := NewSwapStore(sampleSize)
  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    for i := 0; i < steps; i++ {
      store.Put(testString)
    }
    wg.Done()
  }()

  go func() {
    for i := 0; i < steps; i++ {
      current := store.Get()
      if current != testString {
        fmt.Printf("misread!: %s\n", current)
      }
    }
    wg.Done()
  }()

  wg.Wait()
}



func main() {
  fmt.Printf("Going for it\n")
  chanRun()
}
