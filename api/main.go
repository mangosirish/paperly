package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

type Article struct {
    ArticleID      int    `json:"article_id"`
    Title          string `json:"title"`
    Type           string `json:"type"`
    ReceptionDate  string `json:"reception_date"`
    Status         string `json:"status"`
}

var db *sql.DB

func initDB() {
    var err error

    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error cargando el archivo .env: %v\n", err)
    }

    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    sslmode := os.Getenv("DB_SSLMODE")

    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", user, password, dbname, host, sslmode)

    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error abriendo la base de datos: %v\n", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v\n", err)
    }
    fmt.Println("Conexión a la base de datos establecida correctamente")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT article_id, title, type, reception_date, status FROM Articles")
    if err != nil {
        http.Error(w, "Error al obtener los artículos", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    articles := []Article{}
    for rows.Next() {
        var article Article
        if err := rows.Scan(&article.ArticleID, &article.Title, &article.Type, &article.ReceptionDate, &article.Status); err != nil {
            http.Error(w, "Error al escanear los artículos", http.StatusInternalServerError)
            return
        }
        articles = append(articles, article)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("/articles", getArticles)
    fmt.Println("Servidor en ejecución en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}