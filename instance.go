package goyacuna

import (
	"strconv"
	"fmt"
	"github.com/bndr/gopencils"
	"github.com/fatih/structs"
	"errors"
	"time"
	"bytes"
	"net/url"
)

type Instance struct {
	id string
	secret string
	api *gopencils.Resource
}


func (r *Instance) SetId(id string) {
	r.id = id
}

func (r *Instance) SetSecret(secret string) {
	r.secret = secret
}

var ErrUnexpectedType = errors.New("unexpected type in response")
//var ErrClientError = errors.New("client error (4xx)")
//var ErrServerError = errors.New("server error (5xx)")


//	Deal Count
//	http://docs.yacuna.com/api/#api-Deal-Deal_count
func (r *Instance) GetDealCount(req *DealCountRequest) (*CountResponse, error) {

	uriPath := "deal/count"
	respObj := &CountResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*CountResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Deal Get
// http://docs.yacuna.com/api/#api-Deal-Deal_get
func (r *Instance) GetDeal(dealId string) (*GetDealResponse, error) {

	uriPath := "deal/get/" + dealId
	respObj := &GetDealResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*GetDealResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

//	Deal List
//	http://docs.yacuna.com/api/#api-Deal-Deal_list
func (r *Instance) GetDealList(req *DealListRequest) (*DealListResponse, error) {

	uriPath := "deal/list"
	respObj := &DealListResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*DealListResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}


//	Market Count
//	http://docs.yacuna.com/api/#api-Market-Market_count
func (r *Instance) GetMarketCount(req *MarketCountRequest) (*CountResponse, error) {

	uriPath := "market/count"
	respObj := &CountResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*CountResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

//	Market List
//	http://docs.yacuna.com/api/#api-Market-Market_list
func (r *Instance) GetMarketList(req *MarketListRequest) (*MarketListResponse, error) {

	uriPath := "market/list"
	respObj := &MarketListResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)

	re, err := doGet(res)

	ret, ok := re.Response.(*MarketListResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}


// Get MarketDepth
// http://docs.yacuna.com/api/#api-Market-MarketDepth_get
func (r *Instance) GetMarketDepth(currency1 string, currency2 string) (*GetMarketDepthResponse, error) {

	uriPath := fmt.Sprintf("marketdepth/get/%s/%s", currency1, currency2)
	respObj := &GetMarketDepthResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*GetMarketDepthResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Get OrderBook
// http://docs.yacuna.com/api/#api-Market-OrderBook_get
func (r *Instance) GetOrderBook(currency1 string, currency2 string) (*GetOrderBookResponse, error) {

	uriPath := fmt.Sprintf("orderbook/get/%s/%s", currency1, currency2)
	respObj := &GetOrderBookResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*GetOrderBookResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Cancel Order
// http://docs.yacuna.com/api/#api-Order-Order_cancel
func (r *Instance) CancelOrder(orderId string) (*OrderResponse, error) {

	uriPath := "order/cancel/" + orderId
	respObj := &OrderResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setPostAuthentication(res, uriPath)
	re, err := doPost(res)

	ret, ok := re.Response.(*OrderResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Confirm Order
// http://docs.yacuna.com/api/#api-Order-Order_confirm
func (r *Instance) ConfirmOrder(orderId string) (*OrderResponse, error) {

	uriPath := "order/confirm/" + orderId
	respObj := &OrderResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setPostAuthentication(res, uriPath)
	re, err := doPost(res)

	ret, ok := re.Response.(*OrderResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

//	Order Count
//	http://docs.yacuna.com/api/#api-Order-Order_count
func (r *Instance) GetOrderCount(req *OrderCountRequest) (*CountResponse, error) {

	uriPath := "order/count"
	respObj := &CountResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*CountResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Create Order
// http://docs.yacuna.com/api/#api-Order-Order_create
func (r *Instance) CreateOrder(currency1 string, currency2 string, req *CreateOrderRequest) (*OrderResponse, error) {

	uriPath := fmt.Sprintf("order/create/%s/%s", currency1, currency2)
	respObj := &OrderResponse{}
	res := r.api.Res(uriPath, respObj)

	payload := toPayload(toStringMap(req))
	r.setPostAuthentication(res, uriPath, payload)
	re, err := doPost(res, payload)

	ret, ok := re.Response.(*OrderResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Get Order
// http://docs.yacuna.com/api/#api-Order-Order_Get
func (r *Instance) GetOrder(orderId string) (*OrderResponse, error) {

	uriPath := "order/get/" + orderId
	respObj := &OrderResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*OrderResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Get Order by ext ref id
// http://docs.yacuna.com/api/#api-Order-Order_get_by_external_reference_Id
func (r *Instance) GetOrderByExtRefId(walletAccountId string, externalReferenceId string) (*OrderResponse, error) {

	uriPath := "order/getByExternalReferenceId/" + walletAccountId
	respObj := &OrderResponse{}
	res := r.api.Res(uriPath, respObj)

	req := &OrderByExtRefIdRequest{externalReferenceId, walletAccountId}

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*OrderResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}


//	Order List
//	http://docs.yacuna.com/api/#api-Order-Order_list
func (r *Instance) GetOrderList(req *OrderListRequest) (*OrderListResponse, error) {

	uriPath := "order/list"
	respObj := &OrderListResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*OrderListResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}


// Get OrderBook Chart
// http://docs.yacuna.com/api/#api-Public_Charts-Orderbook
func (r *Instance) GetOrderBookChart(currency1 string, currency2 string) (*OrderBookChart, error) {

	uriPath := fmt.Sprintf("charts/orderbook/%s/%s", currency1, currency2)
	respObj := &OrderBookChart{}
	res := r.api.Res(uriPath, respObj)

	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*OrderBookChart)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

// Get Trades Chart
// http://docs.yacuna.com/api/#api-Public_Charts-Trades
func (r *Instance) GetTradesChart(currency1 string, currency2 string, since int) (*TradesChart, error) {

	uriPath := fmt.Sprintf("charts/trades/%s/%s", currency1, currency2)
	respObj := &TradesChart{}

	query := map[string]string{"since": strconv.Itoa(since)}
	res := r.api.Res(uriPath, respObj)
	res.SetQuery(query)
	r.setGetAuthentication(res, uriPath)

	re, err := doGet(res)

	ret, ok := re.Response.(*TradesChart)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}


//	Get Wallet
//	http://docs.yacuna.com/api/#api-Wallet-Wallet_get
func (r *Instance) GetWallet(req *GetWalletRequest) (*GetWalletResponse, error) {

	uriPath := "wallet/get"
	respObj := &GetWalletResponse{}
	res := r.api.Res(uriPath, respObj)

	res.SetQuery(*toStringMap(req))
	r.setGetAuthentication(res, uriPath)
	re, err := doGet(res)

	ret, ok := re.Response.(*GetWalletResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, err

}

func doGet(res *gopencils.Resource) (*gopencils.Resource, error) {
	re, err := res.Get()
	return checkRequestResult(re, err)
}

func toPayload(payload *map[string]string) string {
	vals := make(url.Values)
	for k, v := range *payload {
		vals.Set(k, v)
	}
	return vals.Encode()
}

func doPost(res *gopencils.Resource, payload ...string) (*gopencils.Resource, error) {
	if len(payload) > 0 {
		res.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		res.Payload = bytes.NewBuffer([]byte(payload[0]))
	}
	re, err := res.Post()
	return checkRequestResult(re, err)
}

func checkRequestResult(re *gopencils.Resource, err error) (*gopencils.Resource, error) {
	if err != nil {
		return re, err
	}

	if re.Raw.StatusCode >= 500 {
		return re, errors.New("server error " + strconv.Itoa(re.Raw.StatusCode))
	}
	if re.Raw.StatusCode >= 400 {
		return re, errors.New("client error " + strconv.Itoa(re.Raw.StatusCode))
	}

	return re, err
}

func (r *Instance) setGetAuthentication(res *gopencils.Resource, uriPath string) {

	ti := &apiTokenInput{
		secret: r.secret,
		method: "GET",
		path: 	res.Api.BaseUrl.Path + "/" + uriPath,
		query: 	res.QueryValues.Encode(),
	}

	setAuthentication(res, r.id, ti)
}

func (r *Instance) setPostAuthentication(res *gopencils.Resource, uriPath string, payload ...string) {

	ti := &apiTokenInput{
		secret: r.secret,
		method: "POST",
		path: 	res.Api.BaseUrl.Path + "/" + uriPath,
		query: 	res.QueryValues.Encode(),
	}

	if len(payload) > 0 {
		ti.body = payload[0]
	}

	setAuthentication(res, r.id, ti)
}

func setAuthentication(res *gopencils.Resource, id string, ti *apiTokenInput) {
	res.SetHeader(H_ApiTokenId, id)
	res.SetHeader(H_ApiToken, ApiToken( ti ) )
	res.SetHeader(H_ApiTokenOTP, "")
}

func toStringMap(req interface{}) *map[string]string {

	rm := &map[string]string{}
	for key, value := range structs.Map(req) {
		formatVal(key, value, rm)
	}

	return rm
}

func formatVal(key string, value interface{}, rm *map[string]string) {
	strVal := ""
	switch t := value.(type) {
	default:
		strVal = fmt.Sprintf("%v", t)
	case time.Time:
		strVal = t.Format(time.RFC3339)
	case map[string]interface{}:
		for key2, value2 := range t {
			formatVal(key2, value2, rm)
		}
		return
	}
	(*rm)[key] = strVal
}

