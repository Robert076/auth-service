# 🔐 auth-service

A standalone microservice that implements authorization using a PostgreSQL database and implemented in Go.

One of the main focuses during this project was to concentrate on making the code scalable, and that is achieved by using the strategy pattern wherever possible.

### 🚀 Run command

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

> ❗️ Will be documented once the first release is out.
