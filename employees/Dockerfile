FROM golang:1.27-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app .

CMD ["./app"]