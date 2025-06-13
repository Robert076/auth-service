FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/auth-service .

RUN chmod +x /src/auth-service

RUN apk --no-cache add ca-certificates

EXPOSE 5656

ENTRYPOINT [ "/src/auth-service" ]
