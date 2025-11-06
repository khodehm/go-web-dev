package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// listen and serve on port 8080 and using default serve mux and create following routes and handler for each
// cat
// dog
// /
var templ *template.Template

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	data := struct {
		Title   string
		Message string
	}{
		"Dog",
		"Hello there i am a Dogggggggg",
	}
	if err := templ.ExecuteTemplate(w, "index.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}
func cat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// json.NewEncoder(w).Encode(`<h1>hello iam a garfild</h1>`)
	data := struct {
		Title   string
		Message string
	}{
		"Cat",
		"Hello there i am a cat",
	}
	if err := templ.ExecuteTemplate(w, "index.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}
func init() {
	templ = template.Must(template.ParseFiles("http-servers/router/hands-on-1/index.html"))
}
func main() {
	fmt.Println("server started on port 8090")
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		json.NewEncoder(w).Encode(`<h1>Welecom to my app</h1>`)
	})
	http.ListenAndServe(":8090", nil)
}
