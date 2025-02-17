CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE users (
	no_rekening varchar,
	nama varchar NOT NULL,
	nik varchar NOT NULL,
	no_hp varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (no_rekening)
);

CREATE TABLE user_saldo (
	id uuid DEFAULT public.uuid_generate_v4(),
	no_rekening varchar NOT NULL,
	saldo varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT user_saldo_pkey PRIMARY KEY (id)
);

ALTER TABLE ONLY user_saldo ADD CONSTRAINT fk_users_saldo FOREIGN KEY (no_rekening) REFERENCES users(no_rekening) ON UPDATE CASCADE ON DELETE SET NULL;
