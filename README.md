# Go Anagram Solver


Systemd
```
sudo vi /lib/systemd/system/goanagram.service
```

```
[Unit]
Description=goanagram

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/thefinchmeister/go-anagram-solver/main
WorkingDirectory=/home/thefinchmeister/go-anagram-solver

[Install]
WantedBy=multi-user.target
```

```
sudo service goanagram start
```