#!/bin/bash

set -e -u -x

DIR="${GOPATH}/src/github.com/patrickrand/concourse-maven-resource"
mkdir -p $DIR/assets

go build -o ${DIR}/assets/in ${DIR}/in/main.go
go build -o ${DIR}/assets/check ${DIR}/check/main.go
go build -o ${DIR}/assets/out ${DIR}/out/main.go