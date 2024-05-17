FROM golang:1.18-alpine

RUN apk add --no-cache gcc musl-dev postgresql-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /eulabs cmd/server/main.go

EXPOSE 8080

ENV DB_TYPE=sqlite

CMD ["/eulabs"]