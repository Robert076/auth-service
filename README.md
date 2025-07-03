# 🔐 auth-service

A standalone microservice that implements authorization using a PostgreSQL database and implemented in Go.

Just clone the repo and run it on your machine, plug in your database and just have fun with it.

If you have a different users table please do modify the register DTO to include everything you need. Apart from that you can easily change the db type since it's using the strategy pattern making adjusting the db easy.

## ✍🏻 Diagram

<img width="600" alt="Image" src="https://github.com/user-attachments/assets/df9176bd-2de7-420d-803b-33977e86e2f5" />

---

## 🚀 Run command

```bash
docker compose up --build
```

### ⚙️ Example .env file

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

Make a `POST` request to the `/login` endpoint with email and password, and if password hash matches (from db) it returns 200 (this is how you will know the login was succesful)

Make a `POST` request to the `/authorize` endpoint with the email attached and if the session token + csrf token match (include csrf token in header when making the request) you get 200

Make a `POST` request to `/logout` with the email attached. At the moment it doesn't check for authorization.

## 🧩 Adding another database

You can easily swap out db's with one another since the code is not coupled to a certain database, it actually makes use of interfaces using the strategy pattern to enable you to add whatever db you prefer.
