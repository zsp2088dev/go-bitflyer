package realtime

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"github.com/json-iterator/go"
	"github.com/zsp2088dev/go-bitflyer/auth"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const BaseURL = "wss://ws.lightstream.bitflyer.com/json-rpc"

type JsonRpcClient struct {
	Channels []Channel
	auth     *auth.Auth
}

func New(channels []Channel) *JsonRpcClient {
	return &JsonRpcClient{Channels: channels}
}

func (c *JsonRpcClient) Auth(auth *auth.Auth) *JsonRpcClient {
	c.auth = auth
	return c
}

func (c *JsonRpcClient) ReceiveMessage(ch chan Response) {
	conn, _, err := websocket.DefaultDialer.Dial(BaseURL, nil)
	if err != nil {
		ch <- Response{Error: err}
		return
	}
	defer conn.Close()

	// Privateチャンネルを利用する場合は認証を先に行う
	if c.auth != nil {
		if err := requestAuth(conn, c.auth); err != nil {
			ch <- Response{Error: err}
			return
		}
	}

	// 引数のチャンネルを購読する
	for _, channel := range c.Channels {
		msg := NewSubscribeMessage(channel)
		if err := conn.WriteJSON(msg); err != nil {
			ch <- Response{Error: err}
			return
		}
	}

	// メッセージを受信
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			ch <- Response{Error: err}
			return
		}

		// byte型のJSONからチャンネルを抽出
		channel, _ := jsonparser.GetString(msg, "params", "channel")
		event, err := pickUpEvent(channel)
		if err != nil {
			ch <- Response{Error: err}
			return
		}

		switch event {
		case BoardSnapshot:
			res := new(BoardResponse)
			_ = json.Unmarshal(msg, res)
			ch <- Response{
				Event:         BoardSnapshot,
				BoardSnapshot: res.Params.Message,
			}
		case Board:
			res := new(BoardResponse)
			_ = json.Unmarshal(msg, res)
			ch <- Response{
				Event: Board,
				Board: res.Params.Message,
			}
		case Ticker:
			res := new(TickerResponse)
			_ = json.Unmarshal(msg, res)
			ch <- Response{
				Event:  Ticker,
				Ticker: res.Params.Message,
			}
		case Executions:
			res := new(ExecutionsResponse)
			_ = json.Unmarshal(msg, res)
			ch <- Response{
				Event:      Executions,
				Executions: res.Params.Message,
			}
		case ChildOrderEvents:
			res := new(ChildOrderEventsResponse)
			_ = json.Unmarshal(msg, res)
			ch <- Response{
				Event:            ChildOrderEvents,
				ChildOrderEvents: res.Params.Message,
			}
		case ParentOrderEvents:
			res := new(ParentOrderEventsResponse)
			_ = json.Unmarshal(msg, res)
			ch <- Response{
				Event:             ChildOrderEvents,
				ParentOrderEvents: res.Params.Message,
			}
		default:
			ch <- Response{Error: fmt.Errorf("unexpected event type: %s", event)}
			return
		}
	}
}

func requestAuth(conn *websocket.Conn, auth *auth.Auth) error {
	// 認証情報をリクエスト
	if err := conn.WriteJSON(NewAuthMessage(auth)); err != nil {
		return err
	}

	// 認証の成功を確認
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	res, _ := jsonparser.GetBoolean(msg, "result")
	if !res {
		return err
	}
	return nil
}
