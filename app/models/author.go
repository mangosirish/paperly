package models

type Author struct {
    AuthorID int    `json:"author_id"`
    Notes    string `json:"notes"`
    PersonID int    `json:"person_id"`
}