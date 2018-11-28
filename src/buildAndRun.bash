#!/bin/bash

rm multisite-search;

export GOPATH=$GOPATH:$(pwd)/../
go build -o multisite-search && ./multisite-search
