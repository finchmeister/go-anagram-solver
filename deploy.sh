#!/usr/bin/env bash

git checkout -f origin/master
go get -d ./...
go build main.go
sudo service goanagram restart