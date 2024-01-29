/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package orderbookaccount

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/rluisr/nexapi/okx/orderbookaccount/types"
	okxutils "github.com/rluisr/nexapi/okx/utils"
	"github.com/rluisr/nexapi/utils"
)

type OrderBookAccountClient struct {
	*okxutils.OKXRestClient

	// validate struct fields
	validate *validator.Validate
}

type OrderBookAccountClientCfg struct {
	BaseURL    string `validate:"required"`
	HTTPClient *http.Client
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	Passphrase string `validate:"required"`
	Debug      bool
	// Logger
	Logger *slog.Logger
	IsDemo bool `validate:"-"`
}

func NewOrderBookAccountClient(cfg *OrderBookAccountClientCfg) (*OrderBookAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := okxutils.NewOKXRestClient(&okxutils.OKXRestClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		HTTPClient: cfg.HTTPClient,
		Key:        cfg.Key,
		Secret:     cfg.Secret,
		Passphrase: cfg.Passphrase,
		IsDemo:     cfg.IsDemo,
	})
	if err != nil {
		return nil, err
	}

	return &OrderBookAccountClient{
		OKXRestClient: cli,
		validate:      validator,
	}, nil
}

func (o *OrderBookAccountClient) PlaceOrder(ctx context.Context, param types.PlaceOrderParam) (*types.PlaceOrderResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/api/v5/trade/order",
		Method:  http.MethodPost,
		Body:    param,
	}

	headers, err := o.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.PlaceOrderResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
