CREATE TABLE "Articles" (
  "article_id" int PRIMARY KEY,
  "title" varchar,
  "type" varchar,
  "reception_date" date,
  "status" varchar
);

CREATE TABLE "People" (
  "person_id" int PRIMARY KEY,
  "first_name" varchar,
  "middle_name" varchar,
  "first_surname" varchar,
  "second_surname" varchar,
  "personal_email" varchar,
  "institutional_email" varchar
);

CREATE TABLE "Authors" (
  "author_id" int PRIMARY KEY,
  "notes" text,
  "person_id" integer UNIQUE
);

CREATE TABLE "AcademicCareers" (
  "career_id" int PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "Faculties" (
  "faculty_id" int PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "Institutions" (
  "institution_id" int PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "StudentSocialServices" (
  "social_service_id" int PRIMARY KEY,
  "start_date" date,
  "end_date" date,
  "documentation" text,
  "status" varchar,
  "division" varchar,
  "institution" varchar,
  "person_id" integer UNIQUE,
  "is_enrollment_request_submitted" boolean,
  "is_presentation_letter_submitted" boolean,
  "is_acceptance_letter_submitted" boolean,
  "is_advisor_appointment_submitted" boolean,
  "is_commitment_letter_submitted" boolean,
  "is_intermediate_report_submitted" boolean,
  "is_intermediate_report_validation_submitted" boolean,
  "is_final_report_submitted" boolean,
  "is_completion_letter_submitted" boolean
);

CREATE TABLE "Journals" (
  "journal_id" int PRIMARY KEY,
  "status" varchar,
  "age" int,
  "publication_date" date,
  "start_month_period" varchar,
  "end_month_period" varchar,
  "number" int,
  "volume_number" int,
  "edition_number" int,
  "special_number" int,
  "online_link" text,
  "reserve_number" varchar
);

CREATE TABLE "Specialties" (
  "specialty_id" int PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "TransitiveArticleAuthors" (
  "article_id" int,
  "author_id" int
);

CREATE TABLE "TransitiveArticleJournals" (
  "article_id" int,
  "journal_id" int
);

CREATE TABLE "TransitiveAuthorAcademicProfiles" (
  "author_id" int,
  "academic_profile_id" int
);

CREATE TABLE "AcademicProfiles" (
  "academic_profile_id" int PRIMARY KEY,
  "institution_id" int,
  "faculty_id" int,
  "career_id" int,
  "specialty_id" int
);

CREATE UNIQUE INDEX ON "TransitiveArticleAuthors" ("article_id", "author_id");

CREATE UNIQUE INDEX ON "TransitiveArticleJournals" ("article_id", "journal_id");

CREATE UNIQUE INDEX ON "TransitiveAuthorAcademicProfiles" ("author_id", "academic_profile_id");

CREATE UNIQUE INDEX ON "AcademicProfiles" ("institution_id", "faculty_id", "career_id", "specialty_id");

ALTER TABLE "Authors" ADD FOREIGN KEY ("person_id") REFERENCES "People" ("person_id");

ALTER TABLE "StudentSocialServices" ADD FOREIGN KEY ("person_id") REFERENCES "People" ("person_id");

ALTER TABLE "TransitiveArticleAuthors" ADD FOREIGN KEY ("article_id") REFERENCES "Articles" ("article_id");

ALTER TABLE "TransitiveArticleAuthors" ADD FOREIGN KEY ("author_id") REFERENCES "Authors" ("author_id");

ALTER TABLE "TransitiveArticleJournals" ADD FOREIGN KEY ("article_id") REFERENCES "Articles" ("article_id");

ALTER TABLE "TransitiveArticleJournals" ADD FOREIGN KEY ("journal_id") REFERENCES "Journals" ("journal_id");

ALTER TABLE "TransitiveAuthorAcademicProfiles" ADD FOREIGN KEY ("author_id") REFERENCES "Authors" ("author_id");

ALTER TABLE "TransitiveAuthorAcademicProfiles" ADD FOREIGN KEY ("academic_profile_id") REFERENCES "AcademicProfiles" ("academic_profile_id");

ALTER TABLE "AcademicProfiles" ADD FOREIGN KEY ("institution_id") REFERENCES "Institutions" ("institution_id");

ALTER TABLE "AcademicProfiles" ADD FOREIGN KEY ("faculty_id") REFERENCES "Faculties" ("faculty_id");

ALTER TABLE "AcademicProfiles" ADD FOREIGN KEY ("career_id") REFERENCES "AcademicCareers" ("career_id");

ALTER TABLE "AcademicProfiles" ADD FOREIGN KEY ("specialty_id") REFERENCES "Specialties" ("specialty_id");

-- Insertar datos en la tabla People
INSERT INTO "People" ("person_id", "first_name", "middle_name", "first_surname", "second_surname", "personal_email", "institutional_email")
VALUES 
(1, 'John', 'A.', 'Doe', 'Smith', 'john.doe@gmail.com', 'j.doe@university.edu'),
(2, 'Jane', 'B.', 'Doe', 'Johnson', 'jane.doe@yahoo.com', 'j.doe@university.edu'),
(3, 'Alice', 'C.', 'Johnson', 'Brown', 'alice.johnson@outlook.com', 'a.johnson@university.edu');

-- Insertar datos en la tabla Authors
INSERT INTO "Authors" ("author_id", "notes", "person_id")
VALUES 
(1, 'Expert in Artificial Intelligence', 1),
(2, 'Researcher in Quantum Computing', 2),
(3, 'Specialist in Machine Learning', 3);

-- Insertar datos en la tabla Articles
INSERT INTO "Articles" ("article_id", "title", "type", "reception_date", "status")
VALUES 
(1, 'Artificial Intelligence in Healthcare', 'Research', '2023-01-15', 'Published'),
(2, 'Advancements in Quantum Computing', 'Review', '2023-05-22', 'In Review'),
(3, 'Machine Learning Techniques', 'Research', '2023-07-10', 'Accepted');

-- Insertar datos en la tabla AcademicCareers
INSERT INTO "AcademicCareers" ("career_id", "name")
VALUES 
(1, 'Computer Science'),
(2, 'Electrical Engineering'),
(3, 'Mathematics');

-- Insertar datos en la tabla Faculties
INSERT INTO "Faculties" ("faculty_id", "name")
VALUES 
(1, 'Engineering'),
(2, 'Sciences'),
(3, 'Humanities');

-- Insertar datos en la tabla Institutions
INSERT INTO "Institutions" ("institution_id", "name")
VALUES 
(1, 'MIT'),
(2, 'Stanford University'),
(3, 'Harvard University');

-- Insertar datos en la tabla StudentSocialServices
INSERT INTO "StudentSocialServices" ("social_service_id", "start_date", "end_date", "documentation", "status", "division", "institution", "person_id", "is_enrollment_request_submitted", "is_presentation_letter_submitted", "is_acceptance_letter_submitted", "is_advisor_appointment_submitted", "is_commitment_letter_submitted", "is_intermediate_report_submitted", "is_intermediate_report_validation_submitted", "is_final_report_submitted", "is_completion_letter_submitted")
VALUES 
(1, '2023-02-01', '2023-06-30', 'All documents submitted.', 'Completed', 'Engineering', 'MIT', 1, true, true, true, true, true, true, true, true, true),
(2, '2023-03-01', '2023-07-31', 'Awaiting final report.', 'In Progress', 'Sciences', 'Stanford University', 2, true, true, true, true, true, true, false, false, false),
(3, '2023-01-15', '2023-05-31', 'Intermediate report pending.', 'In Progress', 'Humanities', 'Harvard University', 3, true, true, true, true, true, false, false, false, false);

-- Insertar datos en la tabla Journals
INSERT INTO "Journals" ("journal_id", "status", "age", "publication_date", "start_month_period", "end_month_period", "number", "volume_number", "edition_number", "special_number", "online_link", "reserve_number")
VALUES 
(1, 'Active', 5, '2023-08-01', 'July', 'August', 10, 3, 1, 0, 'http://journal1.com', 'J-001'),
(2, 'Inactive', 10, '2022-05-01', 'April', 'May', 25, 7, 2, 1, 'http://journal2.com', 'J-002'),
(3, 'Active', 7, '2023-09-01', 'August', 'September', 12, 4, 1, 0, 'http://journal3.com', 'J-003');

-- Insertar datos en la tabla Specialties
INSERT INTO "Specialties" ("specialty_id", "name")
VALUES 
(1, 'Artificial Intelligence'),
(2, 'Quantum Computing'),
(3, 'Machine Learning');

-- Insertar datos en la tabla TransitiveArticleAuthors
INSERT INTO "TransitiveArticleAuthors" ("article_id", "author_id")
VALUES 
(1, 1),
(2, 2),
(3, 3);

-- Insertar datos en la tabla TransitiveArticleJournals
INSERT INTO "TransitiveArticleJournals" ("article_id", "journal_id")
VALUES 
(1, 1),
(2, 2),
(3, 3);

-- Insertar datos en la tabla AcademicProfiles
INSERT INTO "AcademicProfiles" ("academic_profile_id", "institution_id", "faculty_id", "career_id", "specialty_id")
VALUES 
(1, 1, 1, 1, 1),
(2, 2, 2, 2, 2),
(3, 3, 3, 3, 3);

-- Insertar datos en la tabla TransitiveAuthorAcademicProfiles
INSERT INTO "TransitiveAuthorAcademicProfiles" ("author_id", "academic_profile_id")
VALUES 
(1, 1),
(2, 2),
(3, 3);

