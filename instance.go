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

var ErrUnexpectedType = errors.New("unexpected type in response")
var ErrClientError = errors.New("client error (4xx)")
var ErrServerError = errors.New("server error (5xx)")


//	Deal Count
//	http://docs.yacuna.com/api/#api-Deal-Deal_count
func (r *Instance) GetDealCount(req *DealCountRequest) (*DealCountResponse, error) {

	uriPath := "deal/count"
	respObj := &DealCountResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setAuthentication(res, uriPath)
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

// Deal Get
// http://docs.yacuna.com/api/#api-Deal-Deal_get
func (r *Instance) GetDeal(dealId string) (*GetDealResponse, error) {

	uriPath := "deal/get/" + dealId
	respObj := &GetDealResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setAuthentication(res, uriPath)
	re, err := doGet(res)
	if err != nil {
		return nil, err
	}

	ret, ok := re.Response.(GetDealResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return &ret, nil

}

//	Deal List
//	http://docs.yacuna.com/api/#api-Deal-Deal_list
func (r *Instance) GetDealList(req *DealListRequest) (*DealListResponse, error) {

	uriPath := "deal/list"
	respObj := &DealListResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setAuthentication(res, uriPath)
	re, err := doGet(res, req)
	if err != nil {
		return nil, err
	}

	ret, ok := re.Response.(DealListResponse)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return &ret, nil

}

//	Get Wallet
//	http://docs.yacuna.com/api/#api-Wallet-Wallet_get
func (r *Instance) GetWallet(req *GetWalletRequest) (*GetWalletResponse, error) {

	uriPath := "wallet/get"
	respObj := &GetWalletResponse{}
	res := r.api.Res(uriPath, respObj)

	r.setAuthentication(res, uriPath)
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

func doGet(res *gopencils.Resource, req ...interface{}) (*gopencils.Resource, error) {

	var re *gopencils.Resource
	var err error
	if len(req) > 0 {
		re, err = res.Get(*toStringMap(req[0]))
	} else {
		re, err = res.Get()
	}

	if err != nil {
		return nil, err
	}


	if re.Raw.StatusCode >= 500 {
		return nil, ErrServerError
	}
	if re.Raw.StatusCode >= 400 {
		return re, ErrClientError
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
