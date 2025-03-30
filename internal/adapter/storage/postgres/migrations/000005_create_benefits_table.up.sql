-- Create benefits table
CREATE TABLE IF NOT EXISTS benefits
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP(3) NOT NULL,
    updated_at TIMESTAMP(3) NOT NULL,
    deleted_at TIMESTAMP(3),
    scheme_id  UUID NOT NULL,
    name       TEXT NOT NULL,
    amount     DOUBLE PRECISION,
    CONSTRAINT fk_benefits_scheme FOREIGN KEY (scheme_id) REFERENCES schemes (id)
);

CREATE INDEX idx_benefits_deleted_at ON benefits (deleted_at);
CREATE INDEX fk_benefits_scheme ON benefits (scheme_id);

-- Create triggers for the benefits table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON benefits
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();