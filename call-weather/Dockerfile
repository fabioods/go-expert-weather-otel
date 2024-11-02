FROM golang:1.23.2-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o weather ./cmd/main.go && chmod +x weather

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/weather /app/weather
ENTRYPOINT ["/app/weather"]
