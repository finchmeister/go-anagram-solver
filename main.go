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
