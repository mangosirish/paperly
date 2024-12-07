package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
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

func GetFacultiesByName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    query := `SELECT faculty_id, name FROM "Faculties" WHERE name = $1`
    rows, err := db.DB.Query(query, name)
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

// Crear un nuevo registro de Faculty
func CreateFaculty(w http.ResponseWriter, r *http.Request) {
	var faculty models.Faculty

	// Decodificar el cuerpo de la solicitud en el modelo Faculty
	err := json.NewDecoder(r.Body).Decode(&faculty)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "Faculties" (name)
			  VALUES ($1) RETURNING faculty_id`
	err = db.DB.QueryRow(query, faculty.Name).Scan(&faculty.FacultyID)
	if err != nil {
		http.Error(w, "Error al crear la facultad", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(faculty)
}

// Actualizar un registro de Faculty
func UpdateFaculty(w http.ResponseWriter, r *http.Request) {
	var faculty models.Faculty

	// Decodificar el cuerpo de la solicitud en el modelo Faculty
	err := json.NewDecoder(r.Body).Decode(&faculty)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el FacultyID esté presente
	if faculty.FacultyID == 0 {
		http.Error(w, "Se requiere un FacultyID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "Faculties" 
			  SET name = $1
			  WHERE faculty_id = $2`
	result, err := db.DB.Exec(query, faculty.Name, faculty.FacultyID)
	if err != nil {
		http.Error(w, "Error al actualizar la facultad", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró la facultad con el ID proporcionado", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(faculty)
}

// Eliminar un registro de Faculty
func DeleteFaculty(w http.ResponseWriter, r *http.Request) {
	// Obtener el FacultyID de los parámetros de la ruta
	vars := mux.Vars(r)
	facultyID := vars["faculty_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "Faculties" WHERE faculty_id = $1`
	result, err := db.DB.Exec(query, facultyID)
	if err != nil {
		http.Error(w, "Error al eliminar la facultad", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró la facultad con el ID proporcionado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}