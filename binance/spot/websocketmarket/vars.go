package websocketmarket

var (
	SpotMarketStreamBaseURL = "wss://data-stream.binance.com"
	CombinedStreamRouter    = "/stream"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)