#!/usr/bin/env bash

sudo apt-get install golang-go

sudo mkdir -p /opt/go-anagram-solver/
sudo chmod -R 775 /opt/
sudo chgrp -R google-sudoers /opt/

cp config/goanagram.service /lib/systemd/system/goanagram.service