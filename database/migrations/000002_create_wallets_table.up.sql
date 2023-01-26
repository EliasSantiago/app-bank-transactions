BEGIN;
CREATE TABLE wallets (
  id VARCHAR(32),
  user_id VARCHAR(32),
  balance DECIMAL(32,16),
  created_at BIGINT DEFAULT 0 NOT NULL,
  updated_at BIGINT DEFAULT 0 NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users (id)
);
COMMIT;