BEGIN;

CREATE TABLE IF NOT EXISTS public.client
(
   id SERIAL PRIMARY KEY,
   clientId character varying(50) NOT NULL UNIQUE,
   name character varying(50) NOT NULL
)
WITH (
  OIDS = FALSE
)
;
ALTER TABLE public.client
  OWNER TO imqs;

CREATE TABLE IF NOT EXISTS public.type
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

CREATE TABLE IF NOT EXISTS public.messages
(
   id SERIAL PRIMARY KEY,
   clientId character varying(50) NOT NULL REFERENCES public.client(clientId),
   type character varying(50) NOT NULL REFERENCES public.type(name),
   message jsonb NOT NULL
)
WITH (
  OIDS = FALSE
)
;
ALTER TABLE public.messages
  OWNER TO imqs;

COMMIT;
