package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgresql://bookstoremanager:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database succesfully connected!")
	tpl = template.Must(template.ParseGlob("template/*"))
}

type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func main() {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./assets/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("GET /books/create", createBookPage)
	mux.HandleFunc("POST /books/create/process", createBookProcess)
	mux.HandleFunc("/", getBooks)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("GET /books/show", getBook)
	http.ListenAndServe(":9090", mux)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	books := make([]Book, 0)

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		bk := Book{}
		err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, bk)
	}
	err = tpl.ExecuteTemplate(w, "index.html", books)
	if err != nil {
		http.Error(w, "failed to excute template", http.StatusInternalServerError)
	}
}
func getBook(w http.ResponseWriter, r *http.Request) {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return

	}
	row := db.QueryRow("select * from books where isbn = $1", isbn)
	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.Error(w, "no record", http.StatusNotFound)
		return
	case err != nil:
		http.Error(w, row.Err().Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, bk)
}
func createBookPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create-book.html", nil)
}
func createBookProcess(w http.ResponseWriter, r *http.Request) {
	b := Book{}
	b.Isbn = r.FormValue("isbn")
	b.Title = r.FormValue("title")
	b.Author = r.FormValue("author")
	p, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	b.Price = float32(p)
	if b.Author == " " || b.Isbn == " " || b.Title == " " {
		http.Error(w, "misiing reqired fields", http.StatusBadRequest)
		return
	}
	_, err := db.Exec("INSERT INTO books (isbn, title, author, price) VALUES($1,$2,$3,$4)", b.Isbn, b.Title, b.Author, b.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "create-book.html", nil)

}
