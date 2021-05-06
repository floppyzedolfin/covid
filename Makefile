build:
	go build -o bin/covid.out main.go

test:
	go test ./...

run: build
	./bin/covid.out
