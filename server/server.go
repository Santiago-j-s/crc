package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Page is a representation of a page of a wiki.
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "CRC"}
	renderTemplate(w, "content/index", page)
}

func handlerCrc(w http.ResponseWriter, r *http.Request) {
	poly := r.FormValue("poly")
	msg := r.FormValue("msg")
	res := crc(poly, msg)
	fmt.Fprintf(w, res)
}

func main() {
	http.HandleFunc("/crc", handlerCrc)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":6080", nil))
}
