BINARY_NAME=ff

build:
	go build -o ./bin/$(BINARY_NAME) -v

run: build
	./bin/$(BINARY_NAME)