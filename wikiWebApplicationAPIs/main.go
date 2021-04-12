package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string `json:title`
	Body  string `json:body`
}

type Data struct {
	Success   bool
	ErrorData string
	PageData  Page
}

func (page *Page) save() error {
	pageTitle := page.Title
	filename := pageTitle + ".txt"
	return ioutil.WriteFile(filename, []byte(page.Body), 0644)
}

func getPage(w http.ResponseWriter, r *http.Request) {
	pageTitle := r.URL.Path[len("/getPage/"):]
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		filename := pageTitle + ".txt"
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			e := Data{Success: false, ErrorData: string(err.Error())}
			json.NewEncoder(w).Encode(e)
			return
		}
		page := Page{Title: pageTitle, Body: string(body)}
		data := Data{Success: true, PageData: page}
		json.NewEncoder(w).Encode(data)
	default:
		return
	}
}

func postPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var page Page
		err := json.NewDecoder(r.Body).Decode(&page)
		if err != nil {
			e := Data{Success: false, ErrorData: string(err.Error())}
			json.NewEncoder(w).Encode(e)
			return
		}
		filepath := page.Title + ".txt"
		if _, err := os.Stat(filepath); os.IsNotExist(err) == false {
			// If the file already exists
			e := Data{Success: false, ErrorData: "This page already exists and if you want to change it please use the `editPage` API"}
			json.NewEncoder(w).Encode(e)
			return
		}
		err = page.save()
		if err != nil {
			e := Data{Success: false, ErrorData: string(err.Error())}
			json.NewEncoder(w).Encode(e)
			return
		}
		data := Data{Success: true, PageData: page}
		json.NewEncoder(w).Encode(data)
	default:
		return
	}
}

func editPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var page Page
		err := json.NewDecoder(r.Body).Decode(&page)
		if err != nil {
			e := Data{Success: false, ErrorData: string(err.Error())}
			json.NewEncoder(w).Encode(e)
			return
		}
		err = page.save()
		if err != nil {
			e := Data{Success: false, ErrorData: string(err.Error())}
			json.NewEncoder(w).Encode(e)
			return
		}
		data := Data{Success: true, PageData: page}
		json.NewEncoder(w).Encode(data)
	default:
		return
	}
}

func main() {
	http.HandleFunc("/getPage/", getPage)
	http.HandleFunc("/postPage/", postPage)
	http.HandleFunc("/editPage/", editPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
