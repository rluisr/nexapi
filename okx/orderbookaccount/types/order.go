package types

import okxutils "github.com/rluisr/nexapi/okx/utils"

// PlaceOrderParam represents the structure for order request
type PlaceOrderParam struct {
	InstId         string          `json:"instId"`                   // Instrument ID, e.g. BTC-USD-190927-5000-C
	TdMode         string          `json:"tdMode"`                   // Trade mode (Margin mode: cross, isolated; Non-Margin mode: cash; spot_isolated: only applicable to SPOT lead trading, tdMode should be spot_isolated for SPOT lead trading.)
	Ccy            string          `json:"ccy,omitempty"`            // Margin currency (Only applicable to cross MARGIN orders in Single-currency margin.)
	ClOrdId        string          `json:"clOrdId,omitempty"`        // Client Order ID as assigned by the client (A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.)
	Tag            string          `json:"tag,omitempty"`            // Order tag (A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters.)
	Side           string          `json:"side"`                     // Order side (buy, sell)
	PosSide        string          `json:"posSide,omitempty"`        // Position side (The default is net in the net mode; required in the long/short mode, and can only be long or short; Only applicable to FUTURES/SWAP.)
	OrdType        string          `json:"ordType"`                  // Order type (market, limit, post_only, fok, ioc, optimal_limit_ioc, mmp, mmp_and_post_only)
	Sz             string          `json:"sz"`                       // Quantity to buy or sell
	Px             string          `json:"px,omitempty"`             // Order price (Only applicable to limit,post_only,fok,ioc,mmp,mmp_and_post_only order; When placing an option order, one of px/pxUsd/pxVol must be filled in, and only one can be filled in)
	PxUsd          string          `json:"pxUsd,omitempty"`          // Place options orders in USD (Only applicable to options; When placing an option order, one of px/pxUsd/pxVol must be filled in, and only one can be filled in)
	PxVol          string          `json:"pxVol,omitempty"`          // Place options orders based on implied volatility, where 1 represents 100% (Only applicable to options; When placing an option order, one of px/pxUsd/pxVol must be filled in, and only one can be filled in)
	ReduceOnly     bool            `json:"reduceOnly,omitempty"`     // Whether orders can only reduce in position size (Valid options: true or false; The default value is false; Only applicable to MARGIN orders, and FUTURES/SWAP orders in net mode; Only applicable to Single-currency margin and Multi-currency margin)
	TgtCcy         string          `json:"tgtCcy,omitempty"`         // Target currency (base_ccy: Base currency, quote_ccy: Quote currency; Only applicable to SPOT Market Orders; Default is quote_ccy for buy, base_ccy for sell)
	BanAmend       bool            `json:"banAmend,omitempty"`       // Whether to disallow the system from amending the size of the SPOT Market Order (Valid options: true or false; The default value is false; If true, system will not amend and reject the market order if user does not have sufficient funds; Only applicable to SPOT Market Orders)
	QuickMgnType   string          `json:"quickMgnType,omitempty"`   // Quick Margin type (Only applicable to Quick Margin Mode of isolated margin; manual, auto_borrow, auto_repay; The default value is manual)
	StpId          string          `json:"stpId,omitempty"`          // Self trade prevention ID (Numerical integers defined by user in the range of 1<= x<= 999999999)
	StpMode        string          `json:"stpMode,omitempty"`        // Self trade prevention mode (Default to cancel maker; cancel_maker,cancel_taker, cancel_both; Cancel both does not support FOK.)
	AttachAlgoOrds []AttachAlgoOrd `json:"attachAlgoOrds,omitempty"` // TP/SL information attached when placing order
}

// AttachAlgoOrd represents the structure for attaching algorithmic orders
type AttachAlgoOrd struct {
	AttachAlgoClOrdId    string `json:"attachAlgoClOrdId,omitempty"`    // Client-supplied Algo ID when placing order attaching TP/SL (A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters; It will be posted to algoClOrdId when placing TP/SL order once the general order is filled completely.)
	TpTriggerPx          string `json:"tpTriggerPx,omitempty"`          // Take-profit trigger price (If you fill in this parameter, you should fill in the take-profit order price as well.)
	TpOrdPx              string `json:"tpOrdPx,omitempty"`              // Take-profit order price (If you fill in this parameter, you should fill in the take-profit trigger price as well; If the price is -1, take-profit will be executed at the market price.)
	SlTriggerPx          string `json:"slTriggerPx,omitempty"`          // Stop-loss trigger price (If you fill in this parameter, you should fill in the stop-loss order price.)
	SlOrdPx              string `json:"slOrdPx,omitempty"`              // Stop-loss order price (If you fill in this parameter, you should fill in the stop-loss trigger price; If the price is -1, stop-loss will be executed at the market price.)
	TpTriggerPxType      string `json:"tpTriggerPxType,omitempty"`      // Take-profit trigger price type (last, index, mark; The Default is last)
	SlTriggerPxType      string `json:"slTriggerPxType,omitempty"`      // Stop-loss trigger price type (last, index, mark; The Default is last)
	Sz                   string `json:"sz,omitempty"`                   // Size (Only applicable to TP order of split TPs, and it is required for TP order of split TPs)
	AmendPxOnTriggerType string `json:"amendPxOnTriggerType,omitempty"` // Whether to enable Cost-price SL (Only applicable to SL order of split TPs; Whether slTriggerPx will move to avgPx when the first TP order is triggered; 0: disable, the default value; 1. Enable “Cost-price SL”)
}

type PlaceOrderResp struct {
	okxutils.Response
	Data []struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}

type GetOrderParam struct {
	InstId string `url:"instId"`
	OrdId  string `url:"ordId"`
}

type GetOrderResp struct {
	okxutils.Response
	Data []struct {
		InstType           string `json:"instType"`
		InstID             string `json:"instId"`
		Ccy                string `json:"ccy"`
		OrdID              string `json:"ordId"`
		ClOrdID            string `json:"clOrdId"`
		Tag                string `json:"tag"`
		Px                 string `json:"px"`
		PxUsd              string `json:"pxUsd"`
		PxVol              string `json:"pxVol"`
		PxType             string `json:"pxType"`
		Sz                 string `json:"sz"`
		Pnl                string `json:"pnl"`
		OrdType            string `json:"ordType"`
		Side               string `json:"side"`
		PosSide            string `json:"posSide"`
		TdMode             string `json:"tdMode"`
		AccFillSz          string `json:"accFillSz"`
		FillPx             string `json:"fillPx"`
		TradeID            string `json:"tradeId"`
		FillSz             string `json:"fillSz"`
		FillTime           string `json:"fillTime"`
		State              string `json:"state"`
		AvgPx              string `json:"avgPx"`
		Lever              string `json:"lever"`
		AttachAlgoClOrdID  string `json:"attachAlgoClOrdId"`
		TpTriggerPx        string `json:"tpTriggerPx"`
		TpTriggerPxType    string `json:"tpTriggerPxType"`
		TpOrdPx            string `json:"tpOrdPx"`
		SlTriggerPx        string `json:"slTriggerPx"`
		SlTriggerPxType    string `json:"slTriggerPxType"`
		SlOrdPx            string `json:"slOrdPx"`
		AttachAlgoOrds     []any  `json:"attachAlgoOrds"`
		StpID              string `json:"stpId"`
		StpMode            string `json:"stpMode"`
		FeeCcy             string `json:"feeCcy"`
		Fee                string `json:"fee"`
		RebateCcy          string `json:"rebateCcy"`
		Rebate             string `json:"rebate"`
		TgtCcy             string `json:"tgtCcy"`
		Category           string `json:"category"`
		ReduceOnly         string `json:"reduceOnly"`
		CancelSource       string `json:"cancelSource"`
		CancelSourceReason string `json:"cancelSourceReason"`
		QuickMgnType       string `json:"quickMgnType"`
		AlgoClOrdID        string `json:"algoClOrdId"`
		AlgoID             string `json:"algoId"`
		UTime              string `json:"uTime"`
		CTime              string `json:"cTime"`
	} `json:"data"`
}
