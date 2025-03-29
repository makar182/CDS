CREATE DATABASE client_db;
CREATE DATABASE logs_db;

\c client_db
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name TEXT NOT NULL,
                       last_name TEXT NOT NULL,
                       middle_name TEXT NOT NULL
);

\c logs_db
CREATE TABLE logs (
                      id SERIAL PRIMARY KEY,
                      message TEXT NOT NULL
);
