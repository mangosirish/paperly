package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query(`SELECT author_id, notes, person_id FROM "Authors"`)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener los autores", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    authors := []models.Author{}
    for rows.Next() {
        var author models.Author
        if err := rows.Scan(&author.AuthorID, &author.Notes, &author.PersonID); err != nil {
            log.Printf("Error al escanear los autores: %v\n", err)
            http.Error(w, "Error al escanear los autores", http.StatusInternalServerError)
            return
        }
        authors = append(authors, author)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(authors)
}