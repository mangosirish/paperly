package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mangosirish/paperly/components"
	"github.com/mangosirish/paperly/db"
	"github.com/mangosirish/paperly/models"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT author_id, notes, person_id FROM "Authors"`)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los autores", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	authors := []models.Author{}
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.AuthorID, &author.Notes, &author.PersonID); err != nil {
			log.Printf("Error al escanear los autores: %v\n", err)
			http.Error(w, "Error al escanear los autores", http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func SearchAuthors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryParam := vars["query"]

	if queryParam == "" {
		http.Error(w, "Parámetro de búsqueda 'query' requerido", http.StatusBadRequest)
		return
	}

	query := `
        SELECT author_id, notes, person_id 
        FROM "Authors" 
        WHERE notes ILIKE '%' || $1 || '%'
    `

	rows, err := db.DB.Query(query, queryParam)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al buscar los autores", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	authors := []models.Author{}
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.AuthorID, &author.Notes, &author.PersonID); err != nil {
			log.Printf("Error al escanear los autores: %v\n", err)
			http.Error(w, "Error al escanear los autores", http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorsByNotes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notes := vars["notes"]

	query := `SELECT author_id, notes, person_id FROM "Authors" WHERE notes = $1`
	rows, err := db.DB.Query(query, notes)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los autores", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	authors := []models.Author{}
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.AuthorID, &author.Notes, &author.PersonID); err != nil {
			log.Printf("Error al escanear los autores: %v\n", err)
			http.Error(w, "Error al escanear los autores", http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorsByPersonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["person_id"]

	query := `SELECT author_id, notes, person_id FROM "Authors" WHERE person_id = $1`
	rows, err := db.DB.Query(query, personID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los autores", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	authors := []models.Author{}
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.AuthorID, &author.Notes, &author.PersonID); err != nil {
			log.Printf("Error al escanear los autores: %v\n", err)
			http.Error(w, "Error al escanear los autores", http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func FetchJoinedAuthorInfo(db *sql.DB) ([]map[string]interface{}, error) {
	query := `
        SELECT 
            CONCAT(pe.first_surname, ' ', pe.second_surname, ', ', pe.first_name) AS "Nombre",
            ac.name AS "Carrera",
            i.name AS "Institución",
            f.name AS "Facultad",
            pe.personal_email AS "Correo electrónico",
            STRING_AGG(art.title, ', ') AS "Lista de artículos",
            a.notes AS "Anotaciones",
            s.name AS "Área de concentración"
        FROM 
            "Authors" a
        JOIN 
            "People" pe ON a.person_id = pe.person_id
        JOIN 
            "TransitiveArticleAuthors" taa ON a.author_id = taa.author_id
        JOIN 
            "Articles" art ON taa.article_id = art.article_id
        JOIN 
            "TransitiveAuthorAcademicProfiles" taap ON a.author_id = taap.author_id
        JOIN 
            "AcademicProfiles" ap ON taap.academic_profile_id = ap.academic_profile_id
        JOIN 
            "AcademicCareers" ac ON ap.career_id = ac.career_id
        JOIN 
            "Institutions" i ON ap.institution_id = i.institution_id
        JOIN 
            "Faculties" f ON ap.faculty_id = f.faculty_id
        JOIN 
            "Specialties" s ON ap.specialty_id = s.specialty_id
        GROUP BY 
            "Nombre", "Carrera", "Institución", "Facultad", "Correo electrónico", "Anotaciones", "Área de concentración";
    `

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	authors := []map[string]interface{}{}
	for rows.Next() {
		var nombre, carrera, institucion, facultad, correo, listaArticulos, anotaciones, areaConcentracion string

		// Escanea los resultados en variables
		if err := rows.Scan(
			&nombre,
			&carrera,
			&institucion,
			&facultad,
			&correo,
			&listaArticulos,
			&anotaciones,
			&areaConcentracion,
		); err != nil {
			log.Printf("Error al escanear los datos: %v\n", err)
			return nil, err
		}

		// Construye el mapa para el autor
		authors = append(authors, map[string]interface{}{
			"Nombre":                nombre,
			"Carrera":               carrera,
			"Institución":           institucion,
			"Facultad":              facultad,
			"Correo electrónico":    correo,
			"Lista de artículos":    listaArticulos,
			"Anotaciones":           anotaciones,
			"Área de concentración": areaConcentracion,
		})
	}

	return authors, nil
}

func GetJoinedAuthorInfo(w http.ResponseWriter, r *http.Request) {
	// Llamar a la función de fetch
	authors, err := FetchJoinedAuthorInfo(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener la información de los autores", http.StatusInternalServerError)
		return
	}

	// Responder con JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// WEB
func RenderAuthorsTable(w http.ResponseWriter, r *http.Request) {
	// Llamar a FetchJoinedAuthorInfo para obtener datos
	authors, err := FetchJoinedAuthorInfo(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener los datos de los autores", http.StatusInternalServerError)
		return
	}

	// Renderizar la vista usando templ
	components.AuthorsTable(authors).Render(r.Context(), w)
}

// Crear un nuevo registro de Author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	// Decodificar el cuerpo de la solicitud en el modelo Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "Authors" (notes, person_id)
			  VALUES ($1, $2) RETURNING author_id`
	err = db.DB.QueryRow(query, author.Notes, author.PersonID).Scan(&author.AuthorID)
	if err != nil {
		http.Error(w, "Error al crear el autor", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// Actualizar un registro de Author
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	// Decodificar el cuerpo de la solicitud en el modelo Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el AuthorID esté presente
	if author.AuthorID == 0 {
		http.Error(w, "Se requiere un AuthorID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "Authors" 
			  SET notes = $1, person_id = $2
			  WHERE author_id = $3`
	result, err := db.DB.Exec(query, author.Notes, author.PersonID, author.AuthorID)
	if err != nil {
		http.Error(w, "Error al actualizar el autor", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el autor con el ID proporcionado", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// Eliminar un registro de Author
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	// Obtener el AuthorID de los parámetros de la ruta
	vars := mux.Vars(r)
	authorID := vars["author_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "Authors" WHERE author_id = $1`
	result, err := db.DB.Exec(query, authorID)
	if err != nil {
		http.Error(w, "Error al eliminar el autor", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el autor con el ID proporcionado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}