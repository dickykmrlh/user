FROM golang:1.20-alpine

WORKDIR /app
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/user cmd/*.go

EXPOSE 8080 8081
CMD ["./bin/user"]