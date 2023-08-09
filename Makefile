
all: noop classify

noop:
	go build -o bin/noop noop.go
classify:
	go build -o bin/classify classify.go

clean:
	rm bin/*
