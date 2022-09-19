FROM golang:1.18-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /api-server

EXPOSE 8080

CMD [ "/api-server" ]