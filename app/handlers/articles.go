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

func GetArticles(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT article_id, title, type, reception_date, status FROM "Articles"`)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los artículos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ArticleID, &article.Title, &article.Type, &article.ReceptionDate, &article.Status); err != nil {
			log.Printf("Error al escanear los artículos: %v\n", err)
			http.Error(w, "Error al escanear los artículos", http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetArticlesByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	query := `SELECT article_id, title, type, reception_date, status FROM "Articles" WHERE title = $1`
	rows, err := db.DB.Query(query, title)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los artículos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ArticleID, &article.Title, &article.Type, &article.ReceptionDate, &article.Status); err != nil {
			log.Printf("Error al escanear los artículos: %v\n", err)
			http.Error(w, "Error al escanear los artículos", http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetArticlesByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleType := vars["type"]

	query := `SELECT article_id, title, type, reception_date, status FROM "Articles" WHERE type = $1`
	rows, err := db.DB.Query(query, articleType)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los artículos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ArticleID, &article.Title, &article.Type, &article.ReceptionDate, &article.Status); err != nil {
			log.Printf("Error al escanear los artículos: %v\n", err)
			http.Error(w, "Error al escanear los artículos", http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetArticlesByReceptionDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receptionDate := vars["reception_date"]

	query := `SELECT article_id, title, type, reception_date, status FROM "Articles" WHERE reception_date = $1`
	rows, err := db.DB.Query(query, receptionDate)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los artículos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ArticleID, &article.Title, &article.Type, &article.ReceptionDate, &article.Status); err != nil {
			log.Printf("Error al escanear los artículos: %v\n", err)
			http.Error(w, "Error al escanear los artículos", http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetArticlesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	query := `SELECT article_id, title, type, reception_date, status FROM "Articles" WHERE status = $1`
	rows, err := db.DB.Query(query, status)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los artículos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ArticleID, &article.Title, &article.Type, &article.ReceptionDate, &article.Status); err != nil {
			log.Printf("Error al escanear los artículos: %v\n", err)
			http.Error(w, "Error al escanear los artículos", http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func FetchJoinedArticleInfo(db *sql.DB) ([]map[string]interface{}, error) {
	query := `
	SELECT 
		art.title AS "Nombre",
		CONCAT(pe.first_name, ' ', pe.first_surname, ', ', pe.second_surname) AS "Autor",
		art.type AS "Tipo",
		EXTRACT(YEAR FROM AGE(CURRENT_DATE, art.reception_date)) AS "Antigüedad",
		art.reception_date AS "Fecha de recepción",
		CONCAT('Año ', j.age, ', vol. ', j.volume_number, ', núm. ', j.number, ', ', j.start_month_period, '-', j.end_month_period, ' ', j.publication_date) AS "Ejemplar",
		art.status AS "Estado",
		CONCAT(pe.first_name, ' ', pe.first_surname) AS "Autor Plano",
		j.online_link AS "Ejemplar Plano",
		a.notes AS "Anotaciones",
		CONCAT('Volumen ', j.volume_number, ', Número ', j.number) AS "Numeración",
		art.title AS "Artículo original"
	FROM 
		"Articles" art
	JOIN 
		"TransitiveArticleAuthors" taa ON art.article_id = taa.article_id
	JOIN 
		"Authors" a ON taa.author_id = a.author_id
	JOIN 
		"People" pe ON a.person_id = pe.person_id
	JOIN 
		"TransitiveArticleJournals" taj ON art.article_id = taj.article_id
	JOIN 
		"Journals" j ON taj.journal_id = j.journal_id;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	articles := []map[string]interface{}{}
	for rows.Next() {
		var nombre, autor, tipo, ejemplar, estado, autorPlano, ejemplarPlano, anotaciones, numeracion, articuloOriginal string
		var antiguedad int
		var fechaRecepcion sql.NullTime

		if err := rows.Scan(
			&nombre,
			&autor,
			&tipo,
			&antiguedad,
			&fechaRecepcion,
			&ejemplar,
			&estado,
			&autorPlano,
			&ejemplarPlano,
			&anotaciones,
			&numeracion,
			&articuloOriginal,
		); err != nil {
			log.Printf("Error al escanear los datos: %v\n", err)
			return nil, err
		}

		articles = append(articles, map[string]interface{}{
			"Nombre":             nombre,
			"Autor":              autor,
			"Tipo":               tipo,
			"Antigüedad":         antiguedad,
			"Fecha de recepción": fechaRecepcion.Time,
			"Ejemplar":           ejemplar,
			"Estado":             estado,
			"Autor Plano":        autorPlano,
			"Ejemplar Plano":     ejemplarPlano,
			"Anotaciones":        anotaciones,
			"Numeración":         numeracion,
			"Artículo original":  articuloOriginal,
		})
	}
	return articles, nil
}

func GetJoinedArticleInfo(w http.ResponseWriter, r *http.Request) {
	articles, err := FetchJoinedAuthorInfo(db.DB)

	if err != nil {
		http.Error(w, "Error al obtener la información de la base de datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func RenderArticlesTable(w http.ResponseWriter, r *http.Request) {
	articles, err := FetchJoinedArticleInfo(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener la información de la base de datos", http.StatusInternalServerError)
		return
	}

	components.ArticlesTable(articles).Render(r.Context(), w)
}

// Crear un nuevo registro de Article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	// Decodificar el cuerpo de la solicitud en el modelo Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "Articles" (title, type, reception_date, status)
			  VALUES ($1, $2, $3, $4) RETURNING article_id`
	err = db.DB.QueryRow(query, article.Title, article.Type, article.ReceptionDate, article.Status).Scan(&article.ArticleID)
	if err != nil {
		http.Error(w, "Error al crear el artículo", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// Actualizar un registro de Article
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	// Decodificar el cuerpo de la solicitud en el modelo Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el ArticleID esté presente
	if article.ArticleID == 0 {
		http.Error(w, "Se requiere un ArticleID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "Articles" 
			  SET title = $1, type = $2, reception_date = $3, status = $4
			  WHERE article_id = $5`
	result, err := db.DB.Exec(query, article.Title, article.Type, article.ReceptionDate, article.Status, article.ArticleID)
	if err != nil {
		http.Error(w, "Error al actualizar el artículo", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el artículo con el ID proporcionado", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// Eliminar un registro de Article
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// Obtener el ArticleID de los parámetros de la ruta
	vars := mux.Vars(r)
	articleID := vars["article_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "Articles" WHERE article_id = $1`
	result, err := db.DB.Exec(query, articleID)
	if err != nil {
		http.Error(w, "Error al eliminar el artículo", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el artículo con el ID proporcionado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}