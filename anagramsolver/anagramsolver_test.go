package anagramsolver

import (
	"fmt"
	"testing"
	"time"
)

func TestEightLetterAnagramYieldsSixWordLengthsAlgo1(t *testing.T) {
	a := NewAnagramSolver(true, false)

	anagrams := a.GetAnagramsAlgo1("lindoite", 3)
	if len(anagrams) != 6 {
		t.Errorf("Expected 6 different word lengths anagrams got %d", len(anagrams))
	}
}

func TestAnagramsFoundAardvarkAlgo1(t *testing.T) {
	a := NewAnagramSolver(true, false)

	allAnagrams := a.GetAnagramsAlgo1("aardvark", 3)
	if len(allAnagrams) != 4 {
		t.Errorf("Expected 4 different word lengths anagrams got %d", len(allAnagrams))
	}

	expectedAnagrams := [][]string{
		{"aardvark"},
		{"arara", "radar", "varda"},
		{"adar", "akra", "arad", "arar", "avar", "dark", "darr", "kava", "raad", "rada", "vara"},
		{"ada", "aka", "ara", "ark", "ava", "dak", "dar", "kra", "rad"},
	}

	for i, anagrams := range allAnagrams {
		for j, word := range anagrams.Words {
			if word != expectedAnagrams[i][j] {
				t.Errorf("Expected %s. Got %s", word, expectedAnagrams[i][j])
			}
		}
	}
}
func TestEightLetterAnagramYieldsSixWordLengthsAlgo3(t *testing.T) {
	a := NewAnagramSolver(true, false)

	anagrams := a.GetAnagramsAlgo3("lindoite", 3)
	if len(anagrams) != 6 {
		t.Errorf("Expected 6 different word lengths anagrams got %d", len(anagrams))
	}
}

func TestAnagramsFoundAardvarkAlgo3(t *testing.T) {
	a := NewAnagramSolver(true, false)

	allAnagrams := a.GetAnagramsAlgo3("aardvark", 3)
	if len(allAnagrams) != 4 {
		t.Errorf("Expected 4 different word lengths anagrams got %d", len(allAnagrams))
	}

	expectedAnagrams := [][]string{
		{"aardvark"},
		{"arara", "radar", "varda"},
		{"adar", "akra", "arad", "arar", "avar", "dark", "darr", "kava", "raad", "rada", "vara"},
		{"ada", "aka", "ara", "ark", "ava", "dak", "dar", "kra", "rad"},
	}

	for i, anagrams := range allAnagrams {
		for j, word := range anagrams.Words {
			if word != expectedAnagrams[i][j] {
				t.Errorf("Expected %s. Got %s", word, expectedAnagrams[i][j])
			}
		}
	}
}

func benchAnagram(letters string) {
	a := NewAnagramSolver(true, false)

	start := time.Now()
	a.GetAnagramsAlgo3(letters, 3)
	fmt.Printf("%f \n", float64(time.Since(start).Nanoseconds())/(float64(time.Millisecond)/float64(time.Nanosecond)))
	//fmt.Printf("%d \n", time.Since(start).Nanoseconds())
}

// Performance testing
func TestAnagrams1(t *testing.T) {
	benchAnagram("aar")
	benchAnagram("aard")
	benchAnagram("aardv")
	benchAnagram("aardva")
	benchAnagram("aardvar")
	benchAnagram("aardvark")
	benchAnagram("aardvarks")
	benchAnagram("aardvarksa")
	benchAnagram("aardvarksab")
}

//
//func TestAnagrams2(t *testing.T) {
//	GetAnagrams2("aar", 3)
//	GetAnagrams2("aard", 3)
//	GetAnagrams2("aardv", 3)
//	GetAnagrams2("aardva", 3)
//	GetAnagrams2("aardvar", 3)
//	GetAnagrams2("aardvark", 3)
//	GetAnagrams2("aardvarks", 3)
//	GetAnagrams2("aardvarksa", 3)
//	GetAnagrams2("aardvarksab", 3)
//}
