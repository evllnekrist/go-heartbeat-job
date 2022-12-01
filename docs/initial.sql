-- Database: climate

-- DROP DATABASE IF EXISTS climate;

CREATE DATABASE climate
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.weathers

-- DROP TABLE IF EXISTS public.weathers;

CREATE TABLE IF NOT EXISTS public.weathers
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    water bigint NOT NULL,
    wind bigint NOT NULL,
    CONSTRAINT weathers_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.weathers
    OWNER to postgres;

-- SEQUENCE: public.weathers_id_seq

-- DROP SEQUENCE IF EXISTS public.weathers_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.weathers_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 1000000
    CACHE 1;

ALTER SEQUENCE public.weathers_id_seq
    OWNER TO postgres;