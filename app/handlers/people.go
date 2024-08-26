package handlers

import (
    "encoding/json"
    "log"
    "net/http"
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