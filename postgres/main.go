package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

// connect to postgres database
// it is local database so you should have postgresql installed
func main() {
	// user := os.Getenv("POSTGRES_USER")
	// db := os.Getenv("POSTGRES_DB")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// host := os.Getenv("HOST")
	// if user == "" || db == "" || password == "" || host == "" {
	// 	log.Fatalln("Please set the environment variables POSTGRES_USER, POSTGRES_DB, POSTGRES_PASSWORD and HOST")
	// }
	// cs := strings.Join([]string{"postgresql://", user, ":", password, "@", host, "/", db, "?sslmode=disable"}, "")

	// connection string
	db, err := sql.Open("postgres", "postgresql://bookstoremanager:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database succesfully connected!")
	SelectFromBooks(db)

}

func SelectFromBooks(d *sql.DB) {
	rows, err := d.Query(`select * from books`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	bks := make([]Book, 0)
	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.isbn, &b.title, &b.author, &b.price)
		if err != nil {
			panic(err)
		}
		bks = append(bks, b)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	for _, bk := range bks {
		fmt.Printf("%#v\n", bk)
	}
}
