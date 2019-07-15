BEGIN TRANSACTION;

DROP TABLE IF EXISTS public.sharks CASCADE;

CREATE TABLE public.sharks (
  id BIGSERIAL NOT NULL,
  name varchar(160) NOT NULL,
  bname varchar(380) NOT NULL,
  description varchar(2048) NULL,
  image varchar(2048) NOT NULL,
  CONSTRAINT sharks_pk PRIMARY KEY (id)
);

COMMIT;