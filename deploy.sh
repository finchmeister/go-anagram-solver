#!/usr/bin/env bash
GREEN='\033[0;32m'
RED='\033[0;31m'
BOLD='\033[1m'
NC='\033[0m' # No Color
# Read variables from key/value pair .env file
export $(grep -v '^#' .env | xargs)
SERVER=${USER}@${SERVER_IP}

printf "About to sync code\n"

rsync -vzO --exclude '.git' --exclude '.idea' --exclude 'main'  . ${SERVER}:/opt/go-anagram-solver/
if [[ $? -gt 0 ]]
then
    printf "${RED}${BOLD}FAIL! Unable to rsync code${NC}\n"
    exit 1
fi
printf "Code synced! About to build binary and start service\n"

ssh ${SERVER} "/usr/local/go/bin/go build /opt/go-anagram-solver/main.go && sudo service goanagram start"
if [[ $? -gt 0 ]]
then
    printf "${RED}${BOLD}FAIL! Unable to build and start service${NC}\n"
    exit 1
fi

printf "${GREEN}${BOLD}DONE!${NC}\n"
printf "View at http://${SERVER_IP}\n"
exit 0