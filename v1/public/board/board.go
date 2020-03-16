package board

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Book struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Response struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Book  `json:"bids"`
	Asks     []Book  `json:"asks"`
}

type Request struct {
	ProductCode types.ProductCode `url:"product_code"`
}

func (req *Request) Endpoint() string {
	return "/v1/getboard"
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
