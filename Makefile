.PHONY: all
all: noop classify exec nonser

prepare:
	@go mod tidy
	@mkdir -p bin/

noop: prepare
	go build -o bin/noop examples/noop/main.go

classify: prepare
	go build -o bin/classify examples/classify/main.go

exec: prepare
	go build -o bin/exec examples/exec/main.go

nonser: prepare 
	go build -o bin/nonser examples/nonser/main.go

clean:
	rm -rf bin
