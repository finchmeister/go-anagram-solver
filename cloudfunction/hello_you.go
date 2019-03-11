package cloudfunction

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"modernc.org/mathutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	stdLogger  = log.New(os.Stdout, "", 0)
	logger     = log.New(os.Stderr, "", 0)
	dictionary = make(map[string]bool, 0)
)

func init() {
	stdLogger.Println("Init - reading dictionary")
	//go readDictIntoMap(c)
	dictionary = readDictIntoMap()
}

func HelloYou(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	q := r.URL.Query().Get("q")

	if len(q) > 9 {
		q = q[:9]
	}

	stdLogger.Println("Start: " + q)

	anagrams := GetAnagrams(q, 3)
	timeTaken := int(time.Since(start).Nanoseconds() / 1000000)
	stdLogger.Println("Time: " + string(timeTaken) + "ms")

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
	//fmt.Fprintf(w, "Hello, %s!", html.EscapeString(q))
}

var debug = false

type Anagrams struct {
	Words  []string
	Length int
}

func GetAnagrams(letters string, minLength int) []Anagrams {
	//var c = make(chan map[string]bool)
	//go readDictIntoMap(c)

	wordPermsMap := make(map[string]string, 0)
	st := strings.Split(strings.ToLower(letters), "")
	stLen := len(st)
	sort.Strings(st)

	for {
		wordPermsMap[strings.Join(st, "")] = ""
		if !mathutil.PermutationNext(sort.StringSlice(st)) {
			break
		}
	}

	for wordLen := stLen - 1; wordLen >= minLength; wordLen-- {
		for perm := range wordPermsMap {
			perm = perm[0:wordLen]
			wordPermsMap[perm] = ""
		}
	}

	stdLogger.Println("Perms done")

	var keys []string
	for k := range wordPermsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	stdLogger.Println("Sorted")

	wordsByLen := make(map[int][]string, 0)
	for _, key := range keys {
		if dictionary[key] == true {
			wordsByLen[len(key)] = append(wordsByLen[len(key)], key)
		}
	}

	allAnagrams := make([]Anagrams, 0)
	for length, words := range wordsByLen {
		allAnagrams = append(allAnagrams, Anagrams{Words: words, Length: length})
	}

	sort.Sort(sort.Reverse(byLetterLength(allAnagrams)))

	printDebug(wordsByLen)

	stdLogger.Println("Searches done")

	return allAnagrams
}

func readDictIntoMap() map[string]bool {

	fileUrl := "https://raw.githubusercontent.com/finchmeister/go-anagram-solver/master/anagramsolver/dictionary.txt"
	stdLogger.Println("Downloading dictionary")

	if err := DownloadFile("/tmp/dictionary.txt", fileUrl); err != nil {
		panic(err)
	}

	stdLogger.Println("Reading dictionary")
	//_, filename, _, _ := runtime.Caller(0)
	//filepath := path.Join(path.Dir(filename), "dictionary.txt")
	//file, err := os.Open(filepath)
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

func printDebug(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}

//func printLog(title string) {
//	if debug {
//		fmt.Println(time.Now().Format(time.RFC3339Nano) + " " + title)
//	}
//}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

type byLetterLength []Anagrams

func (a byLetterLength) Len() int           { return len(a) }
func (a byLetterLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLetterLength) Less(i, j int) bool { return a[i].Length < a[j].Length }
