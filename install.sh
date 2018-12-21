#!/usr/bin/env bash

sudo apt-get install golang-go
sudo apt-get install wbritish

sudo mkdir -p /opt/go-anagram-solver/
sudo chmod -R 775 /opt/
sudo chgrp -R google-sudoers /opt/

echo "[Unit]
Description=goanagram

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/opt/go-anagram-solver/go-anagram-solver/main
WorkingDirectory=/opt/go-anagram-solver/go-anagram-solver

[Install]
WantedBy=multi-user.target" | sudo tee --append /lib/systemd/system/goanagram.service > /dev/null