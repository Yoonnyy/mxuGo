SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: files; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.files (
    id integer NOT NULL,
    "originalFilename" character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    size integer NOT NULL,
    expires bigint NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: shortened; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.shortened (
    id integer NOT NULL,
    "isFile" boolean NOT NULL,
    slug character varying(255) NOT NULL
);


--
-- Name: urls; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.urls (
    id integer NOT NULL,
    slug character varying(255) NOT NULL,
    destination character varying(1024) NOT NULL,
    expires bigint NOT NULL
);


--
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- Name: files files_slug_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_slug_key UNIQUE (slug);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: shortened shortened_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.shortened
    ADD CONSTRAINT shortened_pkey PRIMARY KEY (id);


--
-- Name: shortened shortened_slug_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.shortened
    ADD CONSTRAINT shortened_slug_key UNIQUE (slug);


--
-- Name: urls urls_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT urls_pkey PRIMARY KEY (id);


--
-- Name: urls urls_slug_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT urls_slug_key UNIQUE (slug);


--
-- Name: files_slug_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX files_slug_idx ON public.files USING btree (slug);


--
-- Name: shortened_slug_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX shortened_slug_idx ON public.shortened USING btree (slug);


--
-- Name: urls_slug_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX urls_slug_idx ON public.urls USING btree (slug);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20230914055021');
