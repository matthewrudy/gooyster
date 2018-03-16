run:
	go build -o gooyster .
	./gooyster

test:
	go test -v ./...