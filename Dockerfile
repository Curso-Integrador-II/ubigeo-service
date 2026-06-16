
FROM golang:1.26.1-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o api_municipal .


FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/api_municipal .

EXPOSE 8084

CMD ["./api_municipal"]