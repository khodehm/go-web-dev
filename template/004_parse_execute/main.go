package main

import (
	"log"
	"os"
	"text/template"
)

/*
func ParseFiles(filenames ...string) (*Template, error)
ParseFiles creates a new [Template] and parses the template definitions from the named files. The returned template's name will have the base name

	and parsed contents of the first file. There must be at least one file. If an error occurs, parsing stops and the returned *Template is nil.

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results. For instance,
ParseFiles("a/foo", "b/foo") stores "b/foo" as the template named "foo", while "a/foo" is unavailable.
*/

func main() {
	tpl, err := template.ParseFiles("template/using-package-template/004_parse_execute/templ.txt")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(os.Stdout, nil)
}
