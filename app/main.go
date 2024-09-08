package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mangosirish/paperly/db"
	"github.com/mangosirish/paperly/handlers"
	"github.com/mangosirish/paperly/templates"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	// Frontend
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.IndexView().Render(r.Context(), w)
	})
	http.HandleFunc("/web-table-view", func(w http.ResponseWriter, r *http.Request) {
		templates.TableView(r.URL.Query().Get("table")).Render(r.Context(), w)
	})

	http.HandleFunc("/web-articles", func(w http.ResponseWriter, r *http.Request) {
		templates.TableComponent("Artricles").Render(r.Context(), w)
	})

	// API
	http.HandleFunc("/articles", handlers.GetArticles)
	http.HandleFunc("/people", handlers.GetPeople)
	http.HandleFunc("/authors", handlers.GetAuthors)
	http.HandleFunc("/academic-careers", handlers.GetAcademicCareers)
	http.HandleFunc("/faculties", handlers.GetFaculties)
	http.HandleFunc("/institutions", handlers.GetInstitutions)
	http.HandleFunc("/student-social-services", handlers.GetStudentSocialServices)
	http.HandleFunc("/journals", handlers.GetJournals)
	http.HandleFunc("/specialties", handlers.GetSpecialties)

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
