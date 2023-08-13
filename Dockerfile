FROM golang:latest

WORKDIR /app

COPY . .

RUN go get ./...

RUN CGO_ENABLED=1 go build -o server -ldflags "-s -w" cmd/app/main.go

ENTRYPOINT [ "/app/server" ]