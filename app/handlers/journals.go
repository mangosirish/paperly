package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func GetJournalsWithArticles(w http.ResponseWriter, r *http.Request) {
	query := `
        SELECT 
            CONCAT('Año ', j.age, ', vol.', j.volume_number, ', núm. ', j.number, ', ', j.start_month_period, '-', j.end_month_period, ' ', EXTRACT(YEAR FROM j.publication_date)) AS "ID",
            j.status AS "Estado",
            j.age AS "Año (de antigüedad)",
            j.volume_number AS "Volumen",
            j.number AS "Número",
            COALESCE(NULLIF(j.special_number, 0), 'N/A') AS "Especial",
            CONCAT(j.start_month_period, '-', j.end_month_period) AS "Periodo",
            EXTRACT(YEAR FROM j.publication_date) AS "Año (de publicación)",
            CONCAT(j.edition_number, '.ª ed.') AS "Edición",
            CONCAT('Año ', j.age, ', volumen ', j.volume_number, ', número ', j.number, ', ', j.start_month_period, '-', j.end_month_period, ' ', EXTRACT(YEAR FROM j.publication_date)) AS "Ejemplar (largo)",
            CONCAT('Año ', j.age, ', vol.', j.volume_number, ', núm. ', j.number, ', ', j.start_month_period, '-', j.end_month_period) AS "Ejemplar (medio)",
            CONCAT('Año ', j.age, ', vol.', j.volume_number, ', núm. ', j.number) AS "Ejemplar (corto)",
            STRING_AGG(art.title, ', ') AS "Lista de artículos",
            j.publication_date AS "Fecha"
        FROM 
            Journals j
        JOIN 
            TransitiveArticleJournals taj ON j.journal_id = taj.journal_id
        JOIN 
            Articles art ON taj.article_id = art.article_id
        GROUP BY 
            "ID", "Estado", "Año (de antigüedad)", "Volumen", "Número", "Especial", "Periodo", "Año (de publicación)", "Edición", "Ejemplar (largo)", "Ejemplar (medio)", "Ejemplar (corto)", "Fecha"
    `

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los ejemplares", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	journals := []map[string]interface{}{}
	for rows.Next() {
		var id, estado, anoAntiguedad, volumen, numero, especial, periodo, anoPublicacion, edicion, ejemplarLargo, ejemplarMedio, ejemplarCorto, listaArticulos string
		var fecha sql.NullTime

		if err := rows.Scan(
			&id, &estado, &anoAntiguedad, &volumen, &numero, &especial, &periodo, &anoPublicacion,
			&edicion, &ejemplarLargo, &ejemplarMedio, &ejemplarCorto, &listaArticulos, &fecha,
		); err != nil {
			log.Printf("Error al escanear los ejemplares: %v\n", err)
			http.Error(w, "Error al escanear los ejemplares", http.StatusInternalServerError)
			return
		}

		journals = append(journals, map[string]interface{}{
			"ID":                   id,
			"Estado":               estado,
			"Año (de antigüedad)":  anoAntiguedad,
			"Volumen":              volumen,
			"Número":               numero,
			"Especial":             especial,
			"Periodo":              periodo,
			"Año (de publicación)": anoPublicacion,
			"Edición":              edicion,
			"Ejemplar (largo)":     ejemplarLargo,
			"Ejemplar (medio)":     ejemplarMedio,
			"Ejemplar (corto)":     ejemplarCorto,
			"Lista de artículos":   listaArticulos,
			"Fecha":                fecha.Time,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(journals)
}
