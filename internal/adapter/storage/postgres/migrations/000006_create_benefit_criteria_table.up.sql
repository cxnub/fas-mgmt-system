-- Create benefit_criteria table
CREATE TABLE IF NOT EXISTS benefit_criteria
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP(3) NOT NULL,
    updated_at TIMESTAMP(3) NOT NULL,
    deleted_at TIMESTAMP(3),
    name       TEXT NOT NULL,
    value      TEXT,
    benefit_id UUID NOT NULL,
    CONSTRAINT fk_benefits_criteria FOREIGN KEY (benefit_id) REFERENCES benefits (id)
);

CREATE INDEX idx_benefit_criteria_deleted_at ON benefit_criteria (deleted_at);
CREATE INDEX fk_benefits_criteria ON benefit_criteria (benefit_id);

-- Create triggers for the benefit_criteria table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON benefit_criteria
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();