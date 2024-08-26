package handlers

import (
    "encoding/json"
    "log"
    "net/http"
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
                            &service.Division, &service.Institution, &service.PersonID, &service.IsEnrollmentRequestSubmitted, 
                            &service.IsPresentationLetterSubmitted, &service.IsAcceptanceLetterSubmitted, &service.IsAdvisorAppointmentSubmitted, 
                            &service.IsCommitmentLetterSubmitted, &service.IsIntermediateReportSubmitted, &service.IsIntermediateReportValidationSubmitted, 
                            &service.IsFinalReportSubmitted, &service.IsCompletionLetterSubmitted); err != nil {
            log.Printf("Error al escanear los servicios sociales estudiantiles: %v\n", err)
            http.Error(w, "Error al escanear los servicios sociales estudiantiles", http.StatusInternalServerError)
            return
        }
        services = append(services, service)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(services)
}