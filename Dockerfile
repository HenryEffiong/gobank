FROM golang:1.22.2-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate 
COPY app.env .
COPY wait-for.sh .
COPY start.sh .
COPY db/migration ./migration
  
EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]