#!/usr/bin/env bash
GREEN='\033[0;32m'
BOLD='\033[1m'
NC='\033[0m' # No Color
# Read variables from key/value pair .env file
export $(grep -v '^#' .env | xargs)

printf "About to sync code\n"
rsync -vzO --exclude '.git' --exclude '.idea' --exclude 'main'  . thefinchmeister@35.185.30.127:/opt/go-anagram-solver/
printf "Code synced! About to build binary and start service\n"
ssh ${USER}@${SERVER_IP} "/usr/local/go/bin/go build /opt/go-anagram-solver/main.go && sudo service goanagram start"
printf "${GREEN}${BOLD}DONE!${NC}\n"
printf "View at http://${SERVER_IP}\n"