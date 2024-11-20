package models

type AcademicProfile struct {
	AcademicProfileID int `json:"academic_profile_id"`
	InstitutionID     int `json:"institution_id"`
	FacultyID         int `json:"faculty_id"`
	CareerID          int `json:"career_id"`
	SpecialtyID       int `json:"specialty_id"`
}