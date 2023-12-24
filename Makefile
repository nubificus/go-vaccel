
default: build

build: noop classify exec nonser

prepare:
	@go mod tidy
	@mkdir -p bin


noop: prepare
	go build -o bin/noop noop/main.go

classify: prepare
	go build -o bin/classify classify/main.go

exec: prepare
	go build -o bin/exec exec/main.go

nonser: prepare 
	go build -o bin/nonser nonser/main.go

clean:
	rm bin/*

all: noop classify exec nonser
