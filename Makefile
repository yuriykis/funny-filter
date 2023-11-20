BINARY_NAME=ff
GO=/usr/local/go/bin/go
IP=80.249.99.148
BW_LIMIT=100kbps
P_LIMIT=10
INTERFACE=enp0s5

build:
	@$(GO) build -o ./bin/$(BINARY_NAME) -v

run: build
	@./bin/$(BINARY_NAME)

install: build
	@cp ./bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

test: build
	@$(GO) test -v ./...

set-bw: build
	@./bin/$(BINARY_NAME) bandwidth set --dev $(INTERFACE) --ip $(IP) --limit $(BW_LIMIT)

unset-bw: build
	@./bin/$(BINARY_NAME) bandwidth unset --dev $(INTERFACE) --ip $(IP) --limit $(BW_LIMIT)

set-p: build
	@./bin/$(BINARY_NAME) packets set --ip $(IP) --limit $(P_LIMIT)

unset-p: build
	@./bin/$(BINARY_NAME) packets unset --ip $(IP) --limit $(P_LIMIT)