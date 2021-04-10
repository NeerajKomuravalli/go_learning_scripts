package main

import (
	"fmt"
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
	fmt.Fprintf(w, string(page.Body))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		page.Title, page.Title, page.Body)
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
