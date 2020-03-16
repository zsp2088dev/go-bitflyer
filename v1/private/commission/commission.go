package commission

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode types.ProductCode `url:"product_code"`
}

type Response struct {
	CommissionRate string `json:"commission_rate"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/gettradingcommission"
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
