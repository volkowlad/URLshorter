CREATE TABLE url
(
    id SERIAL,
    alias TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL
);