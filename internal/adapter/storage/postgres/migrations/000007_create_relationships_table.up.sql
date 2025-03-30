-- Create relationship type
CREATE TYPE relationship_type AS ENUM ('spouse', 'child', 'parent', 'sibling');

-- Create relationships table
CREATE TABLE IF NOT EXISTS relationships
(
    id                UUID PRIMARY KEY,
    created_at        TIMESTAMP(3) NOT NULL,
    updated_at        TIMESTAMP(3) NOT NULL,
    deleted_at        TIMESTAMP(3),
    applicant_a_id    UUID NOT NULL,
    applicant_b_id    UUID NOT NULL,
    relationship_type relationship_type NOT NULL,
    CONSTRAINT fk_relationships_applicant_a FOREIGN KEY (applicant_a_id) REFERENCES applicants (id),
    CONSTRAINT fk_relationships_applicant_b FOREIGN KEY (applicant_b_id) REFERENCES applicants (id)
);

CREATE INDEX idx_relationships_deleted_at ON relationships (deleted_at);
CREATE INDEX fk_relationships_applicant_a ON relationships (applicant_a_id);
CREATE INDEX fk_relationships_applicant_b ON relationships (applicant_b_id);

-- Create triggers for the relationships table
CREATE TRIGGER set_timestamps
    BEFORE INSERT OR UPDATE
    ON relationships
    FOR EACH ROW
EXECUTE FUNCTION update_timestamps();