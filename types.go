package goyacuna

import (
	"time"
)

type Response struct {
	RequestId			string			`json:"requestId"`
	Status				string			`json:"status"`
}

type PagingInfo struct {
	RequestedStartWith 	int 			`json:"requestedStartWith"`
	RequestedCount 		int 			`json:"requestedCount"`
	ActualCount 		int 			`json:"actualCount"`
	TotalCount 			int 			`json:"totalCount"`
	TotalPageCount 		int 			`json:"totalPageCount"`
	ShowPrevious 		bool 			`json:"showPrevious"`
	CurrentPageNumber 	int 			`json:"currentPageNumber"`
	ShowNext 			bool 			`json:"showNext"`
}

type Money struct {
	Currency			string			`json:"currency"`
	Amount 				float64			`json:"amount"`
}


type AccountBalance struct {
	Balance 			Money			`json:"balance"`
	ReservedBalance		Money			`json:"reservedBalance"`
	InOrders			Money			`json:"inOrders"`
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
	Count	 			int				`json:"count"`
	Response
}



// Get Deal
// http://docs.yacuna.com/api/#api-Deal-Deal_get
type GetDealRequest struct {
	DealId				string			`json:"dealId"`
}

type GetDealResponse struct {
	Response
	TradeDeal 			TradeDeal		`json:"tradeDeal"`
}



//	Deal List
//	http://docs.yacuna.com/api/#api-Deal-Deal_list
type DealListRequest struct {

	DealCountRequest
	StartWith			int 			`json:"startWith"`
	Count				int 			`json:"count"`
	Sorting				string 			`json:"sorting"`

}

type DealListResponse struct {
	Response
	PagingInfo 			PagingInfo 		`json:"pagingInfo"`
	TradeDeals			[]TradeDeal		`json:"tradeDeals"`
}



type TradeDealType string
const (
	Buy  string = "Buy"
	Sell string = "Sell"
)

type TradeDealSubtype string
const (
	MarketTaker string = "MarketTaker"
	MarketMaker	string = "MarketMaker"
)

type TradeDealStatus string
const (
	Created 	string = "Created"
	Completed	string = "Completed"
)

type TradeDeal struct {
	Id 					string 			`json:"id"`
	CreationDateTime 	time.Time 		`json:"creationDateTime"`
	WalletId 			string 			`json:"walletId"`
	WalletAccountId 	string 			`json:"walletAccountId"`
	MarketId 			string 			`json:"marketId"`
	OrderId 			string 			`json:"orderId"`
	SequenceNumber 		int 			`json:"sequenceNumber"`
	TradeDealType 		TradeDealType	`json:"tradeDealType"`
	TradeDealSubtype 	TradeDealSubtype`json:"tradeDealSubtype"`
	TradeDealStatus 	TradeDealStatus	`json:"tradeDealStatus"`
	Amount 				Money 			`json:"amount"`
	AmountSold 			Money 			`json:"amountSold"`
	Price 				Money 			`json:"price"`
	FeeAmount 			Money 			`json:"feeAmount"`
}


//	Get Wallet
//	http://docs.yacuna.com/api/#api-Wallet-Wallet_get
type GetWalletRequest struct {
	Currency			string			`json:"currency"`
}

type GetWalletResponse struct {
	Response
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
