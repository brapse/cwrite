package main

type Store interface {
	Put(input string)
	Get() string
}
