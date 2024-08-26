package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/handlers"
)

func main() {
    db.InitDB()
    defer db.CloseDB()

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