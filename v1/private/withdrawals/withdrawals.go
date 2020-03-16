package withdrawals

import (
	"github.com/google/go-querystring/query"
	"github.com/json-iterator/go"
	"net/http"
)

type Request struct {
	Count  int `url:"count,omitempty"`
	Before int `url:"before,omitempty"`
	After  int `url:"after,omitempty"`
}

type Response []Withdrawal

type Withdrawal struct {
	ID           int    `json:"id"`
	OrderID      string `json:"order_id"`
	CurrencyCode string `json:"currency_code"`
	Amount       int    `json:"amount"`
	Status       Status `json:"status"`
	EventDate    string `json:"event_date"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Status string

const (
	Pending   Status = "PENDING"
	Completed Status = "COMPLETED"
)

func (req *Request) Endpoint() string {
	return "/v1/me/getwithdrawals"
}

func (req *Request) Method() string {
	return http.MethodPost
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	b, err := json.Marshal(*req)
	if err != nil {
		return nil
	}
	return b
}
