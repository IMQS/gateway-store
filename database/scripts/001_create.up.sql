BEGIN;

CREATE TABLE IF NOT EXISTS client
(
   id SERIAL PRIMARY KEY,
   clientId character varying(50) NOT NULL UNIQUE,
   name character varying(50) NOT NULL,
   status character varying(10) NOT NULL
)
WITH (
  OIDS = FALSE
)
;
ALTER TABLE client
  OWNER TO imqs;

CREATE TABLE IF NOT EXISTS type
(
   id SERIAL PRIMARY KEY,
   name character varying(50) NOT NULL UNIQUE
)
WITH (
  OIDS = FALSE
)
;
ALTER TABLE public.type
  OWNER TO imqs;

CREATE TABLE IF NOT EXISTS messages
(
   id SERIAL PRIMARY KEY,
   clientId character varying(50) NOT NULL REFERENCES client(clientId),
   type character varying(50) NOT NULL REFERENCES type(name),
   message jsonb NOT NULL
)
WITH (
  OIDS = FALSE
)
;
ALTER TABLE messages
  OWNER TO imqs;

COMMIT;
