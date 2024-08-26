package models

type Journal struct {
    JournalID       int    `json:"journal_id"`
    Status          string `json:"status"`
    Age             int    `json:"age"`
    PublicationDate string `json:"publication_date"`
    StartMonthPeriod string `json:"start_month_period"`
    EndMonthPeriod  string `json:"end_month_period"`
    Number          int    `json:"number"`
    VolumeNumber    int    `json:"volume_number"`
    EditionNumber   int    `json:"edition_number"`
    SpecialNumber   int    `json:"special_number"`
    OnlineLink      string `json:"online_link"`
    ReserveNumber   string `json:"reserve_number"`
}