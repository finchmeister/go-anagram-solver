package cloudfunction

import (
	"fmt"
	"github.com/finchmeister/go-anagram-solver/anagramsolver"
	"net/http"
)

func HelloYou(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")

	if len(q) > 9 {
		q = q[:9]
	}

	anagrams := anagramsolver.GetAnagrams(q, 3)

	if q == "" {
		fmt.Fprint(w, "Hello, you!")
		return
	}
	fmt.Fprintf(w, anagrams[0].Words[0])
	//fmt.Fprintf(w, "Hello, %s!", html.EscapeString(q))
}
