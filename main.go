package main

import (
	"github.com/finchmeister/go-anagram-solver/anagramsolver"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	indexTemplate = template.Must(template.ParseFiles("template/index.html"))
	//debug         = false
)

func main() {
	anagramsolver.GetAnagrams("aar", 3)
	anagramsolver.GetAnagrams("aard", 3)
	anagramsolver.GetAnagrams("aardv", 3)
	anagramsolver.GetAnagrams("aardva", 3)
	anagramsolver.GetAnagrams("aardvar", 3)
	anagramsolver.GetAnagrams("aardvark", 3)
	anagramsolver.GetAnagrams("aardvarks", 3)
	anagramsolver.GetAnagrams("aardvarksa", 3)
	anagramsolver.GetAnagrams("aardvarksab", 3)

	return

	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	q := r.FormValue("q")
	//printDebug(q)

	if len(q) > 9 {
		q = q[:9]
	}

	anagrams := anagramsolver.GetAnagrams(q, 3)

	timeTaken := int(time.Since(start).Nanoseconds() / 1000000)
	params := TemplateFields{Query: q, Anagrams: anagrams, TimeTaken: timeTaken}

	indexTemplate.Execute(w, params)
}

type TemplateFields struct {
	Query     string
	Anagrams  []anagramsolver.Anagrams
	TimeTaken int
}
