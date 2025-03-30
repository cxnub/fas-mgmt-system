-- Create scheme_criteria table
CREATE TABLE IF NOT EXISTS scheme_criteria
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP(3) NOT NULL,
    updated_at TIMESTAMP(3) NOT NULL,
    deleted_at TIMESTAMP(3),
    name       TEXT NOT NULL,
    value      TEXT,
    scheme_id  UUID NOT NULL,
    CONSTRAINT fk_schemes_criteria FOREIGN KEY (scheme_id) REFERENCES schemes (id)
);

CREATE INDEX idx_scheme_criteria_deleted_at ON scheme_criteria (deleted_at);
CREATE INDEX fk_schemes_criteria ON scheme_criteria (scheme_id);

-- Create triggers for the scheme_criteria table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON scheme_criteria
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();