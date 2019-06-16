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

type anagramSolver struct {
	Dict  map[string]bool
	Debug bool
}

func NewAnagramSolver(debug bool) *anagramSolver {
	a := new(anagramSolver)
	a.Debug = debug
	a.Dict = a.loadDictionary()
	return a
}

func (a anagramSolver) loadDictionary() map[string]bool {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "dictionary.txt")
	file, _ := os.Open(filepath)

	defer file.Close()

	words := make(map[string]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}
	return words
}

func (a anagramSolver) GetAnagrams(letters string, minLength int) []Anagrams {
	lettersLength := len(letters)
	var wordLength int

	wordsByLen := make(map[int][]string, 0)

	for word, _ := range a.Dict {
		wordLength = len(word)
		if wordLength < minLength || wordLength > lettersLength {
			continue
		}
		if isWordInLetters(word, letters) {
			wordsByLen[len(word)] = append(wordsByLen[len(word)], word)
		}
	}

	allAnagrams := make([]Anagrams, 0)
	for length, words := range wordsByLen {
		sort.Strings(words)
		allAnagrams = append(allAnagrams, Anagrams{Words: words, Length: length})
	}

	sort.Sort(sort.Reverse(byLetterLength(allAnagrams)))

	return allAnagrams
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
	start := time.Now()

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

	fmt.Println(time.Since(start).String())

	return allAnagrams
}

func GetAnagrams2(letters string, minLength int) []Anagrams {
	printLog("Start")

	var c = make(chan map[string]bool)
	go readDictIntoMap(c)

	dict := <-c

	start := time.Now()

	lettersLength := len(letters)
	var wordLength int

	wordsByLen := make(map[int][]string, 0)

	for word, _ := range dict {
		wordLength = len(word)
		if wordLength < minLength || wordLength > lettersLength {
			continue
		}
		if isWordInLetters(word, letters) {
			wordsByLen[len(word)] = append(wordsByLen[len(word)], word)
		}
	}

	allAnagrams := make([]Anagrams, 0)
	for length, words := range wordsByLen {
		sort.Strings(words)
		allAnagrams = append(allAnagrams, Anagrams{Words: words, Length: length})
	}

	sort.Sort(sort.Reverse(byLetterLength(allAnagrams)))

	//printDebug(wordsByLen)

	//timeTaken := fmt.Sprintf("%.0f", float32(time.Since(start).String()))
	fmt.Println(time.Since(start).String())

	return allAnagrams
}

func isWordInLetters(word string, letters string) bool {
	lettersA := strings.Split(strings.ToLower(letters), "")
	wordA := strings.Split(strings.ToLower(word), "")
	var key int
	for _, letter := range wordA {
		key = find(lettersA, letter)
		if key == len(lettersA) {
			return false
		}
		lettersA = remove(lettersA, key)
	}

	return true
}

func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
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
