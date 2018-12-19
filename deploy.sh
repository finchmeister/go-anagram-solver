#!/usr/bin/env bash

git fetch -a
git checkout -f origin/master
go get -d ./...
go build main.go
sudo service goanagram restart