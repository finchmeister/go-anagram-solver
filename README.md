# Go Anagram Solver

Find anagrams of words based upon the unix dictionary found at `/usr/share/dict/words`. 

## Dev

```
go run main.go
```

View at http://localhost:80

## Deploy

Copy `.env.dist` to `.env` with working vars then run `make deploy`.

## Production

Running as a systemd service.

```
sudo service goanagram start
```

To view logs:
```
journalctl -u goanagram.service
```

## Tests

```
go test -v
```