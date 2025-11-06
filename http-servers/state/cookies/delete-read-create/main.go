package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<a href="/set">set a cookie</a>`)
}
func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie",
		Value: "valueeeee",
	})
	fmt.Fprint(w, `<a href="/read">read</a>`)
}
func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<div>
		<p>Domain :%v </p>
		<p>cookie:%v </p>
		<a href="/expire">expire</a>
	</div>`, c.Domain, c)
}
func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 //expires now !
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
