package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", correctAnswer)
	// http.HandleFunc("/", setCookies)
	// http.HandleFunc("/visits", countVisits)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

// my answeer
//var visits int
// func countVisits(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("visits")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	fmt.Println(c)
// 	visits++
// 	fmt.Fprintf(w, "number of visits :%v\n", visits)
// }
// func setCookies(w http.ResponseWriter, r *http.Request) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:  "visits",
// 		Value: fmt.Sprint(visits),
// 	})
// }

// correct answer
func correctAnswer(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visits")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w, c)
	fmt.Fprintf(w, "you visted this webtsite %v times \n\n", count)
}
