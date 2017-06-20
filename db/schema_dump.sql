--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.3
-- Dumped by pg_dump version 9.6.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: api_keys_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE api_keys_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE api_keys_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: api_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE api_keys (
    id integer DEFAULT nextval('api_keys_id_seq'::regclass) NOT NULL,
    key character varying(128) NOT NULL,
    admin boolean DEFAULT false,
    channel_modify integer,
    channel_add_message integer
);


ALTER TABLE api_keys OWNER TO postgres;

--
-- Name: channel_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE channel_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE channel_id_seq OWNER TO postgres;

--
-- Name: channels; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE channels (
    id integer DEFAULT nextval('channel_id_seq'::regclass) NOT NULL,
    name character varying(128) NOT NULL,
    string_id character varying(64) NOT NULL
);


ALTER TABLE channels OWNER TO postgres;

--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE goose_db_version OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goose_db_version_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goose_db_version_id_seq OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE goose_db_version_id_seq OWNED BY goose_db_version.id;


--
-- Name: message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE message_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE message_id_seq OWNER TO postgres;

--
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE messages (
    id integer DEFAULT nextval('message_id_seq'::regclass) NOT NULL,
    body text NOT NULL,
    "timestamp" timestamp without time zone NOT NULL,
    hash bigint DEFAULT 0,
    channel character varying(64),
    sender character varying(64)
);


ALTER TABLE messages OWNER TO postgres;

--
-- Name: senders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE senders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE senders_id_seq OWNER TO postgres;

--
-- Name: senders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE senders (
    id integer DEFAULT nextval('senders_id_seq'::regclass) NOT NULL,
    string_id character varying(64) NOT NULL,
    banned boolean DEFAULT false,
    ban_expire_time timestamp without time zone DEFAULT to_timestamp((0)::double precision),
    last_ban_length integer DEFAULT 0,
    channel character varying(64)
);


ALTER TABLE senders OWNER TO postgres;

--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goose_db_version ALTER COLUMN id SET DEFAULT nextval('goose_db_version_id_seq'::regclass);


--
-- Name: api_keys api_keys_key_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY api_keys
    ADD CONSTRAINT api_keys_key_key UNIQUE (key);


--
-- Name: api_keys api_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY api_keys
    ADD CONSTRAINT api_keys_pkey PRIMARY KEY (id);


--
-- Name: channels channels_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY channels
    ADD CONSTRAINT channels_pkey PRIMARY KEY (id);


--
-- Name: channels channels_string_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY channels
    ADD CONSTRAINT channels_string_id_key UNIQUE (string_id);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: senders senders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY senders
    ADD CONSTRAINT senders_pkey PRIMARY KEY (id);


--
-- Name: senders senders_string_id_channel_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY senders
    ADD CONSTRAINT senders_string_id_channel_key UNIQUE (string_id, channel);


--
-- Name: senders senders_string_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY senders
    ADD CONSTRAINT senders_string_id_key UNIQUE (string_id);


--
-- Name: messages_hash_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX messages_hash_idx ON messages USING btree (hash);


--
-- Name: api_keys api_keys_channel_add_message_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY api_keys
    ADD CONSTRAINT api_keys_channel_add_message_fkey FOREIGN KEY (channel_add_message) REFERENCES channels(id) ON DELETE CASCADE;


--
-- Name: api_keys api_keys_channel_modify_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY api_keys
    ADD CONSTRAINT api_keys_channel_modify_fkey FOREIGN KEY (channel_modify) REFERENCES channels(id) ON DELETE CASCADE;


--
-- Name: messages messages_channel_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_channel_fkey FOREIGN KEY (channel) REFERENCES channels(string_id) ON DELETE CASCADE;


--
-- Name: messages messages_sender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_sender_fkey FOREIGN KEY (sender) REFERENCES senders(string_id) ON DELETE CASCADE;


--
-- Name: senders senders_channel_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY senders
    ADD CONSTRAINT senders_channel_fkey FOREIGN KEY (channel) REFERENCES channels(string_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

