package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
)

func GetJournals(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query(`SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                                  number, volume_number, edition_number, special_number, online_link, reserve_number 
                           FROM "Journals"`)
    if err != nil {
        log.Printf("Error al ejecutar la consulta: %v\n", err)
        http.Error(w, "Error al obtener las revistas", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    journals := []models.Journal{}
    for rows.Next() {
        var journal models.Journal
        if err := rows.Scan(&journal.JournalID, &journal.Status, &journal.Age, &journal.PublicationDate, &journal.StartMonthPeriod, 
                            &journal.EndMonthPeriod, &journal.Number, &journal.VolumeNumber, &journal.EditionNumber, &journal.SpecialNumber, 
                            &journal.OnlineLink, &journal.ReserveNumber); err != nil {
            log.Printf("Error al escanear las revistas: %v\n", err)
            http.Error(w, "Error al escanear las revistas", http.StatusInternalServerError)
            return
        }
        journals = append(journals, journal)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(journals)
}