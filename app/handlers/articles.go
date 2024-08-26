package handlers

import (
    "encoding/json"
    "log"
    "net/http"
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