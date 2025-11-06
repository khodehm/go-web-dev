package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("template_files/tmplate2.txt"))
}
func main() {
	err := temp.ExecuteTemplate(os.Stdout, "tmplate2.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
