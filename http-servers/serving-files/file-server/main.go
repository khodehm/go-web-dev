package main

import (
	"log"
	"net/http"
)

type img string

func main() {
	http.Handle("/", http.FileServer(http.Dir("http-servers/serving-files/")))
	log.Fatal(http.ListenAndServe(":8090", nil))
}

// func (i img) image(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, `<img src="https://images.unsplash.com/photo-1506744038136-46273834b3fb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80" alt="random image from unsplash">`)
// 	// io.WriteString(w.Header(),"")
// }
