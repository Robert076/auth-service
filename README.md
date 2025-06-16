# 🔐 auth-service

A standalone microservice that implements authorization using a PostgreSQL database and implemented in Go.

Just clone the repo and run it on your machine, plug in your database and just have fun with it.

If you have a different users table please do modify the register DTO to include everything you need. Apart from that you can easily change the db type since it's using the strategy pattern making adjusting the db easy.

## 🚀 Run command

```bash
docker compose up --build
```

### Example .env file

```
ENDPOINT_PORT=5656
DB_TYPE=postgres
DB_HOST=postgres-service
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
DB_PORT=5432
POSTGRES_NAME=authservicedb
DB_SSLMODE=disable
ENVIRONMENT=PRODUCTION
```

## 💻 How it works

Make a `POST` request to the `/register` endpoint with the corresponding body (check the register dto) and the account gets created

Make a `GET` request to the `/login` endpoint with email and password, and if password hash matches (from db) it returns 200 (in the future will return a `JWT`)

## 🧩 Adding another database

You can easily swap out db's with one another since the code is not coupled to a certain database, it actually makes use of interfaces using the strategy pattern to enable you to add whatever db you prefer.
