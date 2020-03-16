package v1

import (
	"github.com/pkg/errors"
	api "github.com/zsp2088dev/go-bitflyer"
	"github.com/zsp2088dev/go-bitflyer/auth"
	"github.com/zsp2088dev/go-bitflyer/httpclient"
	"github.com/zsp2088dev/go-bitflyer/v1/private/addresses"
	"github.com/zsp2088dev/go-bitflyer/v1/private/balance"
	"github.com/zsp2088dev/go-bitflyer/v1/private/balancehistory"
	"github.com/zsp2088dev/go-bitflyer/v1/private/bankaccounts"
	"github.com/zsp2088dev/go-bitflyer/v1/private/cancelchildorder"
	"github.com/zsp2088dev/go-bitflyer/v1/private/cancels"
	"github.com/zsp2088dev/go-bitflyer/v1/private/childorders"
	"github.com/zsp2088dev/go-bitflyer/v1/private/coinins"
	"github.com/zsp2088dev/go-bitflyer/v1/private/coinouts"
	"github.com/zsp2088dev/go-bitflyer/v1/private/collateral"
	"github.com/zsp2088dev/go-bitflyer/v1/private/collateralhistory"
	"github.com/zsp2088dev/go-bitflyer/v1/private/commission"
	"github.com/zsp2088dev/go-bitflyer/v1/private/deposits"
	"github.com/zsp2088dev/go-bitflyer/v1/private/parentorder"
	"github.com/zsp2088dev/go-bitflyer/v1/private/parentorders"
	"github.com/zsp2088dev/go-bitflyer/v1/private/permissions"
	"github.com/zsp2088dev/go-bitflyer/v1/private/positions"
	"github.com/zsp2088dev/go-bitflyer/v1/private/sendchildorder"
	"github.com/zsp2088dev/go-bitflyer/v1/private/sendparentorder"
	"github.com/zsp2088dev/go-bitflyer/v1/private/withdrawals"
	"github.com/zsp2088dev/go-bitflyer/v1/public/board"
	"github.com/zsp2088dev/go-bitflyer/v1/public/boardstate"
	"github.com/zsp2088dev/go-bitflyer/v1/public/executions"
	e "github.com/zsp2088dev/go-bitflyer/v1/public/executions"
	"github.com/zsp2088dev/go-bitflyer/v1/public/health"
	"github.com/zsp2088dev/go-bitflyer/v1/public/markets"
	"github.com/zsp2088dev/go-bitflyer/v1/public/ticker"
	"net/http"
)

const BaseURL = "https://api.bitflyer.com"

type Client struct {
	BaseURL string
	Auth    *auth.Auth
}

func NewClient(auth *auth.Auth) *Client {
	return &Client{
		BaseURL: BaseURL,
		Auth:    auth,
	}
}

/*
	Public API
*/

// マーケットの一覧
func (c *Client) Markets(req *markets.Request) (*markets.Response, *http.Response, error) {
	res := new(markets.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 板情報
func (c *Client) Board(req *board.Request) (*board.Response, *http.Response, error) {
	res := new(board.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// Ticker
func (c *Client) Ticker(req *ticker.Request) (*ticker.Response, *http.Response, error) {
	res := new(ticker.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 約定履歴
func (c *Client) Executions(req *executions.Request) (*executions.Response, *http.Response, error) {
	res := new(executions.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 板の状態
func (c *Client) Health(req *health.Request) (*health.Response, *http.Response, error) {
	res := new(health.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 取引所の状態
func (c *Client) BoardState(req *boardstate.Request) (*boardstate.Response, *http.Response, error) {
	res := new(boardstate.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

/*
	Private API
*/

// APIキーの権限を取得
func (c *Client) Permissions(req *permissions.Request) (*permissions.Response, *http.Response, error) {
	res := new(permissions.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 資産残高を取得
func (c *Client) Balance(req *balance.Request) (*balance.Response, *http.Response, error) {
	res := new(balance.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 証拠金の状態を取得
func (c *Client) Collateral(req *collateral.Request) (*collateral.Response, *http.Response, error) {
	res := new(collateral.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 預入用アドレス取得
func (c *Client) Addresses(req *addresses.Request) (*addresses.Response, *http.Response, error) {
	res := new(addresses.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 仮想通貨預入履歴
func (c *Client) CoinIns(req *coinins.Request) (*coinins.Response, *http.Response, error) {
	res := new(coinins.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 仮想通貨送付履歴
func (c *Client) CoinOuts(req *coinouts.Request) (*coinouts.Response, *http.Response, error) {
	res := new(coinouts.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 銀行口座一覧取得
func (c *Client) BankAccounts(req *bankaccounts.Request) (*bankaccounts.Response, *http.Response, error) {
	res := new(bankaccounts.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 入金履歴
func (c *Client) Deposits(req *deposits.Request) (*deposits.Response, *http.Response, error) {
	res := new(deposits.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 出金履歴
func (c *Client) Withdrawals(req *withdrawals.Request) (*withdrawals.Response, *http.Response, error) {
	res := new(withdrawals.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 新規注文を出す
func (c *Client) SendChildOrder(req *sendchildorder.Request) (*sendchildorder.Response, *http.Response, error) {
	res := new(sendchildorder.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 注文をキャンセルする
func (c *Client) CancelChildOrder(req *cancelchildorder.Request) (*cancelchildorder.Response, *http.Response, error) {
	res := new(cancelchildorder.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 新規の親注文を出す（特殊注文）
func (c *Client) SendParentOrder(req *sendparentorder.Request) (*sendparentorder.Response, *http.Response, error) {
	res := new(sendparentorder.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 親注文をキャンセルする
func (c *Client) CancelParentOrder(req *cancelchildorder.Request) (*cancelchildorder.Response, *http.Response, error) {
	res := new(cancelchildorder.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// すべての注文をキャンセルする
func (c *Client) CancelAllChildOrders(req *cancels.Request) (*cancels.Response, *http.Response, error) {
	res := new(cancels.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 注文の一覧を取得
func (c *Client) ChildOrders(req *childorders.Request) (*childorders.Response, *http.Response, error) {
	res := new(childorders.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 親注文の一覧を取得
func (c *Client) ParentOrders(req *parentorders.Request) (*parentorders.Response, *http.Response, error) {
	res := new(parentorders.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 親注文の詳細を取得
func (c *Client) ParentOrder(req *parentorder.Request) (*parentorder.Response, *http.Response, error) {
	res := new(parentorder.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 約定の一覧を取得
func (c *Client) ExecutionsMe(req *e.Request) (*e.Response, *http.Response, error) {
	res := new(e.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 残高履歴を取得
func (c *Client) BalanceHistory(req *balancehistory.Request) (*balancehistory.Response, *http.Response, error) {
	res := new(balancehistory.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 建玉の一覧を取得
func (c *Client) Positions(req *positions.Request) (*positions.Response, *http.Response, error) {
	res := new(positions.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 証拠金の変動履歴を取得
func (c *Client) CollateralHistory(req *collateralhistory.Request) (*collateralhistory.Response, *http.Response, error) {
	res := new(collateralhistory.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// 取引手数料を取得
func (c *Client) Commission(req *commission.Request) (*commission.Response, *http.Response, error) {
	res := new(commission.Response)
	raw, err := c.do(req, res)
	return res, raw, err
}

// HttpClientを利用してbitFlyerへリクエスト
func (c *Client) do(req api.Request, result interface{}) (*http.Response, error) {
	var raw *http.Response
	var err error

	if c.Auth != nil {
		raw, err = httpclient.New().Auth(c.Auth).Request(NewAPI(c, req.Endpoint()), req, result)
	} else {
		raw, err = httpclient.New().Request(NewAPI(c, req.Endpoint()), req, result)
	}

	if err != nil {
		return nil, errors.Wrap(err, "send request")
	}

	return raw, err
}
