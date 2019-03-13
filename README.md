# Go Anagram Solver

Find anagrams of words based upon the unix dictionary found at `anagramsolver/dictionary.txt`. 

## Lambda Function Implementation

See `/cloudfunction`, with the static page is hosted
on GitHub pages in `/docs`.

To deploy on GCP:
```
cd cloudfunction
gcloud functions deploy HelloYou --runtime go111 --trigger-http --memory=2048 --region=europe-west1
```

Compute anagrams: 
<https://europe-west1-anagram-solver-you.cloudfunctions.net/HelloYou?q=aardvark>

## Go Web App Implementation

### Dev

```
go run main.go
```

View at http://localhost:80

### Deploy

Copy `.env.dist` to `.env` with working vars then run:
 
 ```
make deploy
 ```

### Production

Running as a systemd service.

```
sudo service goanagram start
```

To view logs:
```
journalctl -u goanagram.service
```

### Tests

```
cd anagramsolver
go test -v
```