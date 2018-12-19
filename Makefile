.PHONY: deploy

deploy:
	ssh thefinchmeister@35.185.30.127 "cd ~/go-anagram-solver &&  git fetch -a &&  git checkout -f origin/master && /usr/local/go/bin/go get -d ./... && /usr/local/go/bin/go build main.go && sudo service goanagram restart && echo DONE"