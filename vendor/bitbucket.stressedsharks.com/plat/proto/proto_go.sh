#!/bin/bash

# Script to generate the proto files for go
# Best way to do it is to run through a bind mount docker container
#    docker container run --rm --name proto_gen --mount type=bind,source=/your/local/path/to/platformservices/proto,target=/development/ps/proto -it jwahab/go-protoc:latest /bin/bash -c "cd /development/ps/proto; ./proto_go.sh"
#
# Note : Change appropriately [/path/to/local/proto] to what you have on your local machine
#

cd $(dirname "$0")

# clean all existing generated protobuf go files
find . -type f -name "*pb*.go" | xargs rm ;

GO111MODULE="on" go get -u \
  github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2

GOFLAGS="-mod=mod" go install \
  github.com/golang/protobuf/proto \
  github.com/golang/protobuf/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc

function pc() {
  GOFLAGS="-mod=mod" protoc \
  --proto_path=. \
  --proto_path=$(go env GOPATH)/src \
  --proto_path=$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.16.0/third_party/googleapis \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  $1
}

find . -type f -name "*.proto" -print | while read F; do
  echo Generate grpc files for [$F]
  pc $F
done


# Git to check if generated file has change
# Ignore this message if the change is your intention
modified="$(git status --untracked-files=no --porcelain)"
if [[ -z ${modified} ]]
then
    echo -e "\n\nproto_go[success] : generated .go files has no changes"
else
    echo -e "\n\nproto_go[error] :"
    echo "  Newly generated .go files has changes. To see changes detail : [git diff]"
    echo "  Ignore this error message if these changes are intended"
    exit 1
fi
