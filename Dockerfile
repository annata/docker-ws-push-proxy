FROM golang:1.20-alpine as builder
WORKDIR /data
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags="-w -s"

FROM buildpack-deps:bullseye-curl
COPY --from=builder /data/ws_push_proxy /
CMD ["/ws_push_proxy"]