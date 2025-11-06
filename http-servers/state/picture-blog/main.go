package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./static/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":9090", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	if r.Method == http.MethodPost {
		f, fh, err := r.FormFile("file")
		checkError(err)
		defer f.Close()
		ext := strings.Split(fh.Filename, ".")[1]
		fmt.Println(ext)
		h := sha1.New()
		io.Copy(h, f)
	}
	err := tpl.ExecuteTemplate(w, "index.html", c)
	checkError(err)

}
func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		uid, err := uuid.NewV6()
		checkError(err)
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
	}
	http.SetCookie(w, c)
	return c
}
func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += fmt.Sprint(c.Value, "|", fname)
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
