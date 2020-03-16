package coinins

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	Count  int `url:"count,omitempty"`
	Before int `url:"before,omitempty"`
	After  int `url:"after,omitempty"`
}

type Response []CoinIn

type CoinIn struct {
	ID           int                `json:"id"`
	OrderID      string             `json:"order_id"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Amount       float64            `json:"amount"`
	Address      string             `json:"address"`
	TxHash       string             `json:"tx_hash"`
	Status       Status             `json:"status"`
	EventDate    string             `json:"event_date"`
}

type Status string

const (
	Pending   Status = "PENDING"
	Completed Status = "COMPLETED"
)

func (req *Request) Endpoint() string {
	return "/v1/me/getcoinins"
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
