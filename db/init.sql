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

