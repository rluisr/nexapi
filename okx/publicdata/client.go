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

package publicdata

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/rluisr/nexapi/okx/publicdata/types"
	okxutils "github.com/rluisr/nexapi/okx/utils"
	"github.com/rluisr/nexapi/utils"
)

type PublicDataClient struct {
	*okxutils.OKXRestClient

	// validate struct fields
	validate *validator.Validate
}

func NewPublicDataClient(cfg *okxutils.OKXRestClientCfg) (*PublicDataClient, error) {
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
		IsDemo:     cfg.IsDemo,
	})
	if err != nil {
		return nil, err
	}

	return &PublicDataClient{
		OKXRestClient: cli,
		validate:      validator,
	}, nil
}

func (p *PublicDataClient) GetInstruments(ctx context.Context, param types.GetInstrumentsParam) (*types.Instruments, error) {
	req := utils.HTTPRequest{
		Debug:   p.GetDebug(),
		BaseURL: p.GetBaseURL(),
		Path:    "/api/v5/public/instruments",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := p.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Instruments
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (p *PublicDataClient) GetMarketTickers(ctx context.Context, param types.GetMarketTickersParam) (*types.GetMarketTickersResp, error) {
	req := utils.HTTPRequest{
		Debug:   p.GetDebug(),
		BaseURL: p.GetBaseURL(),
		Path:    "/api/v5/market/tickers",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := p.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.GetMarketTickersResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (p *PublicDataClient) GetIndexTickers(ctx context.Context, param types.GetIndexTickersParam) (*types.GetIndexTickersResp, error) {
	req := utils.HTTPRequest{
		Debug:   p.GetDebug(),
		BaseURL: p.GetBaseURL(),
		Path:    "/api/v5/market/index-tickers",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := p.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.GetIndexTickersResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
