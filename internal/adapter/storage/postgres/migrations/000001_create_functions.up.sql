-- Create trigger function to handle timestamps
CREATE OR REPLACE FUNCTION update_timestamps()
    RETURNS TRIGGER AS
$$
BEGIN
    -- Update the "updated_at" column to the current timestamp
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
