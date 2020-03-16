package types

type CurrencyCode string

type ProductCode string

const (
	BTCJPY   ProductCode = "BTC_JPY"
	FXBTCJPY ProductCode = "FX_BTC_JPY"
	ETHJPY   ProductCode = "ETH_JPY"
	ETHBTC   ProductCode = "ETH_BTC"
)

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type ChildOrderType string

const (
	Market ChildOrderType = "MARKET"
	Limit  ChildOrderType = "LIMIT"
)
