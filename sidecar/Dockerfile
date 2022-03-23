FROM golang:1.17.5-alpine3.15 AS builder

COPY . /github.com/olezhek28/container-patterns/sidecar/
WORKDIR /github.com/olezhek28/container-patterns/sidecar/

RUN go mod download
RUN go build -o ./bin/http_server cmd/http_server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/olezhek28/container-patterns/sidecar/bin/http_server .
EXPOSE 8093

CMD ["./http_server"]