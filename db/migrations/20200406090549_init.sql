-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pgcrypto;
CREATE TABLE IF NOT EXISTS links
(
    id             	UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title          	VARCHAR(125) NOT NULL,
    description     VARCHAR(512) NOT NULL,
    fee             VARCHAR(20)  NOT NULL,
    payment_address	VARCHAR(34)  NULL,
    url 			VARCHAR(34)  NULL,
    created_at     	TIMESTAMP    NOT NULL,
    updated_at     	TIMESTAMP    NULL,
    deleted_at     	TIMESTAMP    NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
