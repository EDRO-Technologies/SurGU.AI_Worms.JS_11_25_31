FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o worm-server .

FROM scratch

WORKDIR /app

COPY --from=builder ./app/worm-server .

COPY --from=builder ./app/.env ./.env

CMD ["./worm-server"]