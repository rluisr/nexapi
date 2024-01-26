package types

import "github.com/rluisr/nexapi/mexc/utils"

type CreateOrderParam struct {
	Symbol        string   `json:"symbol"`
	Side          string   `json:"side"`                    // ENUM: Order Side
	Type          string   `json:"type"`                    // ENUM: Order Type
	Quantity      *float64 `json:"quantity,omitempty"`      // DECIMAL
	QuoteOrderQty *float64 `json:"quoteOrderQty,omitempty"` // DECIMAL
	Price         *float64 `json:"price,omitempty"`         // DECIMAL
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
