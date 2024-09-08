package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/mangosirish/paperly/db"
    "github.com/mangosirish/paperly/models"
)

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