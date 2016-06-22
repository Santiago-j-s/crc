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

func handlerAnalisis(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "Análisis de Detección"}
	renderTemplate(w, "content/analisis", page)
}

func handlerCrc(w http.ResponseWriter, r *http.Request) {
	poly := r.FormValue("poly")
	msg := r.FormValue("msg")

	c, err := crc(poly, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, c)
}

func handlerHamming(w http.ResponseWriter, r *http.Request) {
	poly := r.FormValue("poly")

	h, err := hamming(poly)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, h)
}

func main() {
	http.HandleFunc("/hamming", handlerHamming)
	http.HandleFunc("/analisis", handlerAnalisis)
	http.HandleFunc("/crc", handlerCrc)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":6080", nil))
}
