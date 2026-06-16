CREATE TABLE IF NOT EXISTS occupations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_occupations_code ON occupations(code);
CREATE TRIGGER update_occupations_updated_at BEFORE UPDATE ON occupations FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();