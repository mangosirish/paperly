package models

type StudentSocialService struct {
    SocialServiceID                    int    `json:"social_service_id"`
    StartDate                          string `json:"start_date"`
    EndDate                            string `json:"end_date"`
    Documentation                      string `json:"documentation"`
    Status                             string `json:"status"`
    Division                           string `json:"division"`
    Institution                        string `json:"institution"`
    PersonID                           int    `json:"person_id"`
    IsEnrollmentRequestSubmitted       bool   `json:"is_enrollment_request_submitted"`
    IsPresentationLetterSubmitted      bool   `json:"is_presentation_letter_submitted"`
    IsAcceptanceLetterSubmitted        bool   `json:"is_acceptance_letter_submitted"`
    IsAdvisorAppointmentSubmitted      bool   `json:"is_advisor_appointment_submitted"`
    IsCommitmentLetterSubmitted        bool   `json:"is_commitment_letter_submitted"`
    IsIntermediateReportSubmitted      bool   `json:"is_intermediate_report_submitted"`
    IsIntermediateReportValidationSubmitted bool   `json:"is_intermediate_report_validation_submitted"`
    IsFinalReportSubmitted             bool   `json:"is_final_report_submitted"`
    IsCompletionLetterSubmitted        bool   `json:"is_completion_letter_submitted"`
}