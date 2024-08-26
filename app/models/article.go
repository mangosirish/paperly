package models

type Article struct {
    ArticleID      int    `json:"article_id"`
    Title          string `json:"title"`
    Type           string `json:"type"`
    ReceptionDate  string `json:"reception_date"`
    Status         string `json:"status"`
}