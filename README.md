### assignment-05-golang-backend 
To access:
- /sign-in
- /register
- /wallets
- /wallets/transfer
- /wallets/top-up
- /wallets/transactions?s=&sortBy=&sort=&limit=

To initiate database:

CREATE TABLE users (
                       id SERIAL UNIQUE NOT NULL PRIMARY KEY,
                       name VARCHAR NOT NULL,
                       email VARCHAR UNIQUE NOT NULL,
                       password VARCHAR NOT NULL,
                       created_at TIMESTAMP DEFAULT now(),
                       deleted_at TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE wallets (
                         id SERIAL UNIQUE NOT NULL PRIMARY KEY,
                         user_id INT REFERENCES users(id),
                         balance BIGINT NOT NULL,
                         created_at TIMESTAMP DEFAULT now(),
                         deleted_at TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE transactions (
                              id SERIAL UNIQUE PRIMARY KEY,
                              amount INT NOT NULL,
                              transaction_type VARCHAR NOT NULL,
                              wallet_id INT NOT NULL REFERENCES wallets(id),
                              target_wallet_id INT REFERENCES wallets(id),
                              description VARCHAR(34),
                              created_at TIMESTAMP DEFAULT now(),
                              deleted_at TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT now()
);

ALTER SEQUENCE wallets_id_seq RESTART WITH 100001;

INSERT INTO users (name, email, password, created_at)
VALUES ('user1', 'user1@mail.com', '$2a$10$9UQWAf38vefdy8iOuotD9ePlYWbOKfOZBXXWHYqA01nquGR5ezQdi', '2020-01-01 06:54:20'),
       ('user2', 'user2@mail.com', '$2a$10$OMzIqhNruO7g0D/BCXtSl.QXz8/oeGAnT8pPIMwneGYIE0XWeCy5q', '2020-01-01 06:54:20'),
       ('user3', 'user3@mail.com', '$2a$10$IJqkECfzjbJvzFZbsQlvDeLsWJXIM8cb/LH93IDb5vWh5FmVfkDhK', '2020-01-01 06:54:20'),
       ('user4', 'user4@mail.com', '$2a$10$ofGJvPt52xS4esuA6kaLg.Dwphio8pqWVBlH3YIIrCCBUdRLn89ri', '2020-01-01 06:54:20'),
       ('user5', 'user5@mail.com', '$2a$10$YkLFHhApLAD.3.C5N4qRgu7bOsb3orgDTKkjBiP8et39Cz9GTcEOW', '2020-01-01 06:54:20'),
       ('user6', 'user6@mail.com', '$2a$10$hWxVj.L.0ci2vhksiGWdQu2TEOH1dBFq8WixfZ4jqglWHlkUH./PG', '2020-01-01 06:54:20'),
       ('user7', 'user7@mail.com', '$2a$10$3B0/EDF2R0Xzh/9ZJ.xHN.zglBak3XEpuJMxbzlXFc051ZXs.sbw6', '2020-01-01 06:54:20'),
       ('user8', 'user8@mail.com', '$2a$10$kv4VZ4fgutBC1zyNVxRKQe5JccWf.Hhc.ABU6l6zUKpkCzdA94NaK', '2020-01-01 06:54:20'),
       ('user9', 'user9@mail.com', '$2a$10$y/YLiOGR9El/a/3GFj2hoeUP4IoajdSXKrIJC5FcEK.Ojjmjr8Jyi', '2020-01-01 06:54:20'),
       ('user10', 'user10@mail.com', '$2a$10$xglo0/Coz7zNsiBDxvBgn.H1NCQPFT8Tcdj4qOGVz6aqgA9oXLpeK', '2020-01-01 06:54:20'),
       ('user11', 'user11@mail.com', '$2a$10$SHT0U/vo1eNE.nfByR4o7O8uAbOdSw9.FBC/KgfCh3omzkN8TRjKG', '2020-01-01 06:54:20');

INSERT INTO wallets (user_id, balance, created_at)
VALUES (1, 50000000, '2020-01-01 06:54:20'),
       (2, 5000000, '2020-01-01 06:54:20'),
       (3, 1000000, '2020-01-01 06:54:20'),
       (4, 500000000, '2020-01-01 06:54:20'),
       (5, 500000, '2020-01-01 06:54:20'),
       (6, 10000000, '2020-01-01 06:54:20'),
       (7, 5000000, '2020-01-01 06:54:20'),
       (8, 1000000, '2020-01-01 06:54:20'),
       (9, 500000000, '2020-01-01 06:54:20'),
       (10, 500000, '2020-01-01 06:54:20'),
       (11, 10000000, '2020-01-01 06:54:20');

INSERT INTO wallets (id, balance)
VALUES (1001, 999999);

INSERT INTO wallets (id, balance)
VALUES (1002, 999999);

INSERT INTO wallets (id, balance)
VALUES (1003, 999999);

INSERT INTO transactions (amount, transaction_type, wallet_id, target_wallet_id, description, created_at)
VALUES (50000, 'CREDIT', 100003, 1001, 'Top Up from Bank Transfer', '2020-01-01 06:54:20'),
       (100000, 'CREDIT', 100003, 1001, 'Top Up from Bank Transfer', '2021-06-01 06:54:20'),
       (200000, 'CREDIT', 100003, 1001, 'Top Up from Bank Transfer', '2022-07-01 06:54:20'),
       (400000, 'CREDIT', 100009, 1001, 'Top Up from Bank Transfer', '2022-08-01 06:54:20'),
       (800000, 'CREDIT', 100003, 1002, 'Top Up from Credit Card', '2020-01-02 06:54:20'),
       (1600000, 'CREDIT', 100003, 1003, 'Top Up from Credit Card', '2021-06-02 06:54:20'),
       (3200000, 'CREDIT', 100007, 1003, 'Top Up from Credit Card', '2022-07-02 06:54:20'),
       (6400000, 'CREDIT', 100006, 1003, 'Top Up from Credit Card', '2022-08-02 06:54:20'),
       (10000000, 'CREDIT', 100008, 1002, 'Top Up from Cash', '2020-01-03 06:54:20'),
       (10000000, 'CREDIT', 100009, 1002, 'Top Up from Cash', '2021-06-04 06:54:20'),
       (10000000, 'CREDIT', 100011, 1002, 'Top Up from Cash', '2022-07-03 06:54:20'),
       (10000000, 'CREDIT', 100010, 1002, 'Top Up from Cash', '2022-08-04 06:54:20'),
       (5000, 'DEBIT', 100003, 100006, 'Uang jajan', '2020-01-01 06:54:20'),
       (50000, 'DEBIT', 100007, 100003, 'Uang transport', '2021-06-01 06:54:20'),
       (500000, 'DEBIT', 100003, 100009, 'Uang wifi', '2022-07-01 06:54:20'),
       (5000000, 'DEBIT', 100010, 100008, 'Gaji', '2022-08-09 06:54:20'),
       (50000000, 'DEBIT', 100006, 100007, 'Uang kos', '2021-04-01 06:54:20'),
       (1000, 'DEBIT', 100011, 100010, 'Kopi', '2020-06-01 06:54:20'),
       (10000, 'DEBIT', 100008, 100007, 'Patungan', '2022-04-01 06:54:20'),
       (100000, 'DEBIT', 100003, 100009, 'Nonton', '2021-03-01 06:54:20'),
       (1000000, 'DEBIT', 100007, 100003, 'Makan', '2022-08-02 06:54:20'),
       (10000000, 'DEBIT', 100006, 100011, 'Beli buku', '2020-08-11 06:54:20'),
       (42069000, 'DEBIT', 100003, 100006, 'Uang sekolah', '2022-07-01 06:54:20'),
       (2000, 'DEBIT', 100008, 100009, 'Kopi', '2020-06-01 06:54:20'),
       (20000, 'DEBIT', 100006, 100003, 'Patungan', '2022-04-01 06:54:20'),
       (200000, 'DEBIT', 100008, 100006, 'Nonton', '2021-03-01 06:54:20');

