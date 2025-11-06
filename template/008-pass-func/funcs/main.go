package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tmp *template.Template
var fm = template.FuncMap{
	"lc": stringLen,
	"tf": timeFormatter,
}

func stringLen(s ...string) string {
	st := strings.TrimSpace(strings.Join(s, ""))
	return fmt.Sprintf("character length %d\n", len(st))
}
func timeFormatter() time.Time {
	t := time.Now()
	t.Format(time.DateTime)
	return t
}

func init() {
	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("template_files/tmplate3.txt"))
}
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "tmplate3.txt", "Hello World!")
	if err != nil {
		log.Fatalln(err)
	}
}
