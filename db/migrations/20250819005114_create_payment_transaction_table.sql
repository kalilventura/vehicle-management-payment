-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_transaction (
  id                     UUID DEFAULT uuidv7() PRIMARY KEY,
  payment_id             UUID NOT NULL REFERENCES vehicle_payment(id),
  gateway_transaction_id VARCHAR(255) NOT NULL,
  status                 VARCHAR(20) NOT NULL CHECK (status IN ('processing', 'success', 'failure')),
  response_code          VARCHAR(50),
  response_message       TEXT,
  raw_response           JSONB,
  created_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_transactions_gateway_id ON payment_transaction(gateway_transaction_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS payment_transaction;
-- +goose StatementEnd
