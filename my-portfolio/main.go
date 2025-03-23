package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/about", aboutHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "projects.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func projectsAPIHandler(w http.ResponseWriter, r *http.Request) {
	projects := []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Link        string `json:"link"`
	}{
		{"Project 1", "A cool project", "https://example.com/project1"},
		{"Project 2", "Another awesome project", "https://example.com/project2"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
