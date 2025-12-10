package config

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgresql://bookstoremanager:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Database succesfully connected!")
	fmt.Println("server running on http://localhost:9090")

}
