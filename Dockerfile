FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY config.yml .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

ENV REDIS_HOST localhost:6379
ENV REDIS_PASSWORD ""
ENV REDIS_DB 0

CMD ["./main"]