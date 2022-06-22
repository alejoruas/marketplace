# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod /.
COPY go.sum /.
RUN go mod download

COPY . .

RUN go build -o /marketplace

EXPOSE 3000

CMD [ "/marketplace" ]
