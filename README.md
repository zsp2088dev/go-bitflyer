<div align="center">
<img src="https://user-images.githubusercontent.com/33179078/76741787-781a0700-67b3-11ea-9fc9-78d01d5a3faa.png" alt="image" title="Go bitFlyer"><br><br>
</div>

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)


## インストール
`$ go get github.com/zsp2088dev/go-bitflyer`

## 動作環境
- Go 1.12

## 使用方法
```go
package main

import (
	"fmt"
	"github.com/zsp2088dev/go-bitflyer/auth"
	"github.com/zsp2088dev/go-bitflyer/v1"
	"github.com/zsp2088dev/go-bitflyer/v1/private/positions"
	"github.com/zsp2088dev/go-bitflyer/v1/public/ticker"
	"github.com/zsp2088dev/go-bitflyer/v1/realtime"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"log"
)

func main() {
	config := &auth.Auth{
		APIKey:    "<API KEY>",
		APISecret: "<API SECRET>",
	}

	// Private APIを利用しない場合はconfigはnilでよい
	client := v1.NewClient(config)

	// HTTP Public API
	req1 := &ticker.Request{ProductCode: types.BTCJPY}
	res1, raw1, err := client.Ticker(req1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n\n", res1) // Ticker
		fmt.Printf("%v\n\n", raw1)
	}

	// HTTP Private API
	req2 := &positions.Request{ProductCode: types.FXBTCJPY}
	res2, raw2, err := client.Positions(req2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n\n", res2) // 建玉一覧
		fmt.Printf("%v\n\n", raw2)
	}

	// Realtime API
	channel1 := realtime.Channel{ProductCode: types.FXBTCJPY, Event: realtime.Executions}
	channel2 := realtime.Channel{Event: realtime.ChildOrderEvents}

	// Privateチャンネルを購読する場合は ".Auth(config)" を呼び出す
	wsClient := realtime.
        New().
		Auth(config).
		AddChannel(channel1).
		AddChannel(channel2)

	// realtime.Response型で各種メッセージを受信
	ch := make(chan realtime.Response)
	go wsClient.ReceiveMessage(ch)

	for {
		msg := <-ch
		if msg.Error != nil {
			log.Println(msg.Error)
		}

		// 受信したイベントごとに処理
		switch msg.Event {
		case realtime.Executions:
			fmt.Printf("%v\n\n", msg.Executions)

		case realtime.ChildOrderEvents:
			fmt.Printf("%v\n\n", msg.ChildOrderEvents)
		}
	}
}

```

## ライセンス
Released under the MIT license.<br>
Copyright (c) 2020, Kazuki Yamagata

## 製作者
- Name: Kazuki Yamagata
- Twitter: [@zsp2088dev](https://twitter.com/zsp2088dev)
- Website: [Zu's Portfolio](https://zsp2088dev.netlify.com/)