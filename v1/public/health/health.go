package health

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code"`
}

type Response struct {
	Status Status `json:"status"`
}

type Status string

const (
	Running      Status = "RUNNING"
	Closed       Status = "CLOSED"
	Starting     Status = "Starting"
	PreOpen      Status = "PREOPEN"
	CircuitBreak Status = "CIRCUIT BREAK"
	AwaitingSQ   Status = "AWAITING SQ"
	Matured      Status = "MATURED"
)

func (req *Request) Endpoint() string {
	return "/v1/gethealth"
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
