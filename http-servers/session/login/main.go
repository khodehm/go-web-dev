package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name     string
	Username string
	Password []byte
}

var tpl *template.Template
var dbSession = map[string]string{}
var dbUser = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}

func main() {
	http.HandleFunc("/", dashboard)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":9090", nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		n := r.FormValue("name")
		p := r.FormValue("password")
		if _, ok := dbUser[un]; ok {
			http.Error(w, "username is not avaliable", http.StatusBadRequest)
			return
		}
		uid, err := uuid.NewV6()
		if err != nil {
			http.Error(w, "failed to generate session id", http.StatusInternalServerError)
		}
		c := &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}
		u := user{n, un, bs}
		dbUser[un] = u
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)

}
func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		u, ok := dbUser[un]
		// check if user exist
		if !ok {
			http.Error(w, "user dos not exist", http.StatusNotFound)
			return
		}
		//check password
		if err := bcrypt.CompareHashAndPassword(u.Password, []byte(p)); err != nil {
			http.Error(w, "wrong username or password", http.StatusInternalServerError)
			return
		}
		// create new session for user
		uid, err := uuid.NewV6()
		if err != nil {
			http.Error(w, "failed to generate session id", http.StatusInternalServerError)
		}
		c := &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)

}
func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	//delete session
	delete(dbSession, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
func dashboard(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	u := getUser(r)
	tpl.ExecuteTemplate(w, "dashboard.html", u)

}
