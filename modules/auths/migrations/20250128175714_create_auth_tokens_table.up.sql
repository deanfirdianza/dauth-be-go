CREATE TABLE auths.auth_tokens (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE
);
