package markets

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`
}

type Response []Market

type Market struct {
	ProductCode types.ProductCode `json:"product_code"`
	Alias       types.ProductCode `json:"alias"`
}

func (req *Request) Endpoint() string {
	return "/v1/getmarkets"
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
