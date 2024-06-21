-- +goose Up
-- +goose StatementBegin
CREATE TABLE loft (
    id SERIAL PRIMARY KEY,
    loft_id VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    website VARCHAR(255),
    phone VARCHAR(255),
    whatsapp VARCHAR(255),
    telegram VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE loft;
-- +goose StatementEnd
