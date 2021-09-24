#!/bin/sh

rm -rf docs
mkdir -p docs/internal
mkdir -p docs/external

GO111MODULE="on" go get -u \
  github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
  github.com/golang/protobuf/protoc-gen-go

go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

echo "==========================================================\nInternal:"
protoc \
  --proto_path=. \
  --proto_path=$(go env GOPATH)/src \
  --proto_path=$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.16.0/third_party/googleapis \
  --doc_out=./docs/internal \
  --doc_opt=./bowstave.template.html,index.html \
  --go_out=plugins=grpc:. \
  *.proto

echo "==========================================================\nExternal:"
protoc \
  --proto_path=. \
  --proto_path=$(go env GOPATH)/src \
  --proto_path=$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.16.0/third_party/googleapis \
  --doc_out=./docs/external \
  --doc_opt=./bowstave.template.html,index.html \
  --go_out=plugins=grpc:. \
  crypto_service.proto keystore.proto meta_writer.proto meta_map.proto key_context.proto void.proto
