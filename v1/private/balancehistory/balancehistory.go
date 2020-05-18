package balancehistory

import (
	"github.com/google/go-querystring/query"
	"github.com/zsp2088dev/go-bitflyer/v1/time"
	"github.com/zsp2088dev/go-bitflyer/v1/types"
	"net/http"
)

type Request struct {
	CurrencyCode types.CurrencyCode `url:"currency_code"`
	Count        int                `url:"count,omitempty"`
	Before       int                `url:"before,omitempty"`
	After        int                `url:"after,omitempty"`
}

type Response []BalanceHistory

type BalanceHistory struct {
	ID           int               `json:"id"`
	TradeDate    time.BitFlyerTime `json:"trade_date"`
	ProductCode  types.ProductCode `json:"product_code"`
	CurrencyCode string            `json:"currency_code"`
	TradeType    TradeType         `json:"trade_type"`
	Price        float64           `json:"price"`
	Amount       float64           `json:"amount"`
	Quantity     float64           `json:"quantity"`
	Commission   int               `json:"commission"`
	Balance      float64           `json:"balance"`
	OrderID      string            `json:"order_id"`
}

type TradeType string

const (
	Buy        TradeType = "BUY"         // 買い
	Sell       TradeType = "SELL"        // 売り
	Deposit    TradeType = "DEPOSIT"     // 入金・仮想通貨の預入
	Withdraw   TradeType = "WITHDRAW"    // 出金・仮想通貨の外部送付
	Fee        TradeType = "FEE"         // 手数料
	PostColl   TradeType = "POST_COL"    // 証拠金の預入
	CancelColl TradeType = "CANCEL_COLL" // 証拠金の引出
	Payment    TradeType = "PAYMENT"     //  ビットコイン決済による仮想通貨の移転
	Transfer   TradeType = "TRANSFER"    // その他の一般的な資金移動
)

func (req *Request) Endpoint() string {
	return "/v1/me/getbalancehistory"
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
