FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /main main

ENV PORT="8080"
EXPOSE 8080

CMD ["/app/main"]
