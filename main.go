package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"modernc.org/mathutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type Anagrams struct {
	Words  []string
	Length int
}

var (
	indexTemplate = template.Must(template.ParseFiles("template/index.html"))
	debug         = false
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	q := r.FormValue("q")
	printDebug(q)

	if len(q) > 9 {
		q = q[:9]
	}

	anagrams := getAnagrams(q, 3)

	timeTaken := int(time.Since(start).Nanoseconds() / 1000000)
	params := TemplateFields{Query: q, Anagrams: anagrams, TimeTaken: timeTaken}

	indexTemplate.Execute(w, params)
}

type TemplateFields struct {
	Query     string
	Anagrams  []Anagrams
	TimeTaken int
}

func getAnagrams(letters string, minLength int) []Anagrams {
	printLog("Start")
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

	printLog("Perms done")

	dict, _ := readDictIntoMap()
	printLog("File converted to map")

	var keys []string
	for k := range wordPermsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	printLog("Sorted")

	wordsByLen := make(map[int][]string, 0)
	for _, key := range keys {
		if dict[key] == true {
			wordsByLen[len(key)] = append(wordsByLen[len(key)], key)
		}
	}

	allAnagrams := make([]Anagrams, 0)
	for length, words := range wordsByLen {
		allAnagrams = append(allAnagrams, Anagrams{Words: words, Length: length})
	}

	sort.Sort(sort.Reverse(ByLetterLength(allAnagrams)))

	printDebug(wordsByLen)

	printLog("Searches done")

	return allAnagrams
}

func readDictIntoMap() (map[string]bool, error) {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make(map[string]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}
	return words, scanner.Err()
}

func printDebug(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}

func printLog(title string) {
	if debug {
		fmt.Println(time.Now().Format(time.RFC3339Nano) + " " + title)
	}
}

type ByLetterLength []Anagrams

func (a ByLetterLength) Len() int           { return len(a) }
func (a ByLetterLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLetterLength) Less(i, j int) bool { return a[i].Length < a[j].Length }
