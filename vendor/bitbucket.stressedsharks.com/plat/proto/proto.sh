#!/bin/bash

cd $(dirname "$0")
rm -f ./*.pb.go
rm -rf javascript

mkdir javascript

# npm install -g grpc-tools

GO111MODULE="on" go get -u \
  github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2

function pc() {
  protoc \
  --proto_path=$(go env GOPATH)/src \
  --proto_path=. \
  --proto_path=$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.16.0/third_party/googleapis \
  --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:. --js_out=import_style=commonjs,binary:./javascript \
  --grpc_out=./javascript \
  --plugin=protoc-gen-grpc=$(which grpc_tools_node_protoc_plugin) $1
}

find . -maxdepth 1 -not -path "./examples/*" -not -path "./.history/*" -type f -name "*.proto" -print | while read F; do
  echo $F
  pc $F
done

# echo "GRPC gateway: notary"
#  protoc --proto_path=./ notary.proto \
#   -I$(go env GOPATH)/src \
#   -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#   --grpc-gateway_out=logtostderr=true:. \

# echo "GRPC gateway: funding service"
#  protoc --proto_path=./ funding_service.proto \
#   -I$(go env GOPATH)/src \
#   -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#   --grpc-gateway_out=logtostderr=true:. \
