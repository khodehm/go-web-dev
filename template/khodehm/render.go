package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/khodehm/static/*"))
}
func main() {
	err := temp.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// func RenderHTMLTemplate(htmlPath string, data any) (string, error) {
// 	tmpl, err := template.ParseFS(templateFiles, htmlPath)
// 	if err != nil {
// 		fmt.Printf("error while parsing HTML file %s", err)
// 	}
// 	var body bytes.Buffer
// 	if err := tmpl.Execute(&body, data); err != nil {
// 		return "", err
// 	}
// 	return body.String(), nil
// }

// func CreateTempalte(name string, html string, data any) (string, error) {
// 	tmpl, err := template.New(name).Parse(html)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to render HTML template: %w", err)
// 	}
// 	var buf bytes.Buffer
// 	err = tmpl.Execute(&buf, data)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to execute template: %w", err)
// 	}
// 	return buf.String(), nil
// }
