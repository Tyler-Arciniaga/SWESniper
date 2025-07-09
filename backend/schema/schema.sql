--@block
CREATE TABLE IF NOT EXISTS users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT,
    email TEXT,
    created_at TIMESTAMPTZ
);
--@block
CREATE TABLE IF NOT EXISTS urls (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id UUID REFERENCES auth.users (id) ON DELETE CASCADE,
    url TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL UNIQUE,
    check_interval INT NOT NULL,
    last_checked_at TIMESTAMPTZ,
    last_known_hash TEXT,
    last_known_content JSONB,
    created_at TIMESTAMPTZ
);
--@block
CREATE TABLE IF NOT EXISTS changeLogs (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    url_id INT REFERENCES urls (id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    timestamp TIMESTAMPTZ,
    added JSONB,
    diff_summary TEXT
);