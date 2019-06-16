package anagramsolver

import (
	"testing"
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

// Performance testing
//func TestAnagrams1(t *testing.T) {
//	GetAnagrams("aar", 3)
//	GetAnagrams("aard", 3)
//	GetAnagrams("aardv", 3)
//	GetAnagrams("aardva", 3)
//	GetAnagrams("aardvar", 3)
//	GetAnagrams("aardvark", 3)
//	GetAnagrams("aardvarks", 3)
//	GetAnagrams("aardvarksa", 3)
//	GetAnagrams("aardvarksab", 3)
//}
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
