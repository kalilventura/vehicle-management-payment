-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vehicle_payment (
  id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  vehicle_id UUID NOT NULL,
  cpf        TEXT NOT NULL,
  amount     DECIMAL NOT NULL CHECK (amount > 0),
  status     VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'processing', 'approved', 'failed', 'refunded', 'chargeback')),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL
);
CREATE INDEX idx_payments_vehicle_id        ON vehicle_payment(vehicle_id);
CREATE INDEX idx_payments_customer_document ON vehicle_payment(cpf);
CREATE INDEX idx_payments_status            ON vehicle_payment(status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS vehicle_payment;
-- +goose StatementEnd
