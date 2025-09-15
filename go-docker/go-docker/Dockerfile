# syntax=docker/dockerfile:1
FROM golang:1.25 AS dev

# install Air buat hot-reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# cache deps
COPY go.mod go.sum* ./
RUN go mod download || true

# copy source
COPY . .

ENV GO_ENV=development
EXPOSE 8080

# jalanin Air (baca .air.toml)
CMD ["air", "-c", ".air.toml"]