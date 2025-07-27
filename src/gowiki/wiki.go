package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string `json:"title"`
	Body  []byte `json:"body"`
}

func (p *Page) Save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view"):]
	page, err := LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editorHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := LoadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, "edit", page)
}

func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, page)
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
		return
	}
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a test page.")}
	// p1.Save()

	// p2, _ := LoadPage("TestPage")
	// fmt.Println("Title:", p2.Title)
	// fmt.Println("Body:", string(p2.Body))

	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/edit/", editorHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
