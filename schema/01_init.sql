CREATE DATABASE postgres;
CREATE SCHEMA public;

CREATE TABLE author
(
    id serial,
    name text,
    PRIMARY KEY (id)
);

CREATE TABLE book
(
    id      serial,
    title   text,
    author_id int,
    PRIMARY KEY (id)
);