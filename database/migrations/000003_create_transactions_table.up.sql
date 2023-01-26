BEGIN;
CREATE TABLE transactions (
  id VARCHAR(32),
  "from" VARCHAR(32) NOT NULL,
  "to" VARCHAR(255) NOT NULL,
  status BOOLEAN NOT NULL,
  value DECIMAL(32,16) NOT NULL,
  created_at BIGINT DEFAULT 0 NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY("from") REFERENCES wallets (id),
  FOREIGN KEY("to") REFERENCES wallets (id)
);
COMMIT;