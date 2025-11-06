package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name  string
	Skils map[string]string
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("template_files/*"))
}
func main() {
	f, err := os.Create("template/007-pass-variable/aggregate-types/log.html")
	if err != nil {
		log.Fatalln(err)
	}
	m := map[string]string{"name": "alireza hm"}
	// map in template
	err = tmp.ExecuteTemplate(f, "tmplate4.txt", m)
	if err != nil {
		log.Fatalln(err)

	}
	s := []string{"hello", "iam", "alireza hm"}
	// slice of string in template
	err = tmp.ExecuteTemplate(f, "tmplate5.txt", s)
	if err != nil {
		log.Fatalln(err)

	}

	st := person{
		Name:  "alitezahm",
		Skils: map[string]string{"go programming": "i have knowledge of go and master go skils"},
	}
	err = tmp.ExecuteTemplate(f, "tmplate1.txt", st)
	if err != nil {
		log.Fatalln(err)

	}
}
