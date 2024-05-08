CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  bank_name VARCHAR(255) NOT NULL,
  owner_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  account_id INTEGER NOT NULL,
  description TEXT NOT NULL,
  amount NUMERIC NOT NULL,
  currency VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL,
  payment_method VARCHAR(255) NOT NULL,
  payment_at TIMESTAMPTZ NOT NULL,
  category VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_transactions_account_id FOREIGN KEY (account_id) REFERENCES accounts(id)
);