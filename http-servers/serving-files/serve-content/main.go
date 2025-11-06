package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type img string

func main() {
	var i img
	http.HandleFunc("/", i.image)
	http.HandleFunc("/loaded", i.serveimg)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalln(err)
	}
}

func (i img) image(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<img src="https://images.unsplash.com/photo-1506744038136-46273834b3fb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80" alt="random image from unsplash">`)
	// io.WriteString(w.Header(),"")
}
func (i img) serveimg(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("http-servers/serving-files/sharefile.webp")
	if err != nil {
		http.Error(w, "something went wrong", 500)
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "something went wrong", 500)
	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
