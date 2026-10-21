FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/api ./cmd/api

FROM alpine:3.20
WORKDIR /app
COPY --from=build /bin/api /app/api
COPY migrations /app/migrations
ENTRYPOINT ["/app/api"]
