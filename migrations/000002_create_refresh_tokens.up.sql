CREATE TABLE refresh_tokens (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,

    token_hash VARCHAR(255) NOT NULL,

    device_name VARCHAR(100),
    user_agent TEXT,
    ip_address VARCHAR(45),

    expires_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,
    revoked_reason VARCHAR(255),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_refresh_user
        FOREIGN KEY (user_id) REFERENCES auth_users(id)
        ON DELETE CASCADE
);
