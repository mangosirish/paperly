package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
)

func GetFaculties(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query(`SELECT faculty_id, name FROM "Faculties"`)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las facultades", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    faculties := []models.Faculty{}
    for rows.Next() {
        var faculty models.Faculty
        if err := rows.Scan(&faculty.FacultyID, &faculty.Name); err != nil {
            log.Printf("Error al escanear las facultades: %v\n", err)
            http.Error(w, "Error al escanear las facultades", http.StatusInternalServerError)
            return
        }
        faculties = append(faculties, faculty)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(faculties)
}