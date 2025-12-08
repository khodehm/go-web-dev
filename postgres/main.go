package main

import (
	"net/http"

	"github.com/khodehm/pg/books"
)

func main() {
	mux := http.NewServeMux()
	books.RegisterRoutes(mux)
	http.ListenAndServe(":9090", mux)
}
