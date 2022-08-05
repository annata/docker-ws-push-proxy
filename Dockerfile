FROM golang:1.18-alpine as builder
WORKDIR /data
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /data/ws_push_proxy /
CMD ["/ws_push_proxy"]