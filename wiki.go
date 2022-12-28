package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// holds wiki page information
type Page struct {
	Title string
	Body  []byte
}

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600) // 0600 -> only owner has full read/write access
// }

// loads page from local filesystem into a Page struct object
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil // return a pointer to the newly created page struct
}

// function that handles serving up views to the user
func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// loads page and displays an HTML form
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, `
		<div style="text-align: center;">
			<h1>Editing page: %s</h1>
			<form action="/save/%s" method="POST">
				<textarea style="height: 400px; width: 500px;" name="body">%s</textarea>
				<br>
				<br>
				<input type="submit" value="save">
			</form>
		</div>
	`, p.Title, p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
