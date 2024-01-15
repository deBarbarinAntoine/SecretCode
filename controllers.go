package SecretCode

import (
	"fmt"
	"log"
	"net/http"
)

var Input string

// rootHandler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// Index page handler.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Input)
	err := tmpl.ExecuteTemplate(w, "index", Input)
	if err != nil {
		log.Fatal(err)
	}
}

func decode(str string) string {
	var result string
	for _, char := range str {
		char += 11
		char += 13
		char -= 712 % 'M'
		char -= 3
		result += string(char)
	}
	return result
}

// Treatment handler.
func treatmentHandler(w http.ResponseWriter, r *http.Request) {
	Input = r.FormValue("input")
	fmt.Println(Input)
	Input = decode(Input)
	fmt.Println(Input)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
