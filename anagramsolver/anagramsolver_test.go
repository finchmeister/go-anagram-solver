package anagramsolver

import "testing"

func TestEightLetterAnagramYieldsSixWordLengths(t *testing.T) {
	anagrams := GetAnagrams("lindoite", 3)
	if len(anagrams) != 6 {
		t.Errorf("Expected 6 different word lengths anagrams got %d", len(anagrams))
	}
}

func TestAnagramsFoundAardvark(t *testing.T) {
	allAnagrams := GetAnagrams("aardvark", 3)
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
