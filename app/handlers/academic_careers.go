package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
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

func GetAcademicCareersByName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    query := `SELECT career_id, name FROM "AcademicCareers" WHERE name = $1`
    rows, err := db.DB.Query(query, name)
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

// Crear un nuevo registro de AcademicCareer
func CreateAcademicCareer(w http.ResponseWriter, r *http.Request) {
	var career models.AcademicCareer

	// Decodificar el cuerpo de la solicitud en el modelo AcademicCareer
	err := json.NewDecoder(r.Body).Decode(&career)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "AcademicCareers" (name) VALUES ($1) RETURNING career_id`
	err = db.DB.QueryRow(query, career.Name).Scan(&career.CareerID)
	if err != nil {
		http.Error(w, "Error al crear el registro de carrera académica", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(career)
}

// Actualizar un registro de AcademicCareer
func UpdateAcademicCareer(w http.ResponseWriter, r *http.Request) {
	var career models.AcademicCareer

	// Decodificar el cuerpo de la solicitud en el modelo AcademicCareer
	err := json.NewDecoder(r.Body).Decode(&career)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el CareerID esté presente
	if career.CareerID == 0 {
		http.Error(w, "Se requiere un CareerID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "AcademicCareers" SET name = $1 WHERE career_id = $2`
	result, err := db.DB.Exec(query, career.Name, career.CareerID)
	if err != nil {
		http.Error(w, "Error al actualizar el registro de carrera académica", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el registro de carrera académica", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(career)
}

// Eliminar un registro de AcademicCareer
func DeleteAcademicCareer(w http.ResponseWriter, r *http.Request) {
	// Obtener el CareerID de los parámetros de la ruta
	vars := mux.Vars(r)
	careerID := vars["career_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "AcademicCareers" WHERE career_id = $1`
	result, err := db.DB.Exec(query, careerID)
	if err != nil {
		http.Error(w, "Error al eliminar el registro de carrera académica", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el registro de carrera académica", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}