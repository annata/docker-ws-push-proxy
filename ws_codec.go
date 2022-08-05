package main

import (
	"errors"
	"golang.org/x/net/websocket"
)

var AnyMessage = websocket.Codec{Marshal: websocket.Message.Marshal, Unmarshal: anyUnmarshal}

type AnyMessages struct {
	Msg         []byte
	PayloadType byte
}

func anyUnmarshal(msg []byte, payloadType byte, v interface{}) error {
	m, ok := v.(*AnyMessages)
	if !ok {
		return errors.New("not AnyMessages")
	}
	m.PayloadType = payloadType
	m.Msg = msg
	return nil
}

var WsPing = websocket.Codec{Marshal: marshal, Unmarshal: unmarshal}

var emptyMsg = make([]byte, 0)

func marshal(v interface{}) (msg []byte, payloadType byte, err error) {
	return emptyMsg, websocket.PingFrame, err
}

func unmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	return nil
}
