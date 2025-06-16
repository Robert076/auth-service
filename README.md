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
DB_HOST=auth-service
DB_USER=admin
DB_PASSWORD=admin
DB_PORT=5432
DB_NAME=users
ENVIRONMENT=PRODUCTION
```

> ❗️ Will be documented once the first release is out.
