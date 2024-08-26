package models

type Person struct {
    PersonID           int    `json:"person_id"`
    FirstName          string `json:"first_name"`
    MiddleName         string `json:"middle_name"`
    FirstSurname       string `json:"first_surname"`
    SecondSurname      string `json:"second_surname"`
    PersonalEmail      string `json:"personal_email"`
    InstitutionalEmail string `json:"institutional_email"`
}