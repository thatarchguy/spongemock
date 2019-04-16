# Go parameters
BINARY_NAME=spongemock

all: test build
build:
	go build -o $(BINARY_NAME)_ -v
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY_NAME)_*
run:
	go build -o $(BINARY_NAME)_ -v ./...
	./$(BINARY_NAME)_

# Cross compilation
build-all:
	gox -output "bin/{{.Dir}}_{{.OS}}_{{.Arch}}"