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