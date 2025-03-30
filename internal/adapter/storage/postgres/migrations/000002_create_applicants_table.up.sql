-- Create custom types for enums
CREATE TYPE employment_status AS ENUM ('employed', 'unemployed');
CREATE TYPE marital_status AS ENUM ('single', 'married', 'widowed', 'divorce');
CREATE TYPE sex AS ENUM ('male', 'female');

-- Create applicants table
CREATE TABLE IF NOT EXISTS applicants
(
    id                UUID PRIMARY KEY,
    created_at        TIMESTAMP(3) NOT NULL,
    updated_at        TIMESTAMP(3) NOT NULL,
    deleted_at        TIMESTAMP(3),
    name              TEXT NOT NULL,
    employment_status employment_status DEFAULT 'unemployed' NOT NULL,
    marital_status    marital_status    DEFAULT 'single' NOT NULL,
    sex               sex               DEFAULT 'male' NOT NULL,
    date_of_birth     DATE NOT NULL
);

CREATE INDEX idx_applicants_deleted_at ON applicants (deleted_at);

-- Create triggers for the applicants table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON applicants
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();