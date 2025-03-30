CREATE TABLE IF NOT EXISTS applications
(
    id           UUID PRIMARY KEY,
    created_at   TIMESTAMP(3) NOT NULL,
    updated_at   TIMESTAMP(3) NOT NULL,
    deleted_at   TIMESTAMP(3),
    applicant_id UUID NOT NULL,
    scheme_id    UUID NOT NULL,
    CONSTRAINT fk_applications_applicant FOREIGN KEY (applicant_id) REFERENCES applicants (id),
    CONSTRAINT fk_applications_scheme FOREIGN KEY (scheme_id) REFERENCES schemes (id)
);

CREATE INDEX idx_applications_deleted_at ON applications (deleted_at);
CREATE INDEX fk_applications_applicant ON applications (applicant_id);
CREATE INDEX fk_applications_scheme ON applications (scheme_id);

-- Create triggers for the applications table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON applications
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();