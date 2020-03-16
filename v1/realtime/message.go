package realtime

import (
	"github.com/zsp2088dev/go-bitflyer/auth"
	"time"
)

type SubscribeMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
	} `json:"params"`
}

func NewSubscribeMessage(channel Channel) *SubscribeMessage {
	return &SubscribeMessage{
		JsonRPC: "2.0",
		Method:  "subscribe",
		Params: struct {
			Channel string `json:"channel"`
		}{Channel: channel.String()},
	}
}

type AuthMessage struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  auth.Params `json:"params"`
	ID      int         `json:"id"`
}

func NewAuthMessage(param *auth.Auth) *AuthMessage {
	return &AuthMessage{
		JsonRPC: "2.0",
		Method:  "auth",
		Params:  auth.CreateAuthParams(param),
		ID:      int(time.Now().Unix()),
	}
}
