package parentorder

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/time"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ParentOrderAcceptanceID string `url:"parent_order_acceptance_id"`
}

type Response struct {
	ID                      int               `json:"id"`
	ParentOrderID           string            `json:"parent_order_id"`
	OrderMethod             string            `json:"order_method"`
	ExpireDate              time.BitFlyerTime `json:"expire_date"`
	Parameters              []Parameter       `json:"parameters"`
	ParentOrderAcceptanceID string            `json:"parent_order_acceptance_id"`
}

type Parameter struct {
	ProductCode   types.ProductCode `json:"product_code"`
	ConditionType string            `json:"condition_type"`
	Side          types.Side        `json:"side"`
	Price         float64           `json:"price"`
	Size          float64           `json:"size"`
	TriggerPrice  int               `json:"trigger_price"`
	Offset        int               `json:"offset"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getparentorder"
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
