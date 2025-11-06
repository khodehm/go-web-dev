package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmpl *template.Template

type handler string

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// parse data
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// create a aggregate type for passign data to template
	data := struct {
		Method string
		Form   url.Values
		Header http.Header
		Host   string
	}{
		r.Method,
		r.Form,
		r.Header,
		r.Host,
	}
	// execute html template and pass in data
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// load the folder with static data
	tmpl = template.Must(template.ParseGlob("http-servers/request/form/static/*"))
}

func main() {
	var h handler
	fmt.Println("server is started on port 9090")
	err := http.ListenAndServe(":9090", h)
	if err != nil {
		log.Fatal(err)
	}
}
