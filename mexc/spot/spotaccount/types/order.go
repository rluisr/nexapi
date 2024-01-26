package types

import "github.com/rluisr/nexapi/mexc/utils"

type CreateOrderParam struct {
	Symbol        string   `url:"symbol"`
	Side          string   `url:"side"`                    // ENUM: Order Side
	Type          string   `url:"type"`                    // ENUM: Order Type
	Quantity      *float64 `url:"quantity,omitempty"`      // DECIMAL
	QuoteOrderQty *float64 `url:"quoteOrderQty,omitempty"` // DECIMAL
	Price         *float64 `url:"price,omitempty"`         // DECIMAL
}

type CreateOrderParams struct {
	CreateOrderParam
	utils.DefaultParam
}

type CreateOrderResp struct {
	Symbol       string  `json:"symbol"`
	OrderId      int64   `json:"orderId"`
	OrderListId  int64   `json:"orderListId"`
	Price        float64 `json:"price"`
	OrigQty      float64 `json:"origQty"`
	Type         string  `json:"type"`
	Side         string  `json:"side"`
	TransactTime int64   `json:"transactTime"`
}

type Order struct {
	Symbol              string  `json:"symbol"`
	OrigClientOrderId   string  `json:"origClientOrderId"`
	OrderId             int64   `json:"orderId"`
	ClientOrderId       string  `json:"clientOrderId"`
	Price               float64 `json:"price"`
	OrigQty             float64 `json:"origQty"`
	ExecutedQty         float64 `json:"executedQty"`
	CummulativeQuoteQty float64 `json:"cummulativeQuoteQty"`
	Status              string  `json:"status"`
	TimeInForce         string  `json:"timeInForce"`
	Type                string  `json:"type"`
	Side                string  `json:"side"`
	StopPrice           float64 `json:"stopPrice"`
	Time                int64   `json:"time"`
	UpdateTime          int64   `json:"updateTime"`
	IsWorking           bool    `json:"isWorking"`
	OrigQuoteOrderQty   float64 `json:"origQuoteOrderQty"`
}
