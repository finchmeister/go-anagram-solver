#!/usr/bin/env bash
GREEN='\033[0;32m'
RED='\033[0;31m'
BOLD='\033[1m'
NC='\033[0m' # No Color

function writeLn {
    printf "==> $1\n"
}

writeLn "Running tests"

go test
if [[ $? -gt 0 ]]
then
    writeLn "${RED}${BOLD}FAIL! Tests failed${NC}\n"
    exit 1
fi

# Read variables from key/value pair .env file
export $(grep -v '^#' .env | xargs)
SERVER=${USER}@${SERVER_IP}

writeLn "About to sync code"

rsync -rvzO --exclude '.git' --exclude '.idea' --exclude 'main'  . ${SERVER}:/opt/go-anagram-solver/
if [[ $? -gt 0 ]]
then
    writeLn "${RED}${BOLD}FAIL! Unable to rsync code${NC}"
    exit 1
fi
writeLn "Code synced! About to build binary and start service"

ssh ${SERVER} "cd /opt/go-anagram-solver && /usr/local/go/bin/go build main.go && sudo service goanagram start"
if [[ $? -gt 0 ]]
then
    writeLn "${RED}${BOLD}FAIL! Unable to build and start service${NC}"
    exit 1
fi

writeLn "${GREEN}${BOLD}SUCCESS!${NC}"
writeLn "View at http://${SERVER_IP}"
exit 0