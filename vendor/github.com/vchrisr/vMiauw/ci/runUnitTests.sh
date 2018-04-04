#!/bin/bash

set -e -u -x

export GOPATH=$(pwd)/gopath
cd gopath/src/github.com/vchrisr/vMiauw/

go get github.com/kr/godep
export PATH=$PATH:$GOPATH/bin

godep go test ./...
