FROM golang:1.22-alpine

WORKDIR /app

COPY .env go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./... -v
RUN go build -o /rate-limiter

EXPOSE 8080

CMD ["/rate-limiter"]