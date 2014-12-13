package goyacuna

import (
	"fmt"
	"github.com/bndr/gopencils"
	"github.com/fatih/structs"
	"errors"
)

const (
	API_PATH="/api/1"
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

//	Deal count
//	http://docs.yacuna.com/api/#api-Deal-Deal_count
type DealCountRequest struct {

	WalletAccountId		string			`json:"walletAccountId"`
	MarketId 			string			`json:"marketId"`
	OrderId 			string			`json:"orderId"`
	TradeDealType 		string			`json:"tradeDealType"`
	TradeDealStatus 	string			`json:"tradeDealStatus"`
	FromDate 			string			`json:"fromDate"`
	UntilDate 			string			`json:"untilDate"`

}

type DealCountResponse struct {

}

//	Get Wallet
//	http://docs.yacuna.com/api/#api-Wallet-Wallet_get
type GetWalletRequest struct {
	Currency			string			`json:"currency"`
}

type GetWalletResponse struct {
	RequestId			string			`json:"requestId"`
	Status				string			`json:"status"`
	Wallet				Wallet			`json:"wallet"`
}

type Wallet struct {
	WalletId			string			`json:"walletId"`
	WalletStatus 		string			`json:"walletStatus"`
	Accounts			[]WalletAccount	`json:"accounts"`
}

type WalletAccount struct {
	WalletAccountId 	string			`json:"walletAccountId"`
	Currency 			string			`json:"currency"`
	AccountType 		string			`json:"accountType"`
	AccountStatus		string			`json:"accountStatus"`
	AccountName			string			`json:"accountName"`
	AccountBalance		AccountBalance	`json:"accountBalance"`
}

type AccountBalance struct {
	Balance 			Money			`json:"balance"`
	ReservedBalance		Money			`json:"reservedBalance"`
	InOrders			Money			`json:"inOrders"`
}

type Money struct {
	Currency			string			`json:"currency"`
	Amount 				float64			`json:"amount"`
}

var ErrUnexpectedType = errors.New("unexpected type in response")
var ErrClientError = errors.New("client error (4xx)")
var ErrServerError = errors.New("server error (5xx)")


func (r *Instance) GetDealCount(req *DealCountRequest) (*DealCountResponse, error) {

	respObj := &DealCountResponse{}
	res := r.api.Res("deal/count", respObj)

	r.setAuthentication(res, "deal/count")
	re, err := doGet(res, req)
	if err != nil {
		return nil, err
	}

	ret, ok := re.Response.(DealCountResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return &ret, nil

}

func (r *Instance) GetWallet(req *GetWalletRequest) (*GetWalletResponse, error) {

	respObj := &GetWalletResponse{}
	res := r.api.Res("wallet/get", respObj)

	r.setAuthentication(res, "wallet/get")
	re, err := doGet(res, req)
	if err != nil {
		return nil, err
	}

	ret, ok := re.Response.(*GetWalletResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return ret, nil

}

func doGet(res *gopencils.Resource, req interface{}) (*gopencils.Resource, error) {

	re, err := res.Get( *toStringMap(req) )
	if err != nil {
		return nil, err
	}

	if re.Raw.StatusCode >= 500 {
		return nil, ErrServerError
	}
	if re.Raw.StatusCode >= 400 {
		return nil, ErrClientError
	}

	return re, err
}

func (r *Instance) setAuthentication(res *gopencils.Resource, uriPath string) {

	ti := &apiTokenInput{
		secret: r.secret,
		method: "GET",
		path: 	API_PATH + "/" + uriPath,
		query: 	r.api.QueryValues.Encode(),
	}

	res.SetHeader(H_ApiTokenId, r.id)
	res.SetHeader(H_ApiToken, ApiToken( ti ) )
	res.SetHeader(H_ApiTokenOTP, "")

}

func toStringMap(req interface{}) *map[string]string {

	rm := map[string]string{}
	for key, value := range structs.Map(req) {
		strVal := fmt.Sprintf("%v", value)
		if strVal != "" {
			rm[key] = strVal
		}
	}

	return &rm
}
