FROM golang:alpine AS builder

LABEL maintainer="alejoruasuarez@gmail.com"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -installsuffix 'static' -o main

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
