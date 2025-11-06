package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Form)
	http.ListenAndServe(":9090", nil)
}
func Form(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("data")
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<form method="get">
				<input type="text" name="data"/>
				<input type="submit" value="Submit"/>
			</form>
			%v
		`, q)
	}
}
