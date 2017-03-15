-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE public.jobs
(
  id             INTEGER NOT NULL,
  container_name TEXT    NOT NULL,
  frequency      INTEGER NOT NULL,
  config         JSON    NOT NULL,
  enabled        BOOLEAN NOT NULL,
  recipient      TEXT    NOT NULL,
  PRIMARY KEY (id)
)
WITH (
OIDS = FALSE
);

ALTER TABLE public.jobs
  OWNER TO postgres;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE public.jobs;
