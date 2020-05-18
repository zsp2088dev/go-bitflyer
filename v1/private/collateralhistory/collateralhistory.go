package collateralhistory

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/time"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	Count  int `url:"count,omitempty"`
	Before int `url:"before,omitempty"`
	After  int `url:"after,omitempty"`
}

type Response []CollateralHistory

type CollateralHistory struct {
	ID           int                `json:"id"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Change       float64            `json:"change"`
	Amount       float64            `json:"amount"`
	ReasonCode   string             `json:"reason_code"`
	Date         time.BitFlyerTime  `json:"date"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getcollateralhistory"
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
