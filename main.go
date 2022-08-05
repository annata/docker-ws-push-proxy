package main

import (
	"context"
	"flag"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var port = ""
var url = ""
var ctx context.Context

func main() {
	ctx = context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	if !parse() {
		flag.Usage()
		return
	}
	server := &http.Server{Addr: ":" + port, Handler: websocket.Handler(websocketHandle)}
	go stopHttp(server)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func stopHttp(server *http.Server) {
	<-ctx.Done()
	server.Shutdown(context.TODO())
}

func parse() bool {
	flag.StringVar(&port, "port", "8080", "端口")
	flag.StringVar(&url, "url", "", "链接的ws的url")
	flag.Parse()
	portStr := os.Getenv("port")
	if portStr != "" {
		port = portStr
	}
	urlStr := os.Getenv("url")
	if urlStr != "" {
		url = urlStr
	}
	if url == "" {
		return false
	}
	return true
}
