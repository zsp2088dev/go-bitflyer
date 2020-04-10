package realtime

import (
	"github.com/zsp2088dev/go-bitflyer/v1/public/board"
	"github.com/zsp2088dev/go-bitflyer/v1/public/executions"
	"github.com/zsp2088dev/go-bitflyer/v1/public/ticker"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
)

type Response struct {
	Event Event
	Error error

	Board             board.Response
	BoardSnapshot     board.Response
	Ticker            ticker.Response
	Executions        executions.Response
	ChildOrderEvents  []ChildOrderEvent
	ParentOrderEvents []ParentOrderEvent
}

type TickerResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string          `json:"channel"`
		Message ticker.Response `json:"message"`
	} `json:"params"`
}

type ExecutionsResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string              `json:"channel"`
		Message executions.Response `json:"message"`
	} `json:"params"`
}

type BoardResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string         `json:"channel"`
		Message board.Response `json:"message"`
	} `json:"params"`
}

type ChildOrderEventsResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string            `json:"channel"`
		Message []ChildOrderEvent `json:"message"`
	} `json:"params"`
}

type EventType string

const (
	Order        EventType = "ORDER"
	OrderFailed  EventType = "ORDER_FAILED"
	Cancel       EventType = "CANCEL"
	CancelFailed EventType = "CANCEL_FAILED"
	Execution    EventType = "EXECUTION"
	Expire       EventType = "EXPIRE"
)

type ChildOrderEvent struct {
	ProductCode            types.ProductCode    `json:"product_code"`
	ChildOrderID           string               `json:"child_order_id"`
	ChildOrderAcceptanceID string               `json:"child_order_acceptance_id"`
	EventDate              string               `json:"event_date"`
	EventType              EventType            `json:"event_type"`
	ChildOrderType         types.ChildOrderType `json:"child_order_type"`
	ExpireDate             string               `json:"expire_date"`
	Reason                 string               `json:"reason"`
	ExecID                 int                  `json:"exec_id"`
	Side                   types.Side           `json:"side"`
	Price                  float64              `json:"price"`
	Size                   float64              `json:"size"`
	Commission             float64              `json:"commission"`
	SFD                    float64              `json:"sfd"`
}

type ParentOrderEventsResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string             `json:"channel"`
		Message []ParentOrderEvent `json:"message"`
	} `json:"params"`
}

type ParentOrderEvent struct {
	ProductCode             types.ProductCode    `json:"product_code"`
	ParentOrderID           string               `json:"parent_order_id"`
	ParentOrderAcceptanceID string               `json:"parent_order_acceptance_id"`
	EventDate               string               `json:"event_date"`
	EventType               EventType            `json:"event_type"`
	ParentOrderType         string               `json:"parent_order_type"`
	Reason                  string               `json:"reason"`
	ChildOrderType          types.ChildOrderType `json:"child_order_type"`
	ParameterIndex          int                  `json:"parameter_index"`
	ChildOrderAcceptanceID  string               `json:"child_order_acceptance_id"`
	Side                    types.Side           `json:"side"`
	Price                   float64              `json:"price"`
	Size                    float64              `json:"size"`
	ExpireDate              string               `json:"expire_date"`
}
