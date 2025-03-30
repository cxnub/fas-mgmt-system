-- Create schemes table
CREATE TABLE IF NOT EXISTS schemes
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP(3) NOT NULL,
    updated_at TIMESTAMP(3) NOT NULL,
    deleted_at TIMESTAMP(3),
    name       TEXT NOT NULL
);

CREATE INDEX idx_schemes_deleted_at ON schemes (deleted_at);

-- Create triggers for the schemes table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON schemes
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();