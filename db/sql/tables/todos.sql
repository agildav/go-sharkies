BEGIN TRANSACTION;

DROP TABLE IF EXISTS public.todos CASCADE;

CREATE TABLE public.todos (
	id serial NOT NULL,
	title varchar(160) NOT NULL,
	body varchar(380) NULL,
	CONSTRAINT todos_pk PRIMARY KEY (id)
);

COMMIT;