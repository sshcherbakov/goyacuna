package goyacuna

import (
	"time"
)

type Response struct {
	RequestId    string         `json:"requestId"`
	Status       string         `json:"status"`
	ErrorDetails []ErrorDetails `json:"errorDetails"`
}

type ErrorDetails struct {
	SimpleName        string `json:"simpleName"`
	Description       string `json:"description"`
	ErrorName         string `json:"errorName"`
	ErrorNumber       int    `json:"errorNumber"`
	ErrorMessageShort string `json:"errorMessageShort"`
	ErrorMessageLong  string `json:"errorMessageLong"`
	FieldName         string `json:"fieldName"`
	LimitsViolated    string `json:"limitsViolated"`
}

type PagingInfo struct {
	RequestedStartWith int  `json:"requestedStartWith"`
	RequestedCount     int  `json:"requestedCount"`
	ActualCount        int  `json:"actualCount"`
	TotalCount         int  `json:"totalCount"`
	TotalPageCount     int  `json:"totalPageCount"`
	ShowPrevious       bool `json:"showPrevious"`
	CurrentPageNumber  int  `json:"currentPageNumber"`
	ShowNext           bool `json:"showNext"`
}

type Money struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type AccountBalance struct {
	Balance         Money `json:"balance"`
	ReservedBalance Money `json:"reservedBalance"`
	InOrders        Money `json:"inOrders"`
}

type CountResponse struct {
	Count int `json:"count"`
	Response
}

//	Deal count
//	http://docs.yacuna.com/api/#api-Deal-Deal_count
type DealCountRequest struct {
	WalletAccountId string 		`json:"walletAccountId"`
	MarketId        string 		`json:"marketId"`
	OrderId         string 		`json:"orderId"`
	TradeDealType   string 		`json:"tradeDealType"`
	TradeDealStatus string 		`json:"tradeDealStatus"`
	FromDate        time.Time 	`json:"fromDate"`
	UntilDate       time.Time	`json:"untilDate"`
}

// Get Deal
// http://docs.yacuna.com/api/#api-Deal-Deal_get
type GetDealRequest struct {
	DealId string `json:"dealId"`
}

type GetDealResponse struct {
	Response
	TradeDeal TradeDeal `json:"tradeDeal"`
}

//	Deal List
//	http://docs.yacuna.com/api/#api-Deal-Deal_list
type DealListRequest struct {
	DealCountRequest
	StartWith int    `json:"startWith"`
	Count     int    `json:"count"`
	Sorting   string `json:"sorting"`
}

type DealListResponse struct {
	Response
	PagingInfo PagingInfo  `json:"pagingInfo"`
	TradeDeals []TradeDeal `json:"tradeDeals"`
}

type TradeDealType string

const (
	TDT_Buy  string = "Buy"
	TDT_Sell string = "Sell"
)

type TradeDealSubtype string

const (
	TDS_MarketTaker string = "MarketTaker"
	TDS_MarketMaker string = "MarketMaker"
)

type TradeStatus string

const (
	TS_Created   string = "Created"   // Created, but not yet active
	TS_Confirmed string = "Comfirmed" // Confirmed, to or in order-book
	TS_Completed string = "Completed" // Order has been completely executed in at least 1 TradeDeal
	TS_Cancelled string = "Cancelled" // Order has been cancelled
	TS_Expired   string = "Expired"   // Order has been expired - typically after 1 day without being confirmed
	TS_Rejected  string = "Rejected"  // Order has been rejected by the market
)

type TradeCommon struct {
	Id               string    `json:"id"`
	CreationDateTime time.Time `json:"creationDateTime"`
	WalletId         string    `json:"walletId"`
	WalletAccountId  string    `json:"walletAccountId"`
	MarketId         string    `json:"marketId"`
}

type TradeDeal struct {
	TradeCommon
	OrderId          string           `json:"orderId"`
	SequenceNumber   int              `json:"sequenceNumber"`
	TradeDealType    TradeDealType    `json:"tradeDealType"`
	TradeDealSubtype TradeDealSubtype `json:"tradeDealSubtype"`
	TradeDealStatus  TradeStatus      `json:"tradeDealStatus"`
	Amount           Money            `json:"amount"`
	AmountSold       Money            `json:"amountSold"`
	Price            Money            `json:"price"`
	FeeAmount        Money            `json:"feeAmount"`
}

//	Market count
//	http://docs.yacuna.com/api/#api-Market-Market_count
type MarketCountRequest struct {
	Currency1 string `json:"currency1"`
	Currency2 string `json:"currency2"`
}

//	Market list
//	http://docs.yacuna.com/api/#api-Market-Market_list
type MarketListRequest struct {
	MarketCountRequest
	StartWith int `json:"startWith"`
	Count     int `json:"count"`
}

type MarketListResponse struct {
	Response
	PagingInfo PagingInfo `json:"pagingInfo"`
	Markets    []Market   `json:"markets"`
}

type Market struct {
	Id                  string              `json:"id"`
	CurrencyCode1       string              `json:"currencyCode1"`
	CurrencyCode2       string              `json:"currencyCode2"`
	Description         string              `json:"description"`
	MarketStatistics    MarketStatistics    `json:"marketStatistics"`
	Market24hStatistics Market24hStatistics `json:"market24hStatistics"`
	PriceGranularity    Money               `json:"priceGranularity"`
}

type MarketStatistics struct {
	Timestamp      time.Time `json:"time"`
	BuyItemCount   int       `json:"buyItemCount"`
	SellItemCount  int       `json:"sellItemCount"`
	BuyItemAmount  Money     `json:"buyItemAmount"`
	SellItemAmount Money     `json:"sellItemAmount"`
	BuyPrice       Money     `json:"buyPrice"`
	SellPrice      Money     `json:"sellPrice"`
	LastPrice      Money     `json:"lastPrice"`
}

type Market24hStatistics struct {
	OpenPrice    Money `json:"openPrice"`
	ClosePrice   Money `json:"closePrice"`
	LowPrice     Money `json:"lowPrice"`
	HighPrice    Money `json:"highPrice"`
	AveragePrice Money `json:"averagePrice"`
	VolumeC1     Money `json:"volumeC1"`
	VolumeC2     Money `json:"volumeC2"`
}

//	Get MarketDepth
//	http://docs.yacuna.com/api/#api-Market-MarketDepth_get
type GetMarketDepthResponse struct {
	Response
	MarketDepth MarketDepth `json:"marketDepth"`
}

type MarketDepth struct {
	Market Market      `json:"market"`
	OrderBookChart
}

type OrderResponse struct {
	Response
	TradeOrder TradeOrder `json:"tradeOrder"`
}

type TradeOrderType string

const (
	TOT_BuyMarket  string = "BuyMarket"  // Buy at best market price
	TOT_BuyLimit   string = "BuyLimit"   // Buy at defined price or better
	TOT_SellMarket string = "SellMarket" // Sell at best market price
	TOT_SellLimit  string = "SellLimit"  // Sell at defined price or better
)

type TradeOrderMarketStatus string

const (
	TOMS_PendingSendToMarket        string = "PendingSendToMarket"        // Order is to be sent to the market
	TOMS_SentToMarket               string = "SentToMarket"               // Order has been sent to the market
	TOMS_InOrderBook                string = "InOrderBook"                // Order has been received by the market, and placed into the order-book
	TOMS_Completed                  string = "Completed"                  // Order has been completely executed, and removed from the order-book
	TOMS_PendingCancellationRequest string = "PendingCancellationRequest" // Cancellation request is to be sent to the market
	TOMS_CancellationRequested      string = "CancellationRequested"      // Cancellation request has been sent to the market
	TOMS_Cancelled                  string = "Cancelled"                  // Order has been cancelled in the market, and removed from the order-book - though may still have been partially executed
	TOMS_Rejected                   string = "Rejected"                   // Order has been rejected by the market)
)

type TradeOrder struct {
	TradeCommon
	ExternalReferenceId    string                 `json:"externalReferenceId"`
	TradeOrderType         TradeOrderType         `json:"tradeOrderType"`
	TradeOrderStatus       TradeStatus            `json:"tradeOrderStatus"`
	TradeOrderMarketStatus TradeOrderMarketStatus `json:"tradeOrderMarketStatus"`
	PriceLimit             float64                `json:"priceLimit"`
	FillRatio              int                    `json:"fillRatio"`
	Amount                 Money                  `json:"amount"`
	BuyAmount              Money                  `json:"buyAmount"`
	TotalAmountSold        Money                  `json:"totalAmountSold"`
	TotalAmountBought      Money                  `json:"totalAmountBought"`
}

// Get OrderBook
// http://docs.yacuna.com/api/#api-Market-OrderBook_get
type GetOrderBookResponse struct {
	Response
	OrderBook MarketDepth `json:"orderBook"`
}


// Get Order Count
// http://docs.yacuna.com/api/#api-Order-Order_count
type OrderCountRequest struct {
	WalletAccountId			string         			`json:"walletAccountId"`
	MarketId        		string 					`json:"marketId"`
	TradeOrderType			TradeOrderType			`json:"tradeOrderType"`
	TradeOrderStatus		TradeStatus				`json:"tradeOrderStatus"`
	TradeOrderMarketStatus	TradeOrderMarketStatus	`json:"tradeOrderMarketStatus"`
	FromDate        		time.Time 				`json:"fromDate"`
	UntilDate       		time.Time 				`json:"untilDate"`
}

// Create Order
// http://docs.yacuna.com/api/#api-Order-Order_create
type CreateOrderRequest struct {
	WalletAccountId			string         			`json:"walletAccountId"`
	TradeOrderType			TradeOrderType			`json:"tradeOrderType"`
	SellAmount				float64					`json:"sellAmount"`
	SellCurrency			string					`json:"sellCurrency"`
	BuyAmount				float64					`json:"buyAmount"`
	BuyCurrency				string					`json:"buyCurrency"`
	PriceLimitAmount		float64					`json:"priceLimitAmount"`
	PriceLimitCurrency		string					`json:"priceLimitCurrency"`
	ExternalReferenceId		string					`json:"externalReferenceId"`
}


// Get Order by ext ref id
// http://docs.yacuna.com/api/#api-Order-Order_get_by_external_reference_Id
type OrderByExtRefIdRequest struct {
	ExternalReferenceId		string					`json:"externalReferenceId"`
	WalletAccountId			string         			`json:"walletAccountId"`
}


//	Order List
//	http://docs.yacuna.com/api/#api-Order-Order_list
type OrderListRequest struct {
	OrderCountRequest
	StartWith int    `json:"startWith"`
	Count     int    `json:"count"`
}

type OrderListResponse struct {
	Response
	PagingInfo 	PagingInfo  	`json:"pagingInfo"`
	TradeOrders []TradeOrder 	`json:"tradeOrders"`
}


// Get OrderBook Chart
// http://docs.yacuna.com/api/#api-Public_Charts-Orderbook
type OrderBookChart struct {
	Asks   [][]float64 `json:"asks"`
	Bids   [][]float64 `json:"bids"`
}

// Get Trades Chart
// http://docs.yacuna.com/api/#api-Public_Charts-Trades
type TradesChart []Trade

type Trade struct {
	Tid 	int 	`json:"tid"`
	Price 	float64	`json:"price"`
	Amount 	float64	`json:"amount"`
	Date 	int		`json:"date"`
}

//	Get Wallet
//	http://docs.yacuna.com/api/#api-Wallet-Wallet_get
type GetWalletRequest struct {
	Currency string `json:"currency"`
}

type GetWalletResponse struct {
	Response
	Wallet Wallet `json:"wallet"`
}

type Wallet struct {
	WalletId     string          `json:"walletId"`
	WalletStatus string          `json:"walletStatus"`
	Accounts     []WalletAccount `json:"accounts"`
}

type WalletAccount struct {
	WalletAccountId string         `json:"walletAccountId"`
	Currency        string         `json:"currency"`
	AccountType     string         `json:"accountType"`
	AccountStatus   string         `json:"accountStatus"`
	AccountName     string         `json:"accountName"`
	AccountBalance  AccountBalance `json:"accountBalance"`
}
