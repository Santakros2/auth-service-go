CREATE TABLE login_attempts (
    id CHAR(36) PRIMARY KEY,
    email VARCHAR(255),
    ip_address VARCHAR(45),
    success BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
