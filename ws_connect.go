package main

import (
	"golang.org/x/net/websocket"
	"strconv"
	"sync/atomic"
)

func wsConnect(closeFlag <-chan any, ws *websocket.Conn) {
	sn := strconv.FormatUint(atomic.AddUint64(&snn, 1), 10)
	topic := ws.Request().RequestURI
	addTopic(topic, sn, ws)
	defer removeTopic(topic, sn)
	<-closeFlag
}
