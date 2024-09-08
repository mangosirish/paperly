package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
    "github.com/gorilla/mux"
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

func GetJournalsByStatus(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    status := vars["status"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE status = $1`
    rows, err := db.DB.Query(query, status)
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

func GetJournalsByAge(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    age := vars["age"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE age = $1`
    rows, err := db.DB.Query(query, age)
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

func GetJournalsByPublicationDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    publicationDate := vars["publication_date"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE publication_date = $1`
    rows, err := db.DB.Query(query, publicationDate)
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

func GetJournalsByStartMonthPeriod(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    startMonthPeriod := vars["start_month_period"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE start_month_period = $1`
    rows, err := db.DB.Query(query, startMonthPeriod)
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

func GetJournalsByEndMonthPeriod(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    endMonthPeriod := vars["end_month_period"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE end_month_period = $1`
    rows, err := db.DB.Query(query, endMonthPeriod)
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

func GetJournalsByNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    number := vars["number"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE number = $1`
    rows, err := db.DB.Query(query, number)
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

func GetJournalsByVolumeNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    volumeNumber := vars["volume_number"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE volume_number = $1`
    rows, err := db.DB.Query(query, volumeNumber)
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

func GetJournalsByEditionNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    editionNumber := vars["edition_number"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE edition_number = $1`
    rows, err := db.DB.Query(query, editionNumber)
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

func GetJournalsBySpecialNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    specialNumber := vars["special_number"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE special_number = $1`
    rows, err := db.DB.Query(query, specialNumber)
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

func GetJournalsByOnlineLink(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    onlineLink := vars["online_link"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE online_link = $1`
    rows, err := db.DB.Query(query, onlineLink)
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

func GetJournalsByReserveNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    reserveNumber := vars["reserve_number"]

    query := `SELECT journal_id, status, age, publication_date, start_month_period, end_month_period, 
                      number, volume_number, edition_number, special_number, online_link, reserve_number 
               FROM "Journals" WHERE reserve_number = $1`
    rows, err := db.DB.Query(query, reserveNumber)
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