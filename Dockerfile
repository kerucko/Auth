FROM golang:1.21.5-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go build -o app ./cmd/main.go

CMD ["./app"]