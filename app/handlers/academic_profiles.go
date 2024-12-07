package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
	"github.com/mangosirish/paperly/components"
	"fmt"
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

func FetchAcademicProfileInfo(db *sql.DB) ([]map[string]interface{}, error) {
	query := `
        SELECT 
            ap.academic_profile_id AS "ID de Perfil Académico",
            i.name AS "Institución",
            f.name AS "Facultad",
            ac.name AS "Carrera",
            s.name AS "Área de Especialización"
        FROM 
            "AcademicProfiles" ap
        JOIN 
            "Institutions" i ON ap.institution_id = i.institution_id
        JOIN 
            "Faculties" f ON ap.faculty_id = f.faculty_id
        JOIN 
            "AcademicCareers" ac ON ap.career_id = ac.career_id
        JOIN 
            "Specialties" s ON ap.specialty_id = s.specialty_id
    `

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	profiles := []map[string]interface{}{}
	for rows.Next() {
		var academicProfileID, institution, faculty, career, specialty string

		// Escanea los resultados en variables
		if err := rows.Scan(
			&academicProfileID,
			&institution,
			&faculty,
			&career,
			&specialty,
		); err != nil {
			log.Printf("Error al escanear los datos: %v\n", err)
			return nil, err
		}

		// Construye el mapa para el perfil académico
		profiles = append(profiles, map[string]interface{}{
			"ID de Perfil Académico": academicProfileID,
			"Institución":            institution,
			"Facultad":               faculty,
			"Carrera":                career,
			"Área de Especialización": specialty,
		})
	}

	return profiles, nil
}

func GetJoinedAcademicProfileInfo(w http.ResponseWriter, r *http.Request) {
	// Llamar a la función de fetch
	profiles, err := FetchAcademicProfileInfo(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener la información de los perfiles académicos", http.StatusInternalServerError)
		return
	}

	// Responder con JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

// WEB
func RenderAcademicProfilesTable(w http.ResponseWriter, r *http.Request) {
	// Llamar a FetchAcademicProfileInfo para obtener datos
	profiles, err := FetchAcademicProfileInfo(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener los datos de los perfiles académicos", http.StatusInternalServerError)
		return
	}

	// Renderizar la vista usando templ (suponiendo que tengas una función similar en tus componentes)
	components.AcademicProfilesTable(profiles).Render(r.Context(), w)
}

// Crear un nuevo registro de AcademicProfile
func CreateAcademicProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.AcademicProfile

	// Decodificar el cuerpo de la solicitud en el modelo AcademicProfile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "AcademicProfiles" (institution_id, faculty_id, career_id, specialty_id)
			  VALUES ($1, $2, $3, $4) RETURNING academic_profile_id`
	err = db.DB.QueryRow(query, profile.InstitutionID, profile.FacultyID, profile.CareerID, profile.SpecialtyID).Scan(&profile.AcademicProfileID)
	if err != nil {
		http.Error(w, "Error al crear el registro de perfil académico", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"your_project_name/db"
	"your_project_name/models"
)

// CreateAcademicProfile maneja la creación de un nuevo perfil académico
func CreateComposedAcademicProfile(w http.ResponseWriter, r *http.Request) {
	// Decodificar los datos enviados en el cuerpo de la solicitud
	var profile models.AcademicProfile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}

	// Verificar que se haya enviado un ID de persona
	if profile.PersonID == 0 {
		http.Error(w, "El ID de persona es necesario", http.StatusBadRequest)
		return
	}

	// Insertar el perfil académico en la base de datos
	query := `
		INSERT INTO "AcademicProfiles" (institution_id, faculty_id, career_id, specialty_id)
		VALUES ($1, $2, $3, $4)
		RETURNING academic_profile_id`
	err = db.DB.QueryRow(query, profile.InstitutionID, profile.FacultyID, profile.CareerID, profile.SpecialtyID).Scan(&profile.AcademicProfileID)
	if err != nil {
		http.Error(w, "Error al crear el perfil académico", http.StatusInternalServerError)
		return
	}

	// Insertar la relación en la tabla transitiva "TransitiveAuthorAcademicProfiles"
	transitiveQuery := `
		INSERT INTO "TransitiveAuthorAcademicProfiles" (author_id, academic_profile_id)
		VALUES ($1, $2)`
	_, err = db.DB.Exec(transitiveQuery, profile.PersonID, profile.AcademicProfileID)
	if err != nil {
		http.Error(w, "Error al crear la relación transitiva", http.StatusInternalServerError)
		return
	}

	// Responder con el perfil académico recién creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// Actualizar un registro de AcademicProfile
func UpdateAcademicProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.AcademicProfile

	// Decodificar el cuerpo de la solicitud en el modelo AcademicProfile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el AcademicProfileID esté presente
	if profile.AcademicProfileID == 0 {
		http.Error(w, "Se requiere un AcademicProfileID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "AcademicProfiles" 
			  SET institution_id = $1, faculty_id = $2, career_id = $3, specialty_id = $4
			  WHERE academic_profile_id = $5`
	result, err := db.DB.Exec(query, profile.InstitutionID, profile.FacultyID, profile.CareerID, profile.SpecialtyID, profile.AcademicProfileID)
	if err != nil {
		http.Error(w, "Error al actualizar el registro de perfil académico", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el registro de perfil académico", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// Eliminar un registro de AcademicProfile
func DeleteAcademicProfile(w http.ResponseWriter, r *http.Request) {
	// Obtener el AcademicProfileID de los parámetros de la ruta
	vars := mux.Vars(r)
	academicProfileID := vars["academic_profile_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "AcademicProfiles" WHERE academic_profile_id = $1`
	result, err := db.DB.Exec(query, academicProfileID)
	if err != nil {
		http.Error(w, "Error al eliminar el registro de perfil académico", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el registro de perfil académico", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}

// CreateAcademicProfileFromPerson maneja la creación de un perfil académico basado en el ID de la persona
func CreateAcademicProfileFromPerson(w http.ResponseWriter, r *http.Request) {
	// Obtener el 'person_id' de la URL
	vars := mux.Vars(r)
	personID, err := strconv.Atoi(vars["person_id"])
	if err != nil {
		http.Error(w, "El ID de persona no es válido", http.StatusBadRequest)
		return
	}

	// Decodificar los datos enviados en el cuerpo de la solicitud
	var profile models.AcademicProfile
	err = json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}

	// Buscar el 'author_id' basado en 'person_id'
	var authorID int
	query := `SELECT "author_id" FROM "Authors" WHERE "person_id" = $1`
	err = db.DB.QueryRow(query, personID).Scan(&authorID)
	if err != nil {
		http.Error(w, "No se encontró un autor con esa persona", http.StatusNotFound)
		return
	}

	// Insertar el perfil académico en la base de datos
	query = `
		INSERT INTO "AcademicProfiles" (institution_id, faculty_id, career_id, specialty_id)
		VALUES ($1, $2, $3, $4)
		RETURNING academic_profile_id`
	err = db.DB.QueryRow(query, profile.InstitutionID, profile.FacultyID, profile.CareerID, profile.SpecialtyID).Scan(&profile.AcademicProfileID)
	if err != nil {
		http.Error(w, "Error al crear el perfil académico", http.StatusInternalServerError)
		return
	}

	// Insertar la relación en la tabla transitiva "TransitiveAuthorAcademicProfiles"
	transitiveQuery := `
		INSERT INTO "TransitiveAuthorAcademicProfiles" (author_id, academic_profile_id)
		VALUES ($1, $2)`
	_, err = db.DB.Exec(transitiveQuery, authorID, profile.AcademicProfileID)
	if err != nil {
		http.Error(w, "Error al crear la relación transitiva", http.StatusInternalServerError)
		return
	}

	// Responder con el perfil académico recién creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}