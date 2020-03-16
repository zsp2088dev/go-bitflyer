package addresses

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct{}

type Response []Address

type Address struct {
	Type         string             `json:"type"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Address      string             `json:"address"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getaddresses"
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
