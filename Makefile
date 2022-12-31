.PHONY: build genpb clean

build:
	mkdir -p bin
	go build -o bin/main cmd/server/main.go

genpb:
	protoc --go_out=cmd --go-grpc_out=cmd proto/sty.proto

clean:
	rm -rf bin
	rm -rf cmd/pb

runex:
	go run cmd/exercise/*.go