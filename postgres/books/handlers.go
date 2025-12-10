package books

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/khodehm/pg/config"
	errormeessage "github.com/khodehm/pg/error"
)

func RegisterRoutes(mux *http.ServeMux) {
	fileserver := http.FileServer(http.Dir("./assets/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("GET /books/create", CreateBookPage)
	mux.HandleFunc("POST /books/create/process", CreateBookProcess)
	mux.HandleFunc("GET /book/edit", EditBookPage)
	mux.HandleFunc("POST /book/edit/process", EditBookProcess)
	mux.HandleFunc("GET /book/delete", DeleteBookProcess)
	mux.HandleFunc("/", GetBooks)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("GET /books/show", GetBookById)
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = config.Tpl.ExecuteTemplate(w, "index.html", books)
	if err != nil {
		http.Error(w, "failed to excute template", http.StatusInternalServerError)
	}
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return

	}
	row := config.DB.QueryRow("select * from books where isbn = $1", isbn)
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
func CreateBookPage(w http.ResponseWriter, r *http.Request) {
	config.Tpl.ExecuteTemplate(w, "create-book.html", nil)
}
func EditBookPage(w http.ResponseWriter, r *http.Request) {
	i := r.FormValue("isban")
	if i == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	b, err := GetBook(i)
	if err != nil {
		e := errormeessage.Error{
			Error:       "کتاب مورد نظر یافت نشد",
			Description: "کتابی با این شناسه در سیستم ثبت نشده است لطفا شناسه را بررسی کنید",
		}
		config.Tpl.ExecuteTemplate(w, "error.html", e)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "edit-book.html", b)
}
func CreateBookProcess(w http.ResponseWriter, r *http.Request) {
	b := Book{}
	b.Isbn = r.FormValue("isbn")
	b.Title = r.FormValue("title")
	b.Author = r.FormValue("author")
	p, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	b.Price = float32(p)
	if b.Author == " " || b.Isbn == " " || b.Title == " " {
		e := errormeessage.Error{
			Error:       "کتاب ایجاد نشد!",
			Description: "پارامتر های فرم به درستی تکمیل نشده!",
		}
		config.Tpl.ExecuteTemplate(w, "error.html", e)
		return
	}
	_, err := config.DB.Exec("INSERT INTO books (isbn, title, author, price) VALUES($1,$2,$3,$4)", b.Isbn, b.Title, b.Author, b.Price)
	if err != nil {
		e := errormeessage.Error{
			Error:       "کتاب ایجاد نشد",
			Description: err.Error(),
		}
		config.Tpl.ExecuteTemplate(w, "error.html", e)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func EditBookProcess(w http.ResponseWriter, r *http.Request) {
	b := Book{}
	b.Isbn = r.FormValue("isbn")
	b.Title = r.FormValue("title")
	b.Author = r.FormValue("author")
	p, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	b.Price = float32(p)
	if b.Author == " " || b.Isbn == " " || b.Title == " " {
		e := errormeessage.Error{
			Error:       "کتاب ویرای نشد!",
			Description: "پارامتر های فرم به درستی تکمیل نشده!",
		}
		config.Tpl.ExecuteTemplate(w, "error.html", e)
		return
	}
	err := UpdateBook(b)
	if err != nil {
		e := errormeessage.Error{
			Error:       "کتاب ویرایش نشد",
			Description: err.Error(),
		}
		config.Tpl.ExecuteTemplate(w, "error.html", e)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func DeleteBookProcess(w http.ResponseWriter, r *http.Request) {
	i := r.FormValue("isban")
	if i == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	err := DeleteBook(i)
	if err != nil {
		e := errormeessage.Error{
			Error:       "کتاب ایجاد نشد",
			Description: err.Error(),
		}
		config.Tpl.ExecuteTemplate(w, "error.html", e)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
