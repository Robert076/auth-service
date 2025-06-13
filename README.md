# 🔐 auth-service

A standalone microservice that implements authorization using a PostgreSQL database and implemented in Go.

### 🚀 Run command

```bash
docker run -p 5656:5656 --env-file ./.env robert076/auth-service:alpha
```

### Example .env file

```
ENDPOINT_PORT=5656<br>
DB_HOST=auth-service<br>
DB_USER=admin<br>
DB_PASSWORD=admin<br>
DB_PORT=5432<br>
DB_NAME=users<br>
ENVIRONMENT=PRODUCTION<br>
```

> ❗️ Will be documented once the first release is out.
