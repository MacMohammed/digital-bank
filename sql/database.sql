CREATE TABLE IF NOT EXISTS public.accounts
(
    id serial,
    name character varying(120) NOT NULL,
    cpf character varying(16) NOT NULL,
    secret character varying(120) NOT NULL,
    balance numeric(15,2) DEFAULT 0,
    created_at timestamp without time zone DEFAULT now(),
    CONSTRAINT accounts_pkey PRIMARY KEY (cpf)
);

CREATE TABLE IF NOT EXISTS public.deposits
(
    id serial,
    account_origin_id integer NOT NULL,
    amount numeric(15,2) NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.login
(
    cpf character varying(16) NOT NULL,
    secret character varying(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS public.transfers
(
    id serial,
    account_origin_id integer NOT NULL,
    account_destination_id integer NOT NULL,
    amount numeric(15,2) NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL DEFAULT now()
);