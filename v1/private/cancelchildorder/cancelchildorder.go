package cancelchildorder

import (
	"github.com/json-iterator/go"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Request struct {
	ProductCode            types.ProductCode `json:"product_code"`
	ChildOrderAcceptanceId string            `json:"child_order_acceptance_id"`
}

type Response struct{}

func (req *Request) Endpoint() string {
	return "/v1/me/cancelchildorder"
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
