package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type resultTx struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Result  []blockTx `json:"result"`
}
type blockTx struct {
	BlockNumber     string `json:"blockNumber"`
	Timestamp       string `json:"timeStamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
}
type resultPrice struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  blockPrice `json:"result"`
}
type blockPrice struct {
	Ethbtc          string `json:"ethbtc"`
	EthbtcTimestamp string `json:"ethbtc_timestamp"`
	Bnbusd          string `json:"ethusd"`
	BnbusdTimestamp string `json:"ethusd_timestamp"`
}

var APIKEY = "push_you_key_here"
var DefaultBNBPrice float64

func getLastPriceBNB() (usd float64, err error) {
	usd = DefaultBNBPrice
	url := "https://api-testnet.bscscan.com/api?module=stats&action=bnbprice&apikey=" + APIKEY
	resp, err := callBsc(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println("error after call bsc api : ", err)
		return usd, err
	}

	decoder := json.NewDecoder(resp.Body)
	var r resultPrice
	err = decoder.Decode(&r)
	if err != nil {
		log.Println("error after call decode to resultPrice : ", err)
		return usd, err
	}

	usdAsFloat, err := strconv.ParseFloat(r.Result.Bnbusd, 64)
	if err != nil {
		log.Println("error case convert to int usdAsFloat : ", err)
		return usd, err
	}
	return usdAsFloat, nil
}

func changeTxValueToTokenPrice(valueNumberFromApi string) float64 {
	valueAsFloat, err := strconv.ParseFloat(valueNumberFromApi, 64)
	if err != nil {
		log.Println("error after call bsc api : ", err)
		return 0.0
	}
	return valueAsFloat / 1000000000000000000
}

func convertBNBtoUsd(amountBnb float64) float64 {
	lastPriceBNBtoUSD, err := getLastPriceBNB()
	if err != nil {
		log.Println("error case lastPriceBNBtoUSD")
		return 0.0
	}
	return lastPriceBNBtoUSD * amountBnb
}

func callBsc(url string) (resp *http.Response, err error) {
	var client http.Client
	resp, err = client.Get(url)
	if err != nil {
		log.Println("error in callBsc", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		return resp, nil
	}
	err = errors.New("StatusCode is return not 200")
	return nil, err
}

func checkTx(t time.Time) {

	walletAddress := "0x891B68D6B21c64d56dB262D066B38Ea76B6468f6"
	lastBlockLoad := 1
	lastBlockLoadStr := strconv.Itoa(lastBlockLoad)
	url := "https://api-testnet.bscscan.com/api?module=account&action=txlist&address=" + walletAddress + "&startblock=+" + lastBlockLoadStr + "+&endblock=99999999&page=1&offset=5&sort=desc&apikey=" + APIKEY

	resp, err := callBsc(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println("error after call bsc", err)
		return
	}
	//get json body need to decode
	decoder := json.NewDecoder(resp.Body)
	var r resultTx
	err = decoder.Decode(&r)
	if err != nil {
		log.Println("error after Decode checkTx", err)
		return
	}

	value := r.Result[0].Value
	tokenAmount := changeTxValueToTokenPrice(value)

	//todo make it support many more token
	tokenText := "BNB"
	usdAmount := convertBNBtoUsd(tokenAmount)
	usdAmount = math.Floor(usdAmount*100) / 100

	usdAmountStr := fmt.Sprintf("%.2f", usdAmount)
	tokenAmountStr := fmt.Sprintf("%.4f", tokenAmount)

	fromAccount := r.Result[0].From
	toAccount := r.Result[0].To
	lineAccountName := "Patara"

	textGetNewToken := "Hi " + lineAccountName + " you get a new " + tokenAmountStr + " " + tokenText + " (~" + usdAmountStr + " USD) transfer from account" + fromAccount + "!"
	textSentNewToken := "Hi " + lineAccountName + " you have send " + tokenAmountStr + " " + tokenText + " (~" + usdAmountStr + " USD) transfer to account" + toAccount + "!"

	textSendAleartToLine := textGetNewToken
	if fromAccount == walletAddress {
		textSendAleartToLine = textSentNewToken
	}

	log.Println("will send to line with this message ", textSendAleartToLine)
	//todo can check lastBlockNumber and call line-api
	//lastBlockNumber := r.Result[0].BlockNumber
	//if (lastBlockNumber > lastBlockNumberInDatabase) SendPushMessageLine(textSendAlearttoLine)

}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func main() {
	DefaultBNBPrice = 250.0
	//limit free user bsc api call per day is 100,000 and 5 call per sec
	doEvery(5*time.Second, checkTx)
}
