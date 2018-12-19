#!/usr/bin/env bash

git reset --hard
git pull
go get -d ./...
go build main.go
sudo service goanagram restart