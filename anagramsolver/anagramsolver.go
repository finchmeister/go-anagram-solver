package anagramsolver

import (
	"bufio"
	"modernc.org/mathutil"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
)

type Anagrams struct {
	Words  []string
	Length int
}

type anagramSolver struct {
	Dict  map[string]bool
	Debug bool
}

func NewAnagramSolver(useLocalDict bool, debug bool) *anagramSolver {
	a := new(anagramSolver)
	a.Debug = debug
	if useLocalDict {
		a.Dict = a.loadDictionary()
	}
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

func (a anagramSolver) GetAnagramsAlgo3(letters string, minLength int) []Anagrams {
	lettersLength := len(letters)
	var wordLength int

	wordsByLen := make(map[int][]string, 0)

	for word := range a.Dict {
		wordLength = len(word)
		if wordLength < minLength || wordLength > lettersLength {
			continue
		}
		if a.isWordInLetters(word, letters) {
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

func (a anagramSolver) GetAnagramsAlgo1(letters string, minLength int) []Anagrams {
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

	var keys []string
	for k := range wordPermsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	wordsByLen := make(map[int][]string, 0)
	for _, key := range keys {
		if a.Dict[key] == true {
			wordsByLen[len(key)] = append(wordsByLen[len(key)], key)
		}
	}

	allAnagrams := make([]Anagrams, 0)
	for length, words := range wordsByLen {
		allAnagrams = append(allAnagrams, Anagrams{Words: words, Length: length})
	}

	sort.Sort(sort.Reverse(byLetterLength(allAnagrams)))

	return allAnagrams
}

func (a anagramSolver) isWordInLetters(word string, letters string) bool {
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

type byLetterLength []Anagrams

func (a byLetterLength) Len() int           { return len(a) }
func (a byLetterLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLetterLength) Less(i, j int) bool { return a[i].Length < a[j].Length }
