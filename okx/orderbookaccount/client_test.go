package orderbookaccount

import (
	"context"
	"net/http"
	"testing"

	"github.com/rluisr/nexapi/okx/orderbookaccount/types"
	"github.com/rluisr/nexapi/okx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewOrderBookAccountClient(t *testing.T) *OrderBookAccountClient {
	cli, err := NewOrderBookAccountClient(&OrderBookAccountClientCfg{
		Debug:      true,
		IsDemo:     true,
		BaseURL:    utils.RestURL,
		Key:        "3b34b674-3cff-4c9d-87d0-75525536d322",
		Secret:     "AB4E9508E58BFD909B47A83A9A052B78",
		Passphrase: "Luis*o*8526",
		HTTPClient: &http.Client{},
	})

	if err != nil {
		t.Fatalf("Could not create okx private client, %s", err)
	}

	return cli
}

func TestPlaceOrder_Margin(t *testing.T) {
	cli := testNewOrderBookAccountClient(t)

	resp, err := cli.PlaceOrder(context.TODO(), types.PlaceOrderParam{
		InstId:  "BTC-USDT",
		TdMode:  utils.Isolated,
		Side:    utils.Buy,
		OrdType: utils.Market,
		Sz:      "10",
		AttachAlgoOrds: []types.AttachAlgoOrd{
			{
				TpTriggerPx:     "43000",
				TpOrdPx:         "-1",
				SlTriggerPx:     "40000",
				SlOrdPx:         "-1",
				TpTriggerPxType: utils.Index,
				SlTriggerPxType: utils.Index,
				Sz:              "10",
			},
		},
	})

	assert.Nil(t, err)

	if resp.Code != "0" {
		assert.FailNowf(t, "PlaceOrder", "%+v", resp)
	}
}
func TestPlaceOrder_SpotBuy(t *testing.T) {
	cli := testNewOrderBookAccountClient(t)

	resp, err := cli.PlaceOrder(context.TODO(), types.PlaceOrderParam{
		InstId:  "BTC-USDT",
		TdMode:  utils.Cash,
		Side:    utils.Buy,
		OrdType: utils.Market,
		Sz:      "10",
	})
	assert.Nil(t, err)

	if resp.Code != "0" {
		assert.FailNowf(t, "PlaceOrder", "%+v", resp)
	}
}

func TestPlaceOrder_SpotSell(t *testing.T) {
	cli := testNewOrderBookAccountClient(t)

	resp, err := cli.PlaceOrder(context.TODO(), types.PlaceOrderParam{
		InstId:  "BTC-USDT",
		TdMode:  utils.Cash,
		Side:    utils.Sell,
		OrdType: utils.Limit,
		Px:      "",
		Sz:      "10",
	})
	assert.Nil(t, err)

	if resp.Code != "0" {
		assert.FailNowf(t, "PlaceOrder", "%+v", resp)
	}
}
