#!/usr/bin/env bash

git fetch -a
git checkout -f origin/master
go get -d ./...
go build main.go
sudo service goanagram restart
echo "DONE"

#git fetch -a
#git checkout -f origin/master
#/usr/local/go/bin/go get -d ./...
#/usr/local/go/bin/go build main.go
#sudo service goanagram restart
#echo "DONE"

#cd ~/go-anagram-solver &&  git fetch -a &&  git checkout -f origin/master && /usr/local/go/bin/go get -d ./... && /usr/local/go/bin/go build main.go && sudo service goanagram restart && echo "DONE"