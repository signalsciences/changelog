#!/bin/sh
set -ex
go get -u github.com/alecthomas/gometalinter && gometalinter --install
go get -u github.com/golang/dep/... && dep ensure
