package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
)

func SearchAcademicProfiles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryParam := vars["query"]

	if queryParam == "" {
		http.Error(w, "Parámetro de búsqueda 'query' requerido", http.StatusBadRequest)
		return
	}

	query := `
        SELECT academic_profile_id, institution_id, faculty_id, career_id, specialty_id
        FROM "AcademicProfiles"
        WHERE CAST(institution_id AS TEXT) ILIKE '%' || $1 || '%'
           OR CAST(faculty_id AS TEXT) ILIKE '%' || $1 || '%'
           OR CAST(career_id AS TEXT) ILIKE '%' || $1 || '%'
           OR CAST(specialty_id AS TEXT) ILIKE '%' || $1 || '%'
    `

	rows, err := db.DB.Query(query, queryParam)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al buscar perfiles académicos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	profiles := []models.AcademicProfile{}
	for rows.Next() {
		var profile models.AcademicProfile
		if err := rows.Scan(&profile.AcademicProfileID, &profile.InstitutionID, &profile.FacultyID, &profile.CareerID, &profile.SpecialtyID); err != nil {
			log.Printf("Error al escanear los perfiles académicos: %v\n", err)
			http.Error(w, "Error al escanear los perfiles académicos", http.StatusInternalServerError)
			return
		}
		profiles = append(profiles, profile)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}