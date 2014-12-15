package main

import (
	"os"
    "fmt"
	"log"
	"bytes"
	"encoding/json"
	"github.com/sshcherbakov/goyacuna"
)

func main() {

	var buf bytes.Buffer
	logger := log.New(&buf, "yacuna: ", log.Lshortfile)
	logger.Print("Starting test...")

	config := readConfig("conf.json")

	api := goyacuna.Api(config)
/*
	res1, err := api.GetWallet(&goyacuna.GetWalletRequest{})
	printCallResult("GetWallet: ", res1, err )

	res2, err := api.GetDealCount(&goyacuna.DealCountRequest{})
	printCallResult("GetDealCount: ", res2, err )

	res3, err := api.GetDeal("AAEABGOj1rR4C7xxreNhYNjXItSH_Yp93aYNuhH2GuaNiE4UjkoRolFm")
	printCallResult("GetDeal: ", res3, err )

	res4, err := api.GetDealList(&goyacuna.DealListRequest{})
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
		&goyacuna.CreateOrderRequest{ WalletAccountId: "xxx", TradeOrderType: goyacuna.TOT_BuyLimit })
	printCallResult("CreateOrder: ", res9, err )

	res10, err := api.ConfirmOrder(res9.TradeOrder.Id)
	printCallResult("ConfirmOrder: ", res10, err )

	res11, err := api.GetOrderCount(&goyacuna.OrderCountRequest{WalletAccountId: "xxx"})
	printCallResult("GetOrderCount: ", res11, err )

	res12, err := api.GetOrder(res9.TradeOrder.Id)
	printCallResult("GetOrder: ", res12, err )

	res13, err := api.GetOrderByExtRefId("xxx", "")
	printCallResult("GetOrderByExtRefId: ", res13, err )

	res14, err := api.GetOrderList(&goyacuna.OrderListRequest{OrderCountRequest:goyacuna.OrderCountRequest{WalletAccountId: "xxx"}})
	printCallResult("GetOrderList: ", res14, err )

	res15, err := api.CancelOrder(res9.TradeOrder.Id)
	printCallResult("CancelOrder: ", res15, err )
*/
	res16, err := api.GetOrderBookChart("XBT", "EUR")
	printCallResult("GetOrderBookChart: ", res16, err )

	res17, err := api.GetTradesChart("XBT", "EUR", 0)
	printCallResult("GetTradesChart: ", res17, err )

}

func printCallResult(caption string, res interface{}, err error) {
	fmt.Print(caption)
	if err != nil {
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
		return &goyacuna.Config{Url:"https://yacuna.com/api/1"}
	}

	return &config
}
