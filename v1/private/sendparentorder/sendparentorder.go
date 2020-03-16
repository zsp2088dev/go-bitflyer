package sendparentorder

import (
	"github.com/google/go-querystring/query"
	"github.com/json-iterator/go"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Request struct {
	OrderMethod    OrderMethod `json:"order_method,omitempty"`
	MinuteToExpire int         `json:"minute_to_expire,omitempty"`
	TimeInForce    TimeInForce `json:"time_in_force,omitempty"`
	Parameters     []Param     `json:"parameters"`
}

type Response struct {
	ParentOrderAcceptanceId string `json:"parent_order_acceptance_id"`
}

type Param struct {
	ProductCode   types.ProductCode `json:"product_code"`
	ConditionType string            `json:"condition_type"`
	Side          types.Side        `json:"side"`
	Price         float64           `json:"price"`
	TriggerPrice  float64           `json:"trigger_price"`
	Size          float64           `json:"size"`
	Offset        int               `json:"offset,omitempty"`
}

type OrderMethod string

const (
	Simple OrderMethod = "SIMPLE"
	IFD    OrderMethod = "IFD"
	OCO    OrderMethod = "OCO"
	IFDOCO OrderMethod = "IFDOCO"
)

type TimeInForce string

const (
	GTC TimeInForce = "GTC"
	IOC TimeInForce = "IOC"
	FOK TimeInForce = "FOK"
)

func (req *Request) Endpoint() string {
	return "/v1/me/sendparentorder"
}

func (req *Request) Method() string {
	return http.MethodPost
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	b, err := json.Marshal(*req)
	if err != nil {
		return nil
	}
	return b
}
