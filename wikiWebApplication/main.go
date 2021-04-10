package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var savePath string = "SavesPagesData"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() (err error) {
	filename := p.Title + ".txt"
	err = ioutil.WriteFile(filename, p.Body, 0644)
	return err
}

func loadPage(title string) (p *Page, err error) {
	filname := title + ".txt"
	body, err := ioutil.ReadFile(filname)
	return &Page{Title: title, Body: []byte(body)}, err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "404 Page Not Found")
	}
	viewTemplate, _ := template.ParseFiles("view.html")
	viewTemplate.Execute(w, page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	editTemplate, _ := template.ParseFiles("edit.html")
	editTemplate.Execute(w, page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	page.save()
}

func main() {
	// p := &Page{Title: "Testing", Body: []byte("This is just a test doc!")}
	// p.save()
	// loadedPage, _ := loadPage("Testing")
	// fmt.Println(string(loadedPage.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
