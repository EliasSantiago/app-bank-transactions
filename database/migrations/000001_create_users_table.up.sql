BEGIN;
CREATE TABLE users (
  id VARCHAR(32),
  name VARCHAR(50) NOT NULL,
  cpf VARCHAR(11),
  cnpj VARCHAR(14),
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at BIGINT DEFAULT 0 NOT NULL,
  PRIMARY KEY(id),
  UNIQUE (email, cpf, cnpj)
);
COMMIT;