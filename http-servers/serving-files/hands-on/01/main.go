package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

/*
ListenAndServe on port 8080 of localhost
For the default route "/" Have a func called "foo" which writes to the response "foo ran"

For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

This is from dog
" and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
*/

var temp *template.Template

func main() {

	log.Fatal(http.ListenAndServe(":9090", nil))
}
func init() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogimg)
	temp = template.Must(template.ParseFiles("http-servers/serving-files/hands-on/01/index.html"))
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type:", "text/html")
	w.WriteHeader(200)
	d := struct {
		Title       string
		Description string
	}{
		Title:       "Bark",
		Description: "hello there i am a dog",
	}
	err := temp.ExecuteTemplate(w, "index.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	io.WriteString(w, "foo rann !!!!!!!!!")
}
func dogimg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, `http-servers/serving-files/sharefile.webp`)
}
