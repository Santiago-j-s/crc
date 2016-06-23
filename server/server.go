// Copyright 2016
// Dibez Pablo pdibez@gmail.com
// Santana Santiago santana.santiago@gmail.com
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Page is a representation of a page of a wiki.
type Page struct {
	Title string
	Body  []byte
	//TODO: add js and css
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl+".html", "content/menu.html")
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

	c, err := crc8(poly, msg)
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
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("css/")))
	http.Handle("/css/", css)
	
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("js/")))
	http.Handle("/js/", js)

	http.HandleFunc("/hamming", handlerHamming)
	http.HandleFunc("/analisis", handlerAnalisis)
	http.HandleFunc("/crc", handlerCrc)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":6080", nil))
}
