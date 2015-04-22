build:
	go build cwrite.go

run:
	GOMAXPROCS=4 go run -race cwrite.go channeled_store.go swap_store.go

benchmark:
	GOMAXPROCS=4 go test -bench=.
