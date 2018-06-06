package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Cotent Type", "text/html")
		templates := template.New("tempaltes")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"The List of Fruits",
		}

		templates.Lookup("test").Execute(w, context)

	})
	http.ListenAndServe(":8090", nil)
}

type Context struct {
	Fruit [3]string
	Title string
}

const doc = `{{template "header" .Title}}<body><h1>List of Fruits</h1> <ul>{{range .Fruit}} <li>{{.}}</li>{{end}} <ul> </body> `

const header = `<!DOCTYPE HTML> <html> <head><title>Example Title</title> </head>`

const footer = `</html>`
