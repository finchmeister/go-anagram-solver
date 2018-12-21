#!/usr/bin/env bash

echo "About to sync code"
rsync -vzO --exclude '.git' --exclude '.idea' --exclude 'main'  . thefinchmeister@35.185.30.127:/opt/go-anagram-solver/
echo "Code synced! About to build binary and start service"
ssh thefinchmeister@35.185.30.127 "/usr/local/go/bin/go build /opt/go-anagram-solver/main.go && sudo service goanagram start"
echo "DONE!"
echo "View at http://35.185.30.127"