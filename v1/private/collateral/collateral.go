package collateral

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type Request struct{}

type Response struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPNL   float64 `json:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getcollateral"
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
