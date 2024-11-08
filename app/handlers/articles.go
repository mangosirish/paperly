package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
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

func GetJoinedArticleInfo(w http.ResponseWriter, r *http.Request) {
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

    rows, err := db.DB.Query(query)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener la información", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    results := []map[string]interface{}{}
    for rows.Next() {
        var result map[string]interface{}
        err := rows.Scan(
            &result["Nombre"],
            &result["Autor"],
            &result["Tipo"],
            &result["Antigüedad"],
            &result["Fecha de recepción"],
            &result["Ejemplar"],
            &result["Estado"],
            &result["Autor Plano"],
            &result["Ejemplar Plano"],
            &result["Anotaciones"],
            &result["Numeración"],
            &result["Artículo original"],
        )
        if err != nil {
            log.Printf("Error al escanear los datos: %v\n", err)
            http.Error(w, "Error al procesar los datos", http.StatusInternalServerError)
            return
        }
        results = append(results, result)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}