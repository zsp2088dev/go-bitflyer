package cancelparentorder

import (
	"github.com/google/go-querystring/query"
	"github.com/json-iterator/go"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Request struct {
	ProductCode             types.ProductCode `json:"product_code"`
	ParentOrderAcceptanceId string            `json:"parent_order_acceptance_id"`
}

type Response struct{}

func (req *Request) Endpoint() string {
	return "/v1/me/cancelparentorder"
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
