package executions

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode types.ProductCode `url:"product_code"`
}

type Response []Execution

type Execution struct {
	ID                         int        `json:"id"`
	Side                       types.Side `json:"side"`
	Price                      float64    `json:"price"`
	Size                       float64    `json:"size"`
	ExecDate                   string     `json:"exec_date"`
	BuyChildOrderAcceptanceID  string     `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string     `json:"sell_child_order_acceptance_id"`
}

func (req *Request) Endpoint() string {
	return "/v1/getexecutions"
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