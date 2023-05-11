package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/linstohu/nexapi/woox/api/types"
)

func (w *WooXClient) SendOrder(ctx context.Context, params types.SendOrderReq) (*types.SendOrderResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/order"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodPost,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SendOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) CancelOrder(ctx context.Context, params types.CancelOrderParam) (*types.CancelOrderResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/order"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodDelete,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) CancelOrderByClientOrderID(ctx context.Context, params types.CancelOrderByClientOrderIDParam) (*types.CancelOrderResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/client/order"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodDelete,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) CancelOrders(ctx context.Context, params types.CancelOrdersParam) (*types.CancelOrderResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/orders"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodDelete,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetOrder(ctx context.Context, orderID string) (*types.CancelOrderResp, error) {
	if orderID == "" {
		return nil, fmt.Errorf("oid must be given by api [/v1/order/:oid]")
	}

	path := fmt.Sprintf("%s%s", "/v1/order/", orderID)

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetOrderByClientOrderID(ctx context.Context, clientOrderID string) (*types.CancelOrderResp, error) {
	if clientOrderID == "" {
		return nil, fmt.Errorf("client_order_id must be given by api [/v1/client/order/:client_order_id]")
	}

	path := fmt.Sprintf("%s%s", "/v1/client/order/", clientOrderID)

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetOrders(ctx context.Context, params types.GetOrdersParam) (*types.GetOrders, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/orders"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetOrders
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetTrade(ctx context.Context, tradeID string) (*types.GetTrade, error) {
	if tradeID == "" {
		return nil, fmt.Errorf("tid must be given by api [/v1/client/trade/:tid]")
	}

	path := fmt.Sprintf("%s%s", "/v1/client/trade/", tradeID)

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetTradeHistory(ctx context.Context, params types.GetTradeHistoryParam) (*types.GetTradeHistory, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/client/trades"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetTradeHistory
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetBalances(ctx context.Context) (*types.Balance, error) {
	path := "/v3/balances"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV3APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Balance
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetAccountInfo(ctx context.Context) (*types.GetAccountInfo, error) {
	path := "/v3/accountinfo"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV3APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetAccountInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetAssetHisotry(ctx context.Context, params types.GetAssetHisotryParam) (*types.AssetHisotryResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/asset/history"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Query:  params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.AssetHisotryResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetSubAccounts(ctx context.Context) (*types.SubAccounts, error) {
	path := "/v1/sub_account/all"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SubAccounts
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) TransferAsset(ctx context.Context, params types.TransferAssetParam) (*types.TransferAssetResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/asset/main_sub_transfer"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodPost,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.TransferAssetResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) UpdateLeverageSetting(ctx context.Context, params types.UpdateLeverageSettingParam) (*types.Response, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/client/leverage"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodPost,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Response
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetIPRestriction(ctx context.Context) (*types.IPRestriction, error) {
	path := "/v1/sub_account/ip_restriction"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.IPRestriction
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetOnePositionInfo(ctx context.Context, symbol string) (*types.GetOnePositionInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/position/:symbol]")
	}

	path := fmt.Sprintf("%s%s", "/v1/position/", symbol)

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetOnePositionInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetAllPositionInfo(ctx context.Context) (*types.GetAllPositionInfo, error) {
	path := "/v3/positions"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV3APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetAllPositionInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
