package coinouts

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

type Response []CoinOut

type CoinOut struct {
	ID            int                `json:"id"`
	OrderID       string             `json:"order_id"`
	CurrencyCode  types.CurrencyCode `json:"currency_code"`
	Amount        float64            `json:"amount"`
	Address       string             `json:"address"`
	TxHash        string             `json:"tx_hash"`
	Fee           float64            `json:"fee"`
	AdditionalFee float64            `json:"additional_fee"`
	Status        Status             `json:"status"`
	EventDate     time.BitFlyerTime  `json:"event_date"`
}

type Status string

const (
	PENDING   Status = "PENDING"
	COMPLETED Status = "COMPLETED"
)

func (req *Request) Endpoint() string {
	return "/v1/me/getcoinouts"
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
