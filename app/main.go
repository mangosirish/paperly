package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mangosirish/paperly/components"
	"github.com/mangosirish/paperly/db"
	"github.com/mangosirish/paperly/handlers"
	"github.com/mangosirish/paperly/templates"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	router := mux.NewRouter()

	// Frontend
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.IndexView().Render(r.Context(), w)
	}).Methods("GET")

	router.HandleFunc("/web/{table}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		table := vars["table"]
		templates.TableView(table).Render(r.Context(), w)
	}).Methods("GET")

	router.HandleFunc("/web/table/{table}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		table := vars["table"]
		components.Table(table).Render(r.Context(), w)
	}).Methods("GET")

	router.HandleFunc("/web/form/{form}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		form := vars["form"]
		components.CreateItemForm(form).Render(r.Context(), w)
	}).Methods("GET")

	router.HandleFunc("/web/container/{table}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		table := vars["table"]
		components.Container(table).Render(r.Context(), w)

	}).Methods("GET")

	// API
	router.HandleFunc("/articles", handlers.GetArticles).Methods("GET")
	router.HandleFunc("/articles/title/{title}", handlers.GetArticlesByTitle).Methods("GET")
	router.HandleFunc("/articles/type/{type}", handlers.GetArticlesByType).Methods("GET")
	router.HandleFunc("/articles/reception-date/{reception_date}", handlers.GetArticlesByReceptionDate).Methods("GET")
	router.HandleFunc("/articles/status/{status}", handlers.GetArticlesByStatus).Methods("GET")
	router.HandleFunc("/articles/details", handlers.GetJoinedArticleInfo).Methods("GET")
	router.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{article_id:[0-9]+}", handlers.UpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{article_id:[0-9]+}", handlers.DeleteArticle).Methods("DELETE")

	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/search/{query}", handlers.SearchPeople).Methods("GET")
	router.HandleFunc("/people/first_name/{first_name}", handlers.GetPeopleByFirstName).Methods("GET")
	router.HandleFunc("/people/middle_name/{middle_name}", handlers.GetPeopleByMiddleName).Methods("GET")
	router.HandleFunc("/people/first_surname/{first_surname}", handlers.GetPeopleByFirstSurname).Methods("GET")
	router.HandleFunc("/people/second_surname/{second_surname}", handlers.GetPeopleBySecondSurname).Methods("GET")
	router.HandleFunc("/people/personal_email/{personal_email}", handlers.GetPeopleByPersonalEmail).Methods("GET")
	router.HandleFunc("/people/institutional_email/{institutional_email}", handlers.GetPeopleByInstitutionalEmail).Methods("GET")
	router.HandleFunc("/people", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{person_id:[0-9]+}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{person_id:[0-9]+}", handlers.DeletePerson).Methods("DELETE")

	router.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	router.HandleFunc("/authors/search/{query}", handlers.SearchAuthors).Methods("GET")
	router.HandleFunc("/authors/notes/{notes}", handlers.GetAuthorsByNotes).Methods("GET")
	router.HandleFunc("/authors/person-id/{person_id}", handlers.GetAuthorsByPersonID).Methods("GET")
	router.HandleFunc("/authors/details", handlers.GetJoinedAuthorInfo).Methods("GET")
	router.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{author_id:[0-9]+}", handlers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{author_id:[0-9]+}", handlers.DeleteAuthor).Methods("DELETE")

	router.HandleFunc("/academic-careers", handlers.GetAcademicCareers).Methods("GET")
	router.HandleFunc("/academic-careers/name/{name}", handlers.GetAcademicCareersByName).Methods("GET")
	router.HandleFunc("/academic-careers", handlers.CreateAcademicCareer).Methods("POST")
	router.HandleFunc("/academic-careers/{career_id:[0-9]+}", handlers.UpdateAcademicCareer).Methods("PUT")
	router.HandleFunc("/academic-careers/{career_id:[0-9]+}", handlers.DeleteAcademicCareer).Methods("DELETE")

	router.HandleFunc("/academic-profiles/search/{query}", handlers.SearchAcademicProfiles).Methods("GET")
	router.HandleFunc("/academic-profiles", handlers.CreateAcademicProfile).Methods("POST")
	router.HandleFunc("/academic_profiles/composed", handlers.CreateComposedAcademicProfile).Methods("POST")
	router.HandleFunc("/academic_profiles/person/{person_id:[0-9]+}", handlers.CreateAcademicProfileFromPerson).Methods("POST")
	router.HandleFunc("/academic-profiles/{academic_profile_id:[0-9]+}", handlers.UpdateAcademicProfile).Methods("PUT")
	router.HandleFunc("/academic-profiles/{academic_profile_id:[0-9]+}", handlers.DeleteAcademicProfile).Methods("DELETE")

	router.HandleFunc("/faculties", handlers.GetFaculties).Methods("GET")
	router.HandleFunc("/faculties/name/{name}", handlers.GetFacultiesByName).Methods("GET")
	router.HandleFunc("/faculties", handlers.CreateFaculty).Methods("POST")
	router.HandleFunc("/faculties/{faculty_id:[0-9]+}", handlers.UpdateFaculty).Methods("PUT")
	router.HandleFunc("/faculties/{faculty_id:[0-9]+}", handlers.DeleteFaculty).Methods("DELETE")

	router.HandleFunc("/institutions", handlers.GetInstitutions).Methods("GET")
	router.HandleFunc("/institutions/name/{name}", handlers.GetInstitutionsByName).Methods("GET")
	router.HandleFunc("/institutions", handlers.CreateInstitution).Methods("POST")
	router.HandleFunc("/institutions/{institution_id:[0-9]+}", handlers.UpdateInstitution).Methods("PUT")
	router.HandleFunc("/institutions/{institution_id:[0-9]+}", handlers.DeleteInstitution).Methods("DELETE")

	router.HandleFunc("/student-social-services", handlers.GetStudentSocialServices).Methods("GET")
	router.HandleFunc("/student-social-services/start_date/{start_date}", handlers.GetStudentSocialServicesByStartDate).Methods("GET")
	router.HandleFunc("/student-social-services/end_date/{end_date}", handlers.GetStudentSocialServicesByEndDate).Methods("GET")
	router.HandleFunc("/student-social-services/documentation/{documentation}", handlers.GetStudentSocialServicesByDocumentation).Methods("GET")
	router.HandleFunc("/student-social-services/status/{status}", handlers.GetStudentSocialServicesByStatus).Methods("GET")
	router.HandleFunc("/student-social-services/division/{division}", handlers.GetStudentSocialServicesByDivision).Methods("GET")
	router.HandleFunc("/student-social-services/institution/{institution}", handlers.GetStudentSocialServicesByInstitution).Methods("GET")
	router.HandleFunc("/student-social-services/person_id/{person_id:[0-9]+}", handlers.GetStudentSocialServicesByPersonID).Methods("GET")
	router.HandleFunc("/student-social-services/is_enrollment_request_submitted/{is_enrollment_request_submitted}", handlers.GetStudentSocialServicesByIsEnrollmentRequestSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_presentation_letter_submitted/{is_presentation_letter_submitted}", handlers.GetStudentSocialServicesByIsPresentationLetterSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_acceptance_letter_submitted/{is_acceptance_letter_submitted}", handlers.GetStudentSocialServicesByIsAcceptanceLetterSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_advisor_appointment_submitted/{is_advisor_appointment_submitted}", handlers.GetStudentSocialServicesByIsAdvisorAppointmentSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_commitment_letter_submitted/{is_commitment_letter_submitted}", handlers.GetStudentSocialServicesByIsCommitmentLetterSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_intermediate_report_submitted/{is_intermediate_report_submitted}", handlers.GetStudentSocialServicesByIsIntermediateReportSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_intermediate_report_validation_submitted/{is_intermediate_report_validation_submitted}", handlers.GetStudentSocialServicesByIsIntermediateReportValidationSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_final_report_submitted/{is_final_report_submitted}", handlers.GetStudentSocialServicesByIsFinalReportSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/is_completion_letter_submitted/{is_completion_letter_submitted}", handlers.GetStudentSocialServicesByIsCompletionLetterSubmitted).Methods("GET")
	router.HandleFunc("/student-social-services/details", handlers.GetStudentSocialServiceDetails).Methods("GET")
	router.HandleFunc("/student-social-services/search/{query}", handlers.SearchStudentSocialServices).Methods("GET")
	router.HandleFunc("/student_social_services", handlers.CreateStudentSocialService).Methods("POST")
	router.HandleFunc("/student_social_services/{social_service_id:[0-9]+}", handlers.UpdateStudentSocialService).Methods("PUT")
	router.HandleFunc("/student_social_services/{social_service_id:[0-9]+}", handlers.DeleteStudentSocialService).Methods("DELETE")

	router.HandleFunc("/journals", handlers.GetJournals).Methods("GET")
	router.HandleFunc("/journals/search/{query}", handlers.SearchJournals).Methods("GET")
	router.HandleFunc("/journals/status/{status}", handlers.GetJournalsByStatus).Methods("GET")
	router.HandleFunc("/journals/age/{age}", handlers.GetJournalsByAge).Methods("GET")
	router.HandleFunc("/journals/publication_date/{publication_date}", handlers.GetJournalsByPublicationDate).Methods("GET")
	router.HandleFunc("/journals/start_month_period/{start_month_period}", handlers.GetJournalsByStartMonthPeriod).Methods("GET")
	router.HandleFunc("/journals/end_month_period/{end_month_period}", handlers.GetJournalsByEndMonthPeriod).Methods("GET")
	router.HandleFunc("/journals/number/{number}", handlers.GetJournalsByNumber).Methods("GET")
	router.HandleFunc("/journals/volume_number/{volume_number}", handlers.GetJournalsByVolumeNumber).Methods("GET")
	router.HandleFunc("/journals/edition_number/{edition_number}", handlers.GetJournalsByEditionNumber).Methods("GET")
	router.HandleFunc("/journals/special_number/{special_number}", handlers.GetJournalsBySpecialNumber).Methods("GET")
	router.HandleFunc("/journals/online_link/{online_link}", handlers.GetJournalsByOnlineLink).Methods("GET")
	router.HandleFunc("/journals/reserve_number/{reserve_number}", handlers.GetJournalsByReserveNumber).Methods("GET")
	router.HandleFunc("/journals/details", handlers.GetJournalsWithArticles).Methods("GET")
	router.HandleFunc("/journals", handlers.CreateJournal).Methods("POST")
	router.HandleFunc("/journals/{journal_id:[0-9]+}", handlers.UpdateJournal).Methods("PUT")
	router.HandleFunc("/journals/{journal_id:[0-9]+}", handlers.DeleteJournal).Methods("DELETE")

	router.HandleFunc("/specialties", handlers.GetSpecialties).Methods("GET")
	router.HandleFunc("/specialties/name/{name}", handlers.GetSpecialtiesByName).Methods("GET")
	router.HandleFunc("/people", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{person_id:[0-9]+}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{person_id:[0-9]+}", handlers.DeletePerson).Methods("DELETE")

	router.HandleFunc("/create_author", handleForm("Author")).Methods("POST")
	router.HandleFunc("/create_journal", handleForm("Journal")).Methods("POST")
	router.HandleFunc("/create_social_service", handleForm("Social Service")).Methods("POST")
	router.HandleFunc("/create_article", handleForm("Article")).Methods("POST")

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleForm(formType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parsear el formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al analizar el formulario", http.StatusBadRequest)
			return
		}

		// Mostrar cada campo del formulario y sus valores
		fmt.Printf("Formulario %s:\n", formType)
		for key, values := range r.Form { // r.PostForm para datos solo de POST
			fmt.Printf("  Campo: %s, Valores: %v\n", key, values)
		}

		// Opcional: responder al cliente
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Formulario %s recibido con éxito.", formType)
	}
}