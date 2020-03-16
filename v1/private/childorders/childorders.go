package childorders

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode            types.ProductCode `url:"product_code"`
	ChildOrderState        string            `url:"child_order_state,omitempty"`
	ChildOrderID           string            `url:"child_order_id,omitempty"`
	ChildOrderAcceptanceID string            `url:"child_order_acceptance_id,omitempty"`
	ParentOrderID          string            `url:"parent_order_id,omitempty"`
	Count                  int               `url:"count,omitempty"`
	Before                 int               `url:"before,omitempty"`
	After                  int               `url:"after,omitempty"`
}

type Response []ChildOrder

type ChildOrder struct {
	ID                     int                  `json:"id"`
	ChildOrderID           string               `json:"child_order_id"`
	ProductCode            types.ProductCode    `json:"product_code"`
	Side                   types.Side           `json:"side"`
	ChildOrderType         types.ChildOrderType `json:"child_order_type"`
	Price                  float64              `json:"price"`
	AveragePrice           float64              `json:"average_price"`
	Size                   float64              `json:"size"`
	ChildOrderState        string               `json:"child_order_state"`
	ExpireDate             string               `json:"expire_date"`
	ChildOrderDate         string               `json:"child_order_date"`
	ChildOrderAcceptanceID string               `json:"child_order_acceptance_id"`
	OutstandingSize        float64              `json:"outstanding_size"`
	CancelSize             float64              `json:"cancel_size"`
	ExecutedSize           float64              `json:"executed_size"`
	TotalCommission        float64              `json:"total_commission"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getchildorders"
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
