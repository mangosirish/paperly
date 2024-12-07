package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/mangosirish/paperly/db"
	"github.com/mangosirish/paperly/models"
)

func SearchStudentSocialServices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryParam := vars["query"]

	if queryParam == "" {
		http.Error(w, "Parámetro de búsqueda 'query' requerido", http.StatusBadRequest)
		return
	}

	// Consulta SQL ajustada para buscar en el campo "status"
	query := `
        SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id
        FROM "StudentSocialServices"
        WHERE status ILIKE '%' || $1 || '%'
    `

	rows, err := db.DB.Query(query, queryParam)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al buscar servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crear una lista para almacenar los resultados
	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, 
			&service.Status, &service.Division, &service.Institution, &service.PersonID); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	// Enviar la respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServices(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                              is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                              is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                              is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
                           FROM "StudentSocialServices"`)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByStartDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startDate := vars["start_date"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE start_date = $1`
	rows, err := db.DB.Query(query, startDate)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByEndDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	endDate := vars["end_date"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE end_date = $1`
	rows, err := db.DB.Query(query, endDate)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByDocumentation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	documentation := vars["documentation"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE documentation = $1`
	rows, err := db.DB.Query(query, documentation)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE status = $1`
	rows, err := db.DB.Query(query, status)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByDivision(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	division := vars["division"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE division = $1`
	rows, err := db.DB.Query(query, division)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByInstitution(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	institution := vars["institution"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE institution = $1`
	rows, err := db.DB.Query(query, institution)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByPersonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["person_id"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE person_id = $1`
	rows, err := db.DB.Query(query, personID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsEnrollmentRequestSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isEnrollmentRequestSubmitted := vars["is_enrollment_request_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_enrollment_request_submitted = $1`
	rows, err := db.DB.Query(query, isEnrollmentRequestSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsPresentationLetterSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isPresentationLetterSubmitted := vars["is_presentation_letter_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_presentation_letter_submitted = $1`
	rows, err := db.DB.Query(query, isPresentationLetterSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsAcceptanceLetterSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isAcceptanceLetterSubmitted := vars["is_acceptance_letter_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_acceptance_letter_submitted = $1`
	rows, err := db.DB.Query(query, isAcceptanceLetterSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsAdvisorAppointmentSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isAdvisorAppointmentSubmitted := vars["is_advisor_appointment_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_advisor_appointment_submitted = $1`
	rows, err := db.DB.Query(query, isAdvisorAppointmentSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsCommitmentLetterSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isCommitmentLetterSubmitted := vars["is_commitment_letter_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_commitment_letter_submitted = $1`
	rows, err := db.DB.Query(query, isCommitmentLetterSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsIntermediateReportSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isIntermediateReportSubmitted := vars["is_intermediate_report_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_intermediate_report_submitted = $1`
	rows, err := db.DB.Query(query, isIntermediateReportSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsIntermediateReportValidationSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isIntermediateReportValidationSubmitted := vars["is_intermediate_report_validation_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_intermediate_report_validation_submitted = $1`
	rows, err := db.DB.Query(query, isIntermediateReportValidationSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsFinalReportSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isFinalReportSubmitted := vars["is_final_report_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_final_report_submitted = $1`
	rows, err := db.DB.Query(query, isFinalReportSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServicesByIsCompletionLetterSubmitted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isCompletionLetterSubmitted := vars["is_completion_letter_submitted"]

	query := `SELECT social_service_id, start_date, end_date, documentation, status, division, institution, person_id, 
                     is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted, 
                     is_advisor_appointment_submitted, is_commitment_letter_submitted, is_intermediate_report_submitted, 
                     is_intermediate_report_validation_submitted, is_final_report_submitted, is_completion_letter_submitted 
              FROM "StudentSocialServices" WHERE is_completion_letter_submitted = $1`
	rows, err := db.DB.Query(query, isCompletionLetterSubmitted)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los servicios sociales estudiantiles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	services := []models.StudentSocialService{}
	for rows.Next() {
		var service models.StudentSocialService
		if err := rows.Scan(&service.SocialServiceID, &service.StartDate, &service.EndDate, &service.Documentation, &service.Status,
			&service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, &service.IsPresentationLetterSubmitted,
			&service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, &service.IsCommitmentLetterSubmitted,
			&service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, &service.IsFinalReportSubmitted,
			&service.IsCompletionLetterSubmitted); err != nil {
			log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
			http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetStudentSocialServiceDetails(w http.ResponseWriter, r *http.Request) {
	query := `
        SELECT 
            CONCAT(pe.first_surname, ' ', pe.second_surname, ', ', pe.first_name) AS "Nombre",
            CONCAT(pe.first_surname, ' ', pe.second_surname) AS "Apellido(s)",
            pe.first_name AS "Nombre(s)",
            sss.social_service_id AS "Matrícula",
            sss.division AS "División",
            ac.name AS "Carrera",
            sss.start_date AS "Fecha de inicio",
            sss.end_date AS "Fecha de término",
            sss.status AS "Estado",
            sss.documentation AS "Notas",
            ARRAY[
                sss.is_enrollment_request_submitted,
                sss.is_presentation_letter_submitted,
                sss.is_acceptance_letter_submitted,
                sss.is_advisor_appointment_submitted,
                sss.is_commitment_letter_submitted,
                sss.is_intermediate_report_submitted,
                sss.is_intermediate_report_validation_submitted,
                sss.is_final_report_submitted,
                sss.is_completion_letter_submitted
            ] AS "Documentación"
        FROM 
            StudentSocialServices sss
        JOIN 
            People pe ON sss.person_id = pe.person_id
        JOIN 
            Authors a ON pe.person_id = a.person_id
        JOIN 
            TransitiveAuthorAcademicProfiles taap ON a.author_id = taap.author_id
        JOIN 
            AcademicProfiles ap ON taap.academic_profile_id = ap.academic_profile_id
        JOIN 
            AcademicCareers ac ON ap.career_id = ac.career_id`

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		http.Error(w, "Error al obtener los detalles del servicio social", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	studentSocialServices := []map[string]interface{}{}
	for rows.Next() {
		var nombre, apellidos, nombres, matricula, division, carrera, estado, notas string
		var fechaInicio, fechaTermino sql.NullTime
		var documentacion []bool

		if err := rows.Scan(&nombre, &apellidos, &nombres, &matricula, &division, &carrera, &fechaInicio, &fechaTermino, &estado, &notas, pq.Array(&documentacion)); err != nil {
			log.Printf("Error al escanear los detalles del servicio social: %v\n", err)
			http.Error(w, "Error al escanear los detalles del servicio social", http.StatusInternalServerError)
			return
		}

		studentSocialServices = append(studentSocialServices, map[string]interface{}{
			"Nombre":           nombre,
			"Apellido(s)":      apellidos,
			"Nombre(s)":        nombres,
			"Matrícula":        matricula,
			"División":         division,
			"Carrera":          carrera,
			"Fecha de inicio":  fechaInicio.Time,
			"Fecha de término": fechaTermino.Time,
			"Estado":           estado,
			"Notas":            notas,
			"Documentación":    documentacion,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentSocialServices)
}

// Crear un nuevo registro de StudentSocialService
func CreateStudentSocialService(w http.ResponseWriter, r *http.Request) {
	var service models.StudentSocialService

	// Decodificar el cuerpo de la solicitud en el modelo StudentSocialService
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para insertar el nuevo registro
	query := `INSERT INTO "StudentSocialServices" 
			  (start_date, end_date, documentation, status, division, institution, person_id, 
			   is_enrollment_request_submitted, is_presentation_letter_submitted, is_acceptance_letter_submitted,
			   is_advisor_appointment_submitted, is_commitment_letter_submitted, 
			   is_intermediate_report_submitted, is_intermediate_report_validation_submitted, 
			   is_final_report_submitted, is_completion_letter_submitted) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
			  RETURNING social_service_id`

	err = db.DB.QueryRow(query, service.StartDate, service.EndDate, service.Documentation, service.Status, service.Division, service.Institution, service.PersonID,
		service.IsEnrollmentRequestSubmitted, service.IsPresentationLetterSubmitted, service.IsAcceptanceLetterSubmitted, 
		service.IsAdvisorAppointmentSubmitted, service.IsCommitmentLetterSubmitted, service.IsIntermediateReportSubmitted, 
		service.IsIntermediateReportValidationSubmitted, service.IsFinalReportSubmitted, service.IsCompletionLetterSubmitted).Scan(&service.SocialServiceID)
	if err != nil {
		http.Error(w, "Error al crear el servicio social", http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo registro creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

// Actualizar un registro de StudentSocialService
func UpdateStudentSocialService(w http.ResponseWriter, r *http.Request) {
	var service models.StudentSocialService

	// Decodificar el cuerpo de la solicitud en el modelo StudentSocialService
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Verificar que el SocialServiceID esté presente
	if service.SocialServiceID == 0 {
		http.Error(w, "Se requiere un SocialServiceID para la actualización", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para actualizar el registro
	query := `UPDATE "StudentSocialServices" 
			  SET start_date = $1, end_date = $2, documentation = $3, status = $4, division = $5, institution = $6, 
				  person_id = $7, is_enrollment_request_submitted = $8, is_presentation_letter_submitted = $9, 
				  is_acceptance_letter_submitted = $10, is_advisor_appointment_submitted = $11, 
				  is_commitment_letter_submitted = $12, is_intermediate_report_submitted = $13, 
				  is_intermediate_report_validation_submitted = $14, is_final_report_submitted = $15, 
				  is_completion_letter_submitted = $16
			  WHERE social_service_id = $17`
	result, err := db.DB.Exec(query, service.StartDate, service.EndDate, service.Documentation, service.Status, service.Division, service.Institution, 
		service.PersonID, service.IsEnrollmentRequestSubmitted, service.IsPresentationLetterSubmitted, 
		service.IsAcceptanceLetterSubmitted, service.IsAdvisorAppointmentSubmitted, service.IsCommitmentLetterSubmitted, 
		service.IsIntermediateReportSubmitted, service.IsIntermediateReportValidationSubmitted, service.IsFinalReportSubmitted, 
		service.IsCompletionLetterSubmitted, service.SocialServiceID)
	if err != nil {
		http.Error(w, "Error al actualizar el servicio social", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el servicio social con el ID proporcionado", http.StatusNotFound)
		return
	}

	// Responder con el registro actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

// Eliminar un registro de StudentSocialService
func DeleteStudentSocialService(w http.ResponseWriter, r *http.Request) {
	// Obtener el SocialServiceID de los parámetros de la ruta
	vars := mux.Vars(r)
	socialServiceID := vars["social_service_id"]

	// Ejecutar la consulta para eliminar el registro
	query := `DELETE FROM "StudentSocialServices" WHERE social_service_id = $1`
	result, err := db.DB.Exec(query, socialServiceID)
	if err != nil {
		http.Error(w, "Error al eliminar el servicio social", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No se encontró el servicio social con el ID proporcionado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Responder con código 204 No Content
}