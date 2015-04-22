build:
	go build cwrite.go

run:
	GOMAXPROCS=4 go run -race cwrite.go lock_swap_store.go swap_store.go store.go

benchmark:
	GOMAXPROCS=4 go test -bench=.
