FROM golang:1.21.3-alpine

RUN mkdir /app
WORKDIR /app

EXPOSE 8080

COPY src/ .
RUN go mod download

CMD ["go", "run", "cmd/main.go"]