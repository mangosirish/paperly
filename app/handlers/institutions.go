package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
)

func GetInstitutions(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query(`SELECT institution_id, name FROM "Institutions"`)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las instituciones", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    institutions := []models.Institution{}
    for rows.Next() {
        var institution models.Institution
        if err := rows.Scan(&institution.InstitutionID, &institution.Name); err != nil {
            log.Printf("Error al escanear las instituciones: %v\n", err)
            http.Error(w, "Error al escanear las instituciones", http.StatusInternalServerError)
            return
        }
        institutions = append(institutions, institution)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(institutions)
}

func GetInstitutionsByName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    query := `SELECT institution_id, name FROM "Institutions" WHERE name = $1`
    rows, err := db.DB.Query(query, name)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las instituciones", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    institutions := []models.Institution{}
    for rows.Next() {
        var institution models.Institution
        if err := rows.Scan(&institution.InstitutionID, &institution.Name); err != nil {
            log.Printf("Error al escanear las instituciones: %v\n", err)
            http.Error(w, "Error al escanear las instituciones", http.StatusInternalServerError)
            return
        }
        institutions = append(institutions, institution)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(institutions)
}