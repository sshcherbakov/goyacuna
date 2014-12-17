package goyacuna

import (
	"time"
)

type Response struct {
	RequestId    		string         `json:"requestId" structs:"requestId"`
	Status       		string         `json:"status" structs:"status"`
	ErrorDetails 		[]ErrorDetails `json:"errorDetails" structs:"errorDetails"`
}

type ErrorDetails struct {
	FullName			string 			`json:"fullName" structs:"fullName"`
	SimpleName			string 			`json:"simpleName" structs:"simpleName"`
	Description			string 			`json:"description" structs:"description,omitempty"`
	ErrorName			string 			`json:"errorName" structs:"errorName"`
	ErrorNumber			int    			`json:"errorNumber" structs:"errorNumber"`
	ErrorMessageShort	string 			`json:"errorMessageShort" structs:"errorMessageShort"`
	ErrorMessageLong	string 			`json:"errorMessageLong" structs:"errorMessageLong"`
	FieldName			string			`json:"fieldName" structs:"fieldName"`
	LimitsViolated		string 			`json:"limitsViolated" structs:"limitsViolated,omitempty"`
}

type PagingInfo struct {
	RequestedStartWith int  `json:"requestedStartWith" structs:"requestedStartWith"`
	RequestedCount     int  `json:"requestedCount" structs:"requestedCount"`
	ActualCount        int  `json:"actualCount" structs:"actualCount"`
	TotalCount         int  `json:"totalCount" structs:"totalCount"`
	TotalPageCount     int  `json:"totalPageCount" structs:"totalPageCount"`
	ShowPrevious       bool `json:"showPrevious" structs:"showPrevious"`
	CurrentPageNumber  int  `json:"currentPageNumber" structs:"currentPageNumber"`
	ShowNext           bool `json:"showNext" structs:"showNext"`
}

type Money struct {
	Currency string  `json:"currency" structs:"currency"`
	Amount   float64 `json:"amount" structs:"amount"`
}

type AccountBalance struct {
	Balance         Money `json:"balance" structs:"balance"`
	ReservedBalance Money `json:"reservedBalance" structs:"reservedBalance"`
	InOrders        Money `json:"inOrders" structs:"inOrders"`
}

type CountResponse struct {
	Count int `json:"count" structs:"count"`
	Response
}

//	Deal count
//	http://docs.yacuna.com/api/#api-Deal-Deal_count
type DealCountRequest struct {
	WalletAccountId string 		`json:"walletAccountId" structs:"walletAccountId"`
	MarketId        string 		`json:"marketId" structs:"marketId,omitempty"`
	OrderId         string 		`json:"orderId" structs:"orderId,omitempty"`
	TradeDealType   TradeDealType 	`json:"tradeDealType" structs:"tradeDealType,omitnested,omitempty"`
	TradeDealStatus TradeDealStatus	`json:"tradeDealStatus" structs:"tradeDealStatus,omitnested,omitempty"`
	FromDate        time.Time 	`json:"fromDate" structs:"fromDate,omitnested,omitempty"`
	UntilDate       time.Time	`json:"untilDate" structs:"untilDate,omitnested,omitempty"`
}

// Get Deal
// http://docs.yacuna.com/api/#api-Deal-Deal_get
type GetDealResponse struct {
	Response
	TradeDeal TradeDeal `json:"tradeDeal" structs:"tradeDeal"`
}

//	Deal List
//	http://docs.yacuna.com/api/#api-Deal-Deal_list
type DealListRequest struct {
	DealCountRequest
	StartWith int    `json:"startWith" structs:"startWith,omitempty"`
	Count     int    `json:"count" structs:"count,omitempty"`
	Sorting   string `json:"sorting" structs:"sorting,omitempty"`
}

type DealListResponse struct {
	Response
	PagingInfo PagingInfo  `json:"pagingInfo" structs:"pagingInfo"`
	TradeDeals []TradeDeal `json:"tradeDeals" structs:"tradeDeals"`
}

type TradeDealType string

const (
	TDT_Buy  TradeDealType = "Buy"
	TDT_Sell TradeDealType = "Sell"
)

type TradeDealStatus string

const (
	TDS_Created  	TradeDealStatus = "Created"
	TDS_Completed 	TradeDealStatus = "Completed"
)

type TradeDealSubtype string

const (
	TDS_MarketTaker TradeDealSubtype = "MarketTaker"
	TDS_MarketMaker TradeDealSubtype = "MarketMaker"
)

type TradeStatus string

const (
	TS_Created   TradeStatus = "Created"   // Created, but not yet active
	TS_Confirmed TradeStatus = "Comfirmed" // Confirmed, to or in order-book
	TS_Completed TradeStatus = "Completed" // Order has been completely executed in at least 1 TradeDeal
	TS_Cancelled TradeStatus = "Cancelled" // Order has been cancelled
	TS_Expired   TradeStatus = "Expired"   // Order has been expired - typically after 1 day without being confirmed
	TS_Rejected  TradeStatus = "Rejected"  // Order has been rejected by the market
)

type TradeCommon struct {
	Id               string    `json:"id" structs:"id"`
	CreationDateTime time.Time `json:"creationDateTime" structs:"creationDateTime,omitnested"`
	WalletId         string    `json:"walletId" structs:"walletId"`
	WalletAccountId  string    `json:"walletAccountId" structs:"walletAccountId"`
	MarketId         string    `json:"marketId" structs:"marketId"`
}

type TradeDeal struct {
	TradeCommon
	OrderId          string           `json:"orderId" structs:"orderId"`
	SequenceNumber   int              `json:"sequenceNumber" structs:"sequenceNumber"`
	TradeDealType    TradeDealType    `json:"tradeDealType" structs:"tradeDealType"`
	TradeDealSubtype TradeDealSubtype `json:"tradeDealSubtype" structs:"tradeDealSubtype"`
	TradeDealStatus  TradeStatus      `json:"tradeDealStatus" structs:"tradeDealStatus"`
	Amount           Money            `json:"amount" structs:"amount"`
	AmountSold       Money            `json:"amountSold" structs:"amountSold"`
	Price            Money            `json:"price" structs:"price"`
	FeeAmount        Money            `json:"feeAmount" structs:"feeAmount"`
}

//	Market count
//	http://docs.yacuna.com/api/#api-Market-Market_count
type MarketCountRequest struct {
	Currency1 string `json:"currency1" structs:"currency1,omitempty"`
	Currency2 string `json:"currency2" structs:"currency2,omitempty"`
}

//	Market list
//	http://docs.yacuna.com/api/#api-Market-Market_list
type MarketListRequest struct {
	MarketCountRequest
	StartWith int `json:"startWith" structs:"startWith,omitempty"`
	Count     int `json:"count" structs:"count,omitempty"`
}

type MarketListResponse struct {
	Response
	PagingInfo PagingInfo `json:"pagingInfo" structs:"pagingInfo"`
	Markets    []Market   `json:"markets" structs:"markets"`
}

type Market struct {
	Id                  string              `json:"id" structs:"id"`
	CurrencyCode1       string              `json:"currencyCode1" structs:"currencyCode1"`
	CurrencyCode2       string              `json:"currencyCode2" structs:"currencyCode2"`
	Description         string              `json:"description" structs:"description"`
	MarketStatistics    MarketStatistics    `json:"marketStatistics" structs:"marketStatistics"`
	Market24hStatistics Market24hStatistics `json:"market24hStatistics" structs:"market24hStatistics"`
	PriceGranularity    Money               `json:"priceGranularity" structs:"priceGranularity"`
}

type MarketStatistics struct {
	Timestamp      time.Time `json:"time" structs:"time,omitnested"`
	BuyItemCount   int       `json:"buyItemCount" structs:"buyItemCount"`
	SellItemCount  int       `json:"sellItemCount" structs:"sellItemCount"`
	BuyItemAmount  Money     `json:"buyItemAmount" structs:"buyItemAmount"`
	SellItemAmount Money     `json:"sellItemAmount" structs:"sellItemAmount"`
	BuyPrice       Money     `json:"buyPrice" structs:"buyPrice"`
	SellPrice      Money     `json:"sellPrice" structs:"sellPrice"`
	LastPrice      Money     `json:"lastPrice" structs:"lastPrice"`
}

type Market24hStatistics struct {
	OpenPrice    Money `json:"openPrice" structs:"openPrice"`
	ClosePrice   Money `json:"closePrice" structs:"closePrice"`
	LowPrice     Money `json:"lowPrice" structs:"lowPrice"`
	HighPrice    Money `json:"highPrice" structs:"highPrice"`
	AveragePrice Money `json:"averagePrice" structs:"averagePrice"`
	VolumeC1     Money `json:"volumeC1" structs:"volumeC1"`
	VolumeC2     Money `json:"volumeC2" structs:"volumeC2"`
}

//	Get MarketDepth
//	http://docs.yacuna.com/api/#api-Market-MarketDepth_get
type GetMarketDepthResponse struct {
	Response
	MarketDepth MarketDepth `json:"marketDepth" structs:"marketDepth"`
}

type MarketDepth struct {
	Market Market      `json:"market" structs:"market"`
	OrderBookChart
}

type OrderResponse struct {
	Response
	TradeOrder TradeOrder `json:"tradeOrder" structs:"tradeOrder"`
}

type TradeOrderType string

const (
	TOT_BuyMarket  TradeOrderType = "BuyMarket"  // Buy at best market price
	TOT_BuyLimit   TradeOrderType = "BuyLimit"   // Buy at defined price or better
	TOT_SellMarket TradeOrderType = "SellMarket" // Sell at best market price
	TOT_SellLimit  TradeOrderType = "SellLimit"  // Sell at defined price or better
)

type TradeOrderMarketStatus string

const (
	TOMS_PendingSendToMarket        TradeOrderMarketStatus = "PendingSendToMarket"        // Order is to be sent to the market
	TOMS_SentToMarket               TradeOrderMarketStatus = "SentToMarket"               // Order has been sent to the market
	TOMS_InOrderBook                TradeOrderMarketStatus = "InOrderBook"                // Order has been received by the market, and placed into the order-book
	TOMS_Completed                  TradeOrderMarketStatus = "Completed"                  // Order has been completely executed, and removed from the order-book
	TOMS_PendingCancellationRequest TradeOrderMarketStatus = "PendingCancellationRequest" // Cancellation request is to be sent to the market
	TOMS_CancellationRequested      TradeOrderMarketStatus = "CancellationRequested"      // Cancellation request has been sent to the market
	TOMS_Cancelled                  TradeOrderMarketStatus = "Cancelled"                  // Order has been cancelled in the market, and removed from the order-book - though may still have been partially executed
	TOMS_Rejected                   TradeOrderMarketStatus = "Rejected"                   // Order has been rejected by the market)
)

type TradeOrder struct {
	TradeCommon
	ExternalReferenceId    string                 `json:"externalReferenceId" structs:"externalReferenceId"`
	TradeOrderType         TradeOrderType         `json:"tradeOrderType" structs:"tradeOrderType"`
	TradeOrderStatus       TradeStatus            `json:"tradeOrderStatus" structs:"tradeOrderStatus"`
	TradeOrderMarketStatus TradeOrderMarketStatus `json:"tradeOrderMarketStatus" structs:"tradeOrderMarketStatus"`
	PriceLimit             float64                `json:"priceLimit" structs:"priceLimit"`
	FillRatio              int                    `json:"fillRatio" structs:"fillRatio"`
	Amount                 Money                  `json:"amount" structs:"amount"`
	BuyAmount              Money                  `json:"buyAmount" structs:"buyAmount"`
	TotalAmountSold        Money                  `json:"totalAmountSold" structs:"totalAmountSold"`
	TotalAmountBought      Money                  `json:"totalAmountBought" structs:"totalAmountBought"`
}

// Get OrderBook
// http://docs.yacuna.com/api/#api-Market-OrderBook_get
type GetOrderBookResponse struct {
	Response
	OrderBook MarketDepth `json:"orderBook" structs:"orderBook"`
}


// Get Order Count
// http://docs.yacuna.com/api/#api-Order-Order_count
type OrderCountRequest struct {
	WalletAccountId			string         			`json:"walletAccountId" structs:"walletAccountId"`
	MarketId        		string 					`json:"marketId" structs:"marketId,omitempty"`
	TradeOrderType			TradeOrderType			`json:"tradeOrderType" structs:"tradeOrderType,omitempty"`
	TradeOrderStatus		TradeStatus				`json:"tradeOrderStatus" structs:"tradeOrderStatus,omitempty"`
	TradeOrderMarketStatus	TradeOrderMarketStatus	`json:"tradeOrderMarketStatus" structs:"tradeOrderMarketStatus,omitempty"`
	FromDate        		time.Time 				`json:"fromDate" structs:"fromDate,omitnested,omitempty"`
	UntilDate       		time.Time 				`json:"untilDate" structs:"untilDate,omitnested,omitempty"`
}

// Create Order
// http://docs.yacuna.com/api/#api-Order-Order_create
type CreateOrderRequest struct {
	WalletAccountId			string         			`json:"walletAccountId" structs:"walletAccountId"`
	TradeOrderType			TradeOrderType			`json:"tradeOrderType" structs:"tradeOrderType,omitempty"`
	SellAmount				float64					`json:"sellAmount" structs:"sellAmount,omitempty"`
	SellCurrency			string					`json:"sellCurrency" structs:"sellCurrency,omitempty"`
	BuyAmount				float64					`json:"buyAmount" structs:"buyAmount,omitempty"`
	BuyCurrency				string					`json:"buyCurrency" structs:"buyCurrency,omitempty"`
	PriceLimitAmount		float64					`json:"priceLimitAmount" structs:"priceLimitAmount,omitempty"`
	PriceLimitCurrency		string					`json:"priceLimitCurrency" structs:"priceLimitCurrency,omitempty"`
	ExternalReferenceId		string					`json:"externalReferenceId" structs:"externalReferenceId,omitempty"`
}


// Get Order by ext ref id
// http://docs.yacuna.com/api/#api-Order-Order_get_by_external_reference_Id
type OrderByExtRefIdRequest struct {
	ExternalReferenceId		string					`json:"externalReferenceId" structs:"externalReferenceId"`
	WalletAccountId			string         			`json:"walletAccountId" structs:"walletAccountId"`
}


//	Order List
//	http://docs.yacuna.com/api/#api-Order-Order_list
type OrderListRequest struct {
	OrderCountRequest
	StartWith int    `json:"startWith" structs:"startWith,omitempty"`
	Count     int    `json:"count" structs:"count,omitempty"`
}

type OrderListResponse struct {
	Response
	PagingInfo 	PagingInfo  	`json:"pagingInfo" structs:"pagingInfo"`
	TradeOrders []TradeOrder 	`json:"tradeOrders" structs:"tradeOrders"`
}


// Get OrderBook Chart
// http://docs.yacuna.com/api/#api-Public_Charts-Orderbook
type OrderBookChart struct {
	Asks   [][]float64 `json:"asks" structs:"asks"`
	Bids   [][]float64 `json:"bids" structs:"bids"`
}

// Get Trades Chart
// http://docs.yacuna.com/api/#api-Public_Charts-Trades
type TradesChart []Trade

type Trade struct {
	Tid 	int 	`json:"tid" structs:"tid"`
	Price 	float64	`json:"price" structs:"price"`
	Amount 	float64	`json:"amount" structs:"amount"`
	Date 	int		`json:"date" structs:"date"`
}

//	Get Wallet
//	http://docs.yacuna.com/api/#api-Wallet-Wallet_get
type GetWalletRequest struct {
	Currency string `json:"currency" structs:"currency,omitempty"`
}

type GetWalletResponse struct {
	Response
	Wallet Wallet `json:"wallet" structs:"wallet"`
}

type Wallet struct {
	WalletId     string          `json:"walletId" structs:"walletId"`
	WalletStatus string          `json:"walletStatus" structs:"walletStatus"`
	Accounts     []WalletAccount `json:"accounts" structs:"accounts"`
}

type WalletAccount struct {
	WalletAccountId string         `json:"walletAccountId" structs:"walletAccountId"`
	Currency        string         `json:"currency" structs:"currency"`
	AccountType     string         `json:"accountType" structs:"accountType"`
	AccountStatus   string         `json:"accountStatus" structs:"accountStatus"`
	AccountName     string         `json:"accountName" structs:"accountName"`
	AccountBalance  AccountBalance `json:"accountBalance" structs:"accountBalance"`
}
