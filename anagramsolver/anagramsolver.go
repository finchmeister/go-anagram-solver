package anagramsolver

import (
	"bufio"
	"fmt"
	"modernc.org/mathutil"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"time"
)

var debug = false

type Anagrams struct {
	Words  []string
	Length int
}

func GetAnagrams(letters string, minLength int) []Anagrams {
	printLog("Start")

	var c = make(chan map[string]bool)
	go readDictIntoMap(c)

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

	var keys []string
	for k := range wordPermsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	printLog("Sorted")

	dict := <-c
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

	sort.Sort(sort.Reverse(byLetterLength(allAnagrams)))

	printDebug(wordsByLen)

	printLog("Searches done")

	return allAnagrams
}

func readDictIntoMap(c chan map[string]bool) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "dictionary.txt")
	file, _ := os.Open(filepath)

	defer file.Close()

	words := make(map[string]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}
	c <- words

	//return words, scanner.Err()
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

type byLetterLength []Anagrams

func (a byLetterLength) Len() int           { return len(a) }
func (a byLetterLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLetterLength) Less(i, j int) bool { return a[i].Length < a[j].Length }
