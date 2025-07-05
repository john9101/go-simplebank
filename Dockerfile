# build stage
FROM golang:tip-alpine3.22 AS builder
WORKDIR /app
COPY . . 
RUN go build -o main main.go
RUN apk add curl

# run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]