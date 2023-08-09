
all:
	go build -o bin/cgo cgo.go

clean:
	rm bin/*
