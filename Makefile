build:
	go build -o bin/covid.out main.go

run: build
	./bin/covid.out
