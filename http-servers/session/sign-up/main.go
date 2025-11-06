package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	Name     string
	Username string
}

var tpl *template.Template
var userDb = map[string]user{}
var sessionDb = map[string]string{}

func init() {
	tpl = template.Must(template.ParseFiles("./template/index.html"))
}
func main() {
	http.HandleFunc("/", signup)
	// http.HandleFunc("/signup", nil)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	c, err := r.Cookie("session")
	if err != nil {
		uid, _ := uuid.NewV6()
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
		http.SetCookie(w, c)
	}
	var u user
	if un, ok := sessionDb[c.Value]; ok {
		u = userDb[un]
	}
	if r.Method == http.MethodPost {
		n := r.FormValue("name")
		un := r.FormValue("username")
		u = user{n, un}
		sessionDb[c.Value] = un
		userDb[un] = u
	}
	tpl.ExecuteTemplate(w, "index.html", u)
}

// func createSession(w http.ResponseWriter, r *http.Request) {

// }
