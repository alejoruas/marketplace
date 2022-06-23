FROM golang:latest

LABEL maintainer="alejoruasuarez@gmail.com"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD [ "./marketplace" ]

