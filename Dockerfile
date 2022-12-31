FROM golang:1.19

RUN apt update && \
    apt install -y git protobuf-compiler && \
    apt clean

WORKDIR /workspaces/golang-playground

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/yoheimuta/protolint/cmd/protolint@latest && \
    go install golang.org/x/tools/gopls@latest
    
COPY go.* ./
RUN go mod download