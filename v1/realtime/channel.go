package realtime

import (
	"fmt"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"strings"
)

type Event string

const (
	Ticker            Event = "ticker"
	Board             Event = "board"
	BoardSnapshot     Event = "board_snapshot"
	Executions        Event = "executions"
	ChildOrderEvents  Event = "child_order_events"
	ParentOrderEvents Event = "parent_order_events"
)

type Channel struct {
	ProductCode types.ProductCode
	Event       Event
}

func (c *Channel) String() string {
	if c.ProductCode != "" {
		return fmt.Sprintf("lightning_%s_%s", c.Event, c.ProductCode)
	} else {
		return fmt.Sprintf("%s", c.Event)
	}
}

// 受信したメッセージのチャンネル名を抽出してEvent型に変換
func pickUpEvent(channel string) (Event, error) {
	if strings.Contains(channel, string(BoardSnapshot)) {
		return BoardSnapshot, nil
	} else if strings.Contains(channel, string(Board)) {
		return Board, nil
	} else if strings.Contains(channel, string(Ticker)) {
		return Ticker, nil
	} else if strings.Contains(channel, string(Executions)) {
		return Executions, nil
	} else if strings.Contains(channel, string(ChildOrderEvents)) {
		return ChildOrderEvents, nil
	} else if strings.Contains(channel, string(ParentOrderEvents)) {
		return ParentOrderEvents, nil
	} else {
		return "", fmt.Errorf("unexpected channel %s", channel)
	}
}
