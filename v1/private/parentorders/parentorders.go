package parentorders

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode      types.ProductCode `url:"product_code"`
	ParentOrderState string            `url:"parent_order_state,omitempty"`
	Count            int               `url:"count,omitempty"`
	Before           int               `url:"before,omitempty"`
	After            int               `url:"after,omitempty"`
}

type Response []ParentOrder

type ParentOrder struct {
	ID                      int               `json:"id"`
	ParentOrderID           string            `json:"parent_order_id"`
	ProductCode             types.ProductCode `json:"product_code"`
	Side                    types.Side        `json:"side"`
	ParentOrderType         string            `json:"parent_order_type"`
	Price                   int               `json:"price"`
	AveragePrice            int               `json:"average_price"`
	Size                    float64           `json:"size"`
	ParentOrderState        string            `json:"parent_order_state"`
	ExpireDate              string            `json:"expire_date"`
	ParentOrderDate         string            `json:"parent_order_date"`
	ParentOrderAcceptanceID string            `json:"parent_order_acceptance_id"`
	OutstandingSize         int               `json:"outstanding_size"`
	CancelSize              int               `json:"cancel_size"`
	ExecutedSize            float64           `json:"executed_size"`
	TotalCommission         int               `json:"total_commission"`
}

type ParentOrderState string

const (
	ACTIVE    ParentOrderState = "ACTIVE"
	COMPLETED ParentOrderState = "COMPLETED"
	CANCELED  ParentOrderState = "CANCELED"
	EXPIRED   ParentOrderState = "EXPIRED"
	REJECTED  ParentOrderState = "REJECTED"
)

func (req *Request) Endpoint() string {
	return "/v1/me/getparentorders"
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
