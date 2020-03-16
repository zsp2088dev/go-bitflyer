package sendchildorder

import (
	"github.com/json-iterator/go"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Request struct {
	ProductCode    types.ProductCode    `json:"product_code"`
	ChildOrderType types.ChildOrderType `json:"child_order_type"`
	Side           types.Side           `json:"side"`
	Price          float64              `json:"price"`
	Size           float64              `json:"size"`
	MinuteToExpire int                  `json:"minute_to_expire,omitempty"`
	TimeInForce    string               `json:"time_in_force,omitempty"`
}

type Response struct {
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/sendchildorder"
}

func (req *Request) Method() string {
	return http.MethodPost
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	b, err := json.Marshal(*req)
	if err != nil {
		return nil
	}
	return b
}
