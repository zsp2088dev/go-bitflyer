package ticker

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`
}

type Response struct {
	ProductCode     types.ProductCode `json:"product_code"`
	Timestamp       string            `json:"timestamp"`
	TickID          int               `json:"tick_id"`
	BestBid         float64           `json:"best_bid"`
	BestAsk         float64           `json:"best_ask"`
	BestBidSize     float64           `json:"best_bid_size"`
	BestAskSize     float64           `json:"best_ask_size"`
	TotalBidDepth   float64           `json:"total_bid_depth"`
	TotalAskDepth   float64           `json:"total_ask_depth"`
	Ltp             float64           `json:"ltp"`
	Volume          float64           `json:"volume"`
	VolumeByProduct float64           `json:"volume_by_product"`
}

func (req *Request) Endpoint() string {
	return "/v1/ticker"
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
