CREATE DATABASE authservicedb

\c authservicedb

CREATE TABLE "Users"(
    "Id" SERIAL PRIMARY KEY,
    "Name" VARCHAR(50),
    "Email" VARCHAR(100) UNIQUE
);