FROM golang:1.19-alpine

ENV PORT = 8000
WORKDIR /app/encrypter
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build
CMD go run main.go