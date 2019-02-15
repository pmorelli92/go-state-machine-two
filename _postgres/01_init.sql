CREATE SCHEMA go;

CREATE TABLE go.vehicles
(
  id uuid NOT NULL,
  battery numeric NOT NULL,
  current_state text COLLATE pg_catalog."default" NOT NULL,
  last_change_state timestamp without time zone NOT NULL,
  CONSTRAINT vehicles_pkey PRIMARY KEY (id)
)