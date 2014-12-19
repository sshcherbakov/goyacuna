package main

import (
	"os"
    "fmt"
	"encoding/json"
	"github.com/sshcherbakov/goyacuna"
)

var testCount int = 0
var errCount int = 0

func main() {

	config := readConfig("conf.json")

	api := goyacuna.Api(config)


	res1, err := api.GetWallet(&goyacuna.GetWalletRequest{})
	printCallResult("GetWallet: ", res1, err )

	if len(res1.Wallet.Accounts) <= 0 {
		fmt.Println("No wallet accounts")
		return
	}

	waid := res1.WalletIdByCurrency("XBT")

	res2, err := api.GetDealCount(&goyacuna.DealCountRequest{WalletAccountId: waid})
	printCallResult("GetDealCount: ", res2, err )

	res3, err := api.GetDeal("XXXX")
	printCallResult("GetDeal: ", res3, err )

	res4, err := api.GetDealList(&goyacuna.DealListRequest{DealCountRequest:goyacuna.DealCountRequest{WalletAccountId: waid}})
	printCallResult("GetDealList: ", res4, err )

	res5, err := api.GetMarketCount(&goyacuna.MarketCountRequest{})
	printCallResult("GetMarketCount: ", res5, err )

	res6, err := api.GetMarketList(&goyacuna.MarketListRequest{})
	printCallResult("GetMarketList: ", res6, err )

	res7, err := api.GetMarketDepth("XBT", "EUR")
	printCallResult("GetMarketDepth: ", res7, err )

	res8, err := api.GetOrderBook("XBT", "EUR")
	printCallResult("GetOrderBook: ", res8, err )

	res9, err := api.CreateOrder("XBT", "EUR",
		&goyacuna.CreateOrderRequest{
			WalletAccountId: waid,
			TradeOrderType: goyacuna.TOT_SellMarket,
			SellAmount: 0.001,
			SellCurrency: "XBT",
	})
	printCallResult("CreateOrder: ", res9, err )

	res10, err := api.ConfirmOrder(res9.TradeOrder.Id)
	printCallResult("ConfirmOrder: ", res10, err )

	res11, err := api.GetOrderCount(&goyacuna.OrderCountRequest{WalletAccountId: waid})
	printCallResult("GetOrderCount: ", res11, err )

	res12, err := api.GetOrder(res9.TradeOrder.Id)
	printCallResult("GetOrder: ", res12, err )

	res13, err := api.GetOrderByExtRefId(waid, res9.TradeOrder.ExternalReferenceId)
	printCallResult("GetOrderByExtRefId: ", res13, err )

	res14, err := api.GetOrderList(&goyacuna.OrderListRequest{OrderCountRequest:goyacuna.OrderCountRequest{WalletAccountId: waid}})
	printCallResult("GetOrderList: ", res14, err )

	res15, err := api.CancelOrder(res9.TradeOrder.Id)
	printCallResult("CancelOrder: ", res15, err )

	res16, err := api.GetOrderBookChart("XBT", "EUR")
	printCallResult("GetOrderBookChart: ", res16, err )

	res17, err := api.GetTradesChart("XBT", "EUR", 0)
	printCallResult("GetTradesChart: ", res17, err )

	if errCount != testCount {
		fmt.Println()
		fmt.Printf("Failed tests %d/%d", errCount, testCount)
	}

}

func printCallResult(caption string, res interface{}, err error) {
	fmt.Print(caption)
	testCount++
	if err != nil {
		errCount++
		fmt.Printf( "(%s) ", err )
	}
	jres,_ := json.Marshal(res)
	fmt.Println(string(jres))
}

func readConfig(fileName string) *goyacuna.Config {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	config := goyacuna.Config{}
	err := decoder.Decode(&config)
	if err != nil {
		return &goyacuna.Config{Url:"https://sandbox.yacuna.com/api/1"}
	}

	return &config
}
