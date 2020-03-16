package positions

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode types.ProductCode `url:"product_code"`
}

type Response []Position

type Position struct {
	ProductCode         types.ProductCode `json:"product_code"`
	Side                types.Side        `json:"side"`
	Price               float64           `json:"price"`
	Size                float64           `json:"size"`
	Commission          float64           `json:"commission"`
	SwapPointAccumulate float64           `json:"swap_point_accumulate"`
	RequireCollateral   float64           `json:"require_collateral"`
	OpenDate            string            `json:"open_date"`
	Leverage            float64           `json:"leverage"`
	Pnl                 float64           `json:"pnl"`
	Sfd                 float64           `json:"sfd"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getpositions"
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
