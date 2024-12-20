package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mangosirish/paperly/db"
	"github.com/mangosirish/paperly/models"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email FROM "People"`)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func SearchPeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryParam := vars["query"]

	if queryParam == "" {
		http.Error(w, "Parámetro de búsqueda 'query' requerido", http.StatusBadRequest)
		return
	}

	query := `
        SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
        FROM "People" 
        WHERE first_name ILIKE '%' || $1 || '%' 
           OR middle_name ILIKE '%' || $1 || '%' 
           OR first_surname ILIKE '%' || $1 || '%' 
           OR second_surname ILIKE '%' || $1 || '%' 
           OR personal_email ILIKE '%' || $1 || '%'
           OR institutional_email ILIKE '%' || $1 || '%'
    `

	rows, err := db.DB.Query(query, queryParam)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al buscar las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPeopleByFirstName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstName := vars["first_name"]

	query := `SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
              FROM "People" WHERE first_name = $1`
	rows, err := db.DB.Query(query, firstName)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPeopleByMiddleName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	middleName := vars["middle_name"]

	query := `SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
              FROM "People" WHERE middle_name = $1`
	rows, err := db.DB.Query(query, middleName)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPeopleByFirstSurname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstSurname := vars["first_surname"]

	query := `SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
              FROM "People" WHERE first_surname = $1`
	rows, err := db.DB.Query(query, firstSurname)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPeopleBySecondSurname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secondSurname := vars["second_surname"]

	query := `SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
              FROM "People" WHERE second_surname = $1`
	rows, err := db.DB.Query(query, secondSurname)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPeopleByPersonalEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personalEmail := vars["personal_email"]

	query := `SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
              FROM "People" WHERE personal_email = $1`
	rows, err := db.DB.Query(query, personalEmail)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPeopleByInstitutionalEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	institutionalEmail := vars["institutional_email"]

	query := `SELECT person_id, first_name, middle_name, first_surname, second_surname, personal_email, institutional_email 
              FROM "People" WHERE institutional_email = $1`
	rows, err := db.DB.Query(query, institutionalEmail)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener las personas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := []models.Person{}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.FirstName, &person.MiddleName, &person.FirstSurname, &person.SecondSurname, &person.PersonalEmail, &person.InstitutionalEmail); err != nil {
			log.Printf("Error al escanear las personas: %v\n", err)
			http.Error(w, "Error al escanear las personas", http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
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