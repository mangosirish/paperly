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

// Crear un nuevo registro de Person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	// Decodificar el cuerpo de la solicitud en el modelo Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "People" (first_name, middle_name, first_surname, second_surname, personal_email, institutional_email)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING person_id`
	err = db.DB.QueryRow(query, person.FirstName, person.MiddleName, person.FirstSurname, person.SecondSurname, person.PersonalEmail, person.InstitutionalEmail).Scan(&person.PersonID)
	if err != nil {
		http.Error(w, "Error al crear la persona", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// Actualizar un registro de Person
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	// Decodificar el cuerpo de la solicitud en el modelo Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el PersonID esté presente
	if person.PersonID == 0 {
		http.Error(w, "Se requiere un PersonID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "People" 
			  SET first_name = $1, middle_name = $2, first_surname = $3, second_surname = $4, personal_email = $5, institutional_email = $6
			  WHERE person_id = $7`
	result, err := db.DB.Exec(query, person.FirstName, person.MiddleName, person.FirstSurname, person.SecondSurname, person.PersonalEmail, person.InstitutionalEmail, person.PersonID)
	if err != nil {
		http.Error(w, "Error al actualizar la persona", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró la persona con el ID proporcionado", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// Eliminar un registro de Person
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// Obtener el PersonID de los parámetros de la ruta
	vars := mux.Vars(r)
	personID := vars["person_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "People" WHERE person_id = $1`
	result, err := db.DB.Exec(query, personID)
	if err != nil {
		http.Error(w, "Error al eliminar la persona", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró la persona con el ID proporcionado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}