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

// Crear un nuevo registro de Institution
func CreateInstitution(w http.ResponseWriter, r *http.Request) {
	var institution models.Institution

	// Decodificar el cuerpo de la solicitud en el modelo Institution
	err := json.NewDecoder(r.Body).Decode(&institution)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "Institutions" (name)
			  VALUES ($1) RETURNING institution_id`
	err = db.DB.QueryRow(query, institution.Name).Scan(&institution.InstitutionID)
	if err != nil {
		http.Error(w, "Error al crear la institución", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(institution)
}

// Actualizar un registro de Institution
func UpdateInstitution(w http.ResponseWriter, r *http.Request) {
	var institution models.Institution

	// Decodificar el cuerpo de la solicitud en el modelo Institution
	err := json.NewDecoder(r.Body).Decode(&institution)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el InstitutionID esté presente
	if institution.InstitutionID == 0 {
		http.Error(w, "Se requiere un InstitutionID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "Institutions" 
			  SET name = $1
			  WHERE institution_id = $2`
	result, err := db.DB.Exec(query, institution.Name, institution.InstitutionID)
	if err != nil {
		http.Error(w, "Error al actualizar la institución", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró la institución con el ID proporcionado", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(institution)
}

// Eliminar un registro de Institution
func DeleteInstitution(w http.ResponseWriter, r *http.Request) {
	// Obtener el InstitutionID de los parámetros de la ruta
	vars := mux.Vars(r)
	institutionID := vars["institution_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "Institutions" WHERE institution_id = $1`
	result, err := db.DB.Exec(query, institutionID)
	if err != nil {
		http.Error(w, "Error al eliminar la institución", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró la institución con el ID proporcionado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}

