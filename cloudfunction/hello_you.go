package cloudfunction

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/finchmeister/go-anagram-solver/anagramsolver"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	stdLogger = log.New(os.Stdout, "", 0)
	logger    = log.New(os.Stderr, "", 0)
	a         = anagramsolver.NewAnagramSolver(false, false)
)

func init() {
	stdLogger.Println("Init - reading dictionary")
	a.Dict = readDictIntoMap()
}

func HelloYou(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("cache-control", "max-age=86400")
	start := time.Now()
	q := r.URL.Query().Get("q")

	var anagrams []anagramsolver.Anagrams
	stdLogger.Println("Start: " + q)
	if len(q) < 9 {
		// Algo 1
		anagrams = a.GetAnagramsAlgo1(q, 3)
	} else {
		// Algo 3
		anagrams = a.GetAnagramsAlgo3(q, 3)
	}

	timeTaken := fmt.Sprintf("%.0f", float32(time.Since(start).Nanoseconds()/1000000))
	stdLogger.Println("Time: " + timeTaken + "ms")

	if q == "" {
		fmt.Fprint(w, "Hello, you!")
		return
	}

	b, err := json.Marshal(anagrams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}

func readDictIntoMap() map[string]bool {

	fileUrl := "https://raw.githubusercontent.com/finchmeister/go-anagram-solver/master/anagramsolver/dictionary.txt"
	stdLogger.Println("Downloading dictionary")

	if err := DownloadFile("/tmp/dictionary.txt", fileUrl); err != nil {
		panic(err)
	}

	stdLogger.Println("Reading dictionary")
	file, err := os.Open("/tmp/dictionary.txt")

	if err != nil {
		logger.Println(err)
	}

	defer file.Close()

	words := make(map[string]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}
	stdLogger.Println("Dictionary read")

	return words
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
