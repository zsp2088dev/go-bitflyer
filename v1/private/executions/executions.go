package executions

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/time"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode            types.ProductCode `url:"product_code"`
	Count                  int               `url:"count,omitempty"`
	Before                 int               `url:"before,omitempty"`
	After                  int               `url:"after,omitempty"`
	ChildOrderAcceptanceID string            `url:"child_order_acceptance_id,omitempty"`
}

type Response []Execution

type Execution struct {
	ID                     int               `json:"id"`
	ChildOrderID           string            `json:"child_order_id"`
	Side                   types.Side        `json:"side"`
	Price                  float64           `json:"price"`
	Size                   float64           `json:"size"`
	Commission             int               `json:"commission"`
	ExecDate               time.BitFlyerTime `json:"exec_date"`
	ChildOrderAcceptanceID string            `json:"child_order_acceptance_id"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getexecutions"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
