# Build: docker build -t go-docker .
FROM golang:1.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

# Prod
FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]