CREATE TABLE users (
  id VARCHAR(32),
  name VARCHAR(50) NOT NULL,
  cpf VARCHAR(11),
  cnpj VARCHAR(14),
  email VARCHAR(255),
  password VARCHAR(255),
  created_at BIGINT DEFAULT 0 NOT NULL,
  PRIMARY KEY(id),
  UNIQUE KEY uk_email_cpf_cnpj(email, cpf, cnpj)
);