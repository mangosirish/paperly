package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
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

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error abriendo la base de datos: %v\n", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v\n", err)
    }
    fmt.Println("Conexión a la base de datos establecida correctamente")
}

func CloseDB() {
    DB.Close()
}