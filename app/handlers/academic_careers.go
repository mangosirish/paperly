package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
)

func GetAcademicCareers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query(`SELECT career_id, name FROM "AcademicCareers"`)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las carreras académicas", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    careers := []models.AcademicCareer{}
    for rows.Next() {
        var career models.AcademicCareer
        if err := rows.Scan(&career.CareerID, &career.Name); err != nil {
            log.Printf("Error al escanear las carreras académicas: %v\n", err)
            http.Error(w, "Error al escanear las carreras académicas", http.StatusInternalServerError)
            return
        }
        careers = append(careers, career)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(careers)
}