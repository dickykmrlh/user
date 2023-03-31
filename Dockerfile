FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o  /bin/user cmd/*.go

FROM alpine
COPY --from=builder /bin/user /bin/user

EXPOSE 8080 8081
ENTRYPOINT ["./bin/user"]