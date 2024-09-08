package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
)

func GetSpecialties(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query(`SELECT specialty_id, name FROM "Specialties"`)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las especialidades", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    specialties := []models.Specialty{}
    for rows.Next() {
        var specialty models.Specialty
        if err := rows.Scan(&specialty.SpecialtyID, &specialty.Name); err != nil {
            log.Printf("Error al escanear las especialidades: %v\n", err)
            http.Error(w, "Error al escanear las especialidades", http.StatusInternalServerError)
            return
        }
        specialties = append(specialties, specialty)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(specialties)
}

func GetSpecialtiesByName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    query := `SELECT specialty_id, name FROM "Specialties" WHERE name = $1`
    rows, err := db.DB.Query(query, name)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las especialidades", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    specialties := []models.Specialty{}
    for rows.Next() {
        var specialty models.Specialty
        if err := rows.Scan(&specialty.SpecialtyID, &specialty.Name); err != nil {
            log.Printf("Error al escanear las especialidades: %v\n", err)
            http.Error(w, "Error al escanear las especialidades", http.StatusInternalServerError)
            return
        }
        specialties = append(specialties, specialty)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(specialties)
}