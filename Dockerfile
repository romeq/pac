FROM golang:1.20-alpine3.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

RUN apk update && apk add postgresql-client make

CMD ["make", "migrate", "run"]
