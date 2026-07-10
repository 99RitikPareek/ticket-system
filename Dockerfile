FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ticket-system ./cmd

EXPOSE 8080

CMD ["./ticket-system"]