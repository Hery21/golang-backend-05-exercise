--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)

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
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    amount integer NOT NULL,
    transaction_type character varying NOT NULL,
    wallet_id integer NOT NULL,
    target_wallet_id integer,
    description character varying(34),
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets (
    id integer NOT NULL,
    user_id integer,
    balance bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.wallets OWNER TO postgres;

--
-- Name: wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_id_seq OWNER TO postgres;

--
-- Name: wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallets_id_seq OWNED BY public.wallets.id;


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: wallets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets ALTER COLUMN id SET DEFAULT nextval('public.wallets_id_seq'::regclass);


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, amount, transaction_type, wallet_id, target_wallet_id, description, created_at, deleted_at, updated_at) FROM stdin;
1	50000	CREDIT	100003	1001	Top Up from Bank Transfer	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.180473
2	100000	CREDIT	100003	1001	Top Up from Bank Transfer	2021-06-01 06:54:20	\N	2022-08-11 13:46:30.180473
3	200000	CREDIT	100003	1001	Top Up from Bank Transfer	2022-07-01 06:54:20	\N	2022-08-11 13:46:30.180473
4	400000	CREDIT	100009	1001	Top Up from Bank Transfer	2022-08-01 06:54:20	\N	2022-08-11 13:46:30.180473
5	800000	CREDIT	100003	1002	Top Up from Credit Card	2020-01-02 06:54:20	\N	2022-08-11 13:46:30.180473
6	1600000	CREDIT	100003	1003	Top Up from Credit Card	2021-06-02 06:54:20	\N	2022-08-11 13:46:30.180473
7	3200000	CREDIT	100007	1003	Top Up from Credit Card	2022-07-02 06:54:20	\N	2022-08-11 13:46:30.180473
8	6400000	CREDIT	100006	1003	Top Up from Credit Card	2022-08-02 06:54:20	\N	2022-08-11 13:46:30.180473
9	10000000	CREDIT	100008	1002	Top Up from Cash	2020-01-03 06:54:20	\N	2022-08-11 13:46:30.180473
10	10000000	CREDIT	100009	1002	Top Up from Cash	2021-06-04 06:54:20	\N	2022-08-11 13:46:30.180473
11	10000000	CREDIT	100011	1002	Top Up from Cash	2022-07-03 06:54:20	\N	2022-08-11 13:46:30.180473
12	10000000	CREDIT	100010	1002	Top Up from Cash	2022-08-04 06:54:20	\N	2022-08-11 13:46:30.180473
13	5000	DEBIT	100003	100006	Uang jajan	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.180473
14	50000	DEBIT	100007	100003	Uang transport	2021-06-01 06:54:20	\N	2022-08-11 13:46:30.180473
15	500000	DEBIT	100003	100009	Uang wifi	2022-07-01 06:54:20	\N	2022-08-11 13:46:30.180473
16	5000000	DEBIT	100010	100008	Gaji	2022-08-09 06:54:20	\N	2022-08-11 13:46:30.180473
17	50000000	DEBIT	100006	100007	Uang kos	2021-04-01 06:54:20	\N	2022-08-11 13:46:30.180473
18	1000	DEBIT	100011	100010	Kopi	2020-06-01 06:54:20	\N	2022-08-11 13:46:30.180473
19	10000	DEBIT	100008	100007	Patungan	2022-04-01 06:54:20	\N	2022-08-11 13:46:30.180473
20	100000	DEBIT	100003	100009	Nonton	2021-03-01 06:54:20	\N	2022-08-11 13:46:30.180473
21	1000000	DEBIT	100007	100003	Makan	2022-08-02 06:54:20	\N	2022-08-11 13:46:30.180473
22	10000000	DEBIT	100006	100011	Beli buku	2020-08-11 06:54:20	\N	2022-08-11 13:46:30.180473
23	42069000	DEBIT	100003	100006	Uang sekolah	2022-07-01 06:54:20	\N	2022-08-11 13:46:30.180473
24	2000	DEBIT	100008	100009	Kopi	2020-06-01 06:54:20	\N	2022-08-11 13:46:30.180473
25	20000	DEBIT	100006	100003	Patungan	2022-04-01 06:54:20	\N	2022-08-11 13:46:30.180473
26	200000	DEBIT	100008	100006	Nonton	2021-03-01 06:54:20	\N	2022-08-11 13:46:30.180473
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, created_at, deleted_at, updated_at) FROM stdin;
1	user1	user1@mail.com	$2a$10$9UQWAf38vefdy8iOuotD9ePlYWbOKfOZBXXWHYqA01nquGR5ezQdi	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
2	user2	user2@mail.com	$2a$10$OMzIqhNruO7g0D/BCXtSl.QXz8/oeGAnT8pPIMwneGYIE0XWeCy5q	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
3	user3	user3@mail.com	$2a$10$IJqkECfzjbJvzFZbsQlvDeLsWJXIM8cb/LH93IDb5vWh5FmVfkDhK	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
4	user4	user4@mail.com	$2a$10$ofGJvPt52xS4esuA6kaLg.Dwphio8pqWVBlH3YIIrCCBUdRLn89ri	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
5	user5	user5@mail.com	$2a$10$YkLFHhApLAD.3.C5N4qRgu7bOsb3orgDTKkjBiP8et39Cz9GTcEOW	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
6	user6	user6@mail.com	$2a$10$hWxVj.L.0ci2vhksiGWdQu2TEOH1dBFq8WixfZ4jqglWHlkUH./PG	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
7	user7	user7@mail.com	$2a$10$3B0/EDF2R0Xzh/9ZJ.xHN.zglBak3XEpuJMxbzlXFc051ZXs.sbw6	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
8	user8	user8@mail.com	$2a$10$kv4VZ4fgutBC1zyNVxRKQe5JccWf.Hhc.ABU6l6zUKpkCzdA94NaK	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
9	user9	user9@mail.com	$2a$10$y/YLiOGR9El/a/3GFj2hoeUP4IoajdSXKrIJC5FcEK.Ojjmjr8Jyi	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
10	user10	user10@mail.com	$2a$10$xglo0/Coz7zNsiBDxvBgn.H1NCQPFT8Tcdj4qOGVz6aqgA9oXLpeK	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
11	user11	user11@mail.com	$2a$10$SHT0U/vo1eNE.nfByR4o7O8uAbOdSw9.FBC/KgfCh3omzkN8TRjKG	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.150624
\.


--
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wallets (id, user_id, balance, created_at, deleted_at, updated_at) FROM stdin;
100001	1	50000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100002	2	5000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100003	3	1000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100004	4	500000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100005	5	500000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100006	6	10000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100007	7	5000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100008	8	1000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100009	9	500000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100010	10	500000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
100011	11	10000000	2020-01-01 06:54:20	\N	2022-08-11 13:46:30.159318
1001	\N	999999	2022-08-11 13:46:30.164996	\N	2022-08-11 13:46:30.164996
1002	\N	999999	2022-08-11 13:46:30.168915	\N	2022-08-11 13:46:30.168915
1003	\N	999999	2022-08-11 13:46:30.17273	\N	2022-08-11 13:46:30.17273
\.


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 26, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_id_seq', 100001, false);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_target_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_target_wallet_id_fkey FOREIGN KEY (target_wallet_id) REFERENCES public.wallets(id);


--
-- Name: transactions transactions_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_wallet_id_fkey FOREIGN KEY (wallet_id) REFERENCES public.wallets(id);


--
-- Name: wallets wallets_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

