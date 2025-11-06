package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("template/005-parse-glob/template/*"))
}
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "tmplate5.txt", nil)
	fmt.Println()
	if err != nil {
		log.Fatalln(err)
	}
	err = tmp.ExecuteTemplate(os.Stdout, "tmplate4.txt", `<h1 style="color:blue;">hello there</h1>`)
	fmt.Println()
	if err != nil {
		log.Fatalln(err)
	}
}
