#!/usr/bin/env bash

sudo apt-get install golang-go
sudo apt-get install wbritish

chmod +x update.sh
chmod +x install.sh

echo "[Unit]
Description=goanagram

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/thefinchmeister/go-anagram-solver/main
WorkingDirectory=/home/thefinchmeister/go-anagram-solver

[Install]
WantedBy=multi-user.target" | sudo tee --append /lib/systemd/system/goanagram.service > /dev/null

sudo service goanagram start
