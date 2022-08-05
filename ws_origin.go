package main

import (
	"context"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"golang.org/x/net/websocket"
	"time"
)

var wsOrigin = cmap.New[context.CancelFunc]()

func connectWsOrigin(ctx context.Context, topic string) {
	config, e := websocket.NewConfig(url+topic, "http://localhost")
	if e != nil {
		return
	}
	for {
		if ctx.Err() != nil {
			return
		}
		wsDial, err := websocket.DialConfig(config)
		if err == nil {
			for {
				if ctx.Err() != nil {
					wsDial.Close()
					return
				}
				var data = AnyMessages{}
				err := AnyMessage.Receive(wsDial, &data)
				if err != nil {
					break
				}
				if data.PayloadType == websocket.BinaryFrame {
					go sendByte(topic, data.Msg)
				} else if data.PayloadType == websocket.TextFrame {
					go sendMessage(topic, string(data.Msg))
				}
			}
			wsDial.Close()
		} else {
			fmt.Println(err.Error())
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 10):
		}
	}
}

func add(topic string) {
	c, cancel := context.WithCancel(context.TODO())
	wsOrigin.Set(topic, cancel)
	go connectWsOrigin(c, topic)
}

func del(topic string) {
	cancel, ok := wsOrigin.Get(topic)
	if ok {
		cancel()
		wsOrigin.Remove(topic)
	}
}
