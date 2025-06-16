CREATE DATABASE authservicedb

\c authservicedb

CREATE TABLE "Users"(
    "Id" SERIAL PRIMARY KEY,
    "Username" VARCHAR(50),
    "Email" VARCHAR(100) UNIQUE,
    "Password" TEXT
);