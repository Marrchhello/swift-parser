CREATE TABLE IF NOT EXISTS swift_codes (
    id SERIAL PRIMARY KEY,
    swift_code VARCHAR(11) UNIQUE NOT NULL,
    country_iso2 CHAR(2) NOT NULL,
    country_name VARCHAR(100) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    address TEXT,
    is_headquarter BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_swift_codes_country_iso2 ON swift_codes(country_iso2);
CREATE INDEX IF NOT EXISTS idx_swift_codes_swift_code ON swift_codes(swift_code);