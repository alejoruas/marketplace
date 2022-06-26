CREATE DATABASE postgres
    WITH
    OWNER = spuser
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

COMMENT ON DATABASE postgres
    IS 'default administrative connection database';


-- Table: public.Project


CREATE TABLE IF NOT EXISTS public."Project"
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    budget double precision,
    CONSTRAINT "Project_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Project"
    OWNER to spuser;


-- Table: public.DriverImputation

CREATE TABLE IF NOT EXISTS public."DriverImputation"
(
    id_project text COLLATE pg_catalog."default" NOT NULL,
    ceco character(9)[] COLLATE pg_catalog."default" NOT NULL,
    percentage double precision NOT NULL,
    cia character(2)[] COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "DriverImputation_pkey" PRIMARY KEY (id_project, ceco)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."DriverImputation"
    OWNER to spuser;