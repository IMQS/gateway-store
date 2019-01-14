BEGIN;
CREATE TABLE IF NOT EXISTS public.messages
(
   id SERIAL PRIMARY KEY,
   clientId character varying(50) NOT NULL,
   type character varying(50),
   message jsonb NOT NULL
)
WITH (
  OIDS = FALSE
)
;
ALTER TABLE public.messages
  OWNER TO imqs;
COMMIT;
