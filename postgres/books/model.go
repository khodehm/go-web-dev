package books

import (
	"github.com/khodehm/pg/config"
)

type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func GetBook(isbn string) (Book, error) {
	row := config.DB.QueryRow("select * from books where isbn = $1", isbn)
	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return Book{}, err
	}
	return bk, nil
}
func GetAllBooks() ([]Book, error) {
	books := make([]Book, 0)
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bk := Book{}
		err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, bk)
	}
	return books, nil
}
func UpdateBook(bk Book) error {
	_, err := config.DB.Exec("UPDATE books SET isbn=$1, title=$2,author=$3,price=$4 WHERE isbn=$1", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return err
	}
	return nil
}
func DeleteBook(isbn string) error {
	_, err := config.DB.Exec("DELETE From books WHERE isbn=$1", isbn)
	if err != nil {
		return err
	}
	return nil
}
