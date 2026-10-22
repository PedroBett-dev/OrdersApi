CREATE OR REPLACE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS products (
   id VARCHAR(6) PRIMARY KEY,
   name VARCHAR(30) NOT NULL,
   price NUMERIC(10, 2) NOT NULL,
   status VARCHAR(8) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER products_set_updated_at
   BEFORE UPDATE ON products
   FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE IF NOT EXISTS customers (
   id VARCHAR(6) PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   cpf VARCHAR(11) NOT NULL UNIQUE,
   phone_number VARCHAR(11) NOT NULL,
   cep VARCHAR(8) NOT NULL,
   street VARCHAR(100) NOT NULL,
   number VARCHAR(10) NOT NULL,
   complement VARCHAR(50),
   neighborhood VARCHAR(50) NOT NULL,
   city VARCHAR(50) NOT NULL,
   state CHAR(2) NOT NULL,

   status VARCHAR(8) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER customers_set_updated_at
   BEFORE UPDATE ON customers
   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
