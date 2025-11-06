package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tl *template.Template

func init() {
	tl = template.Must(template.ParseGlob("template/005-parse-glob/template/*"))

}
func main() {
	err := tl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tl.ExecuteTemplate(os.Stdout, "tmplate5.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create("template/005-parse-glob/out.txt")
	if err != nil {
		log.Fatalln(err)
	}
	err = tl.ExecuteTemplate(f, "tmplate5.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tl.Execute(f, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tl.Name())
}
