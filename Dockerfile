FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go-weather-challenge ./cmd//server/main.go

FROM alpine:3.18
WORKDIR /
COPY --from=builder /go-weather-challenge /go-weather-challenge
EXPOSE 8080
CMD ["/go-weather-challenge"]