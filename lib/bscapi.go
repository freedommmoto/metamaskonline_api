package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"github.com/freedommmoto/metamaskonline_api/tool"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
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

var DefaultBNBPrice float64

func GetLastPriceBNB(CainInfo db.Chain) (usd float64, err error) {
	usd = DefaultBNBPrice
	chainApi := CainInfo.UrlApi.String

	url := chainApi + "?module=stats&action=bnbprice"
	resp, err := callBsc(url)
	if err != nil {
		return usd, err
	}
	var resultprice resultPrice
	if errUnmarshal := json.Unmarshal(resp, &resultprice); errUnmarshal != nil {
		return usd, errUnmarshal
	}

	usdAsFloat, err := strconv.ParseFloat(resultprice.Result.Bnbusd, 64)
	if err != nil {
		return usd, err
	}
	return usdAsFloat, nil
}

func changeTxValueToTokenPrice(valueNumberFromApi string) float64 {
	valueAsFloat, err := strconv.ParseFloat(valueNumberFromApi, 64)
	if err != nil {
		tool.AddErrorLogIntoFile("error after call bsc api : " + err.Error())
		return 0.0
	}
	return valueAsFloat / 1000000000000000000
}

func convertBNBtoUsd(amountBnb float64) float64 {
	return DefaultBNBPrice * amountBnb
}

func callBsc(url string) (res []byte, err error) {
	//var client http.Client
	req, errNewRequest := http.NewRequest("GET", url, nil)
	req.Close = true
	if errNewRequest != nil {
		return nil, errNewRequest
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody := resp.Body

	b, errReadAll := io.ReadAll(respBody)
	if errReadAll != nil {
		return nil, errReadAll
	}

	//log request and response call bsc api into log file
	tool.AddApiLogIntoFile("request " + "\n" + url)
	tool.AddApiLogIntoFile("response" + "\n" + string(b))

	if err != nil {
		log.Println("error in callBsc", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		return b, nil
	}
	err = errors.New("StatusCode is return not 200")
	return nil, err
}

func GetLastBlockTransactionFromBscScan(CainInfo db.Chain, walletAddress string, lastBlock int32, BSCToken string) (r resultTx, last int32, e error) {
	chainApi := CainInfo.UrlApi.String
	lastBlockStr := strconv.Itoa(int(lastBlock))
	url := chainApi + "?module=account&action=txlist&address=" + walletAddress + "&startblock=+" + lastBlockStr + "+&endblock=99999999&page=1&offset=5&sort=desc&apikey=" + BSCToken
	var nilresp resultTx
	resp, err := callBsc(url)
	if err != nil {
		return nilresp, 0, err
	}

	var tx resultTx
	if errUnmarshal := json.Unmarshal(resp, &tx); errUnmarshal != nil {
		return nilresp, 0, errUnmarshal
	}

	if len(tx.Result) < 1 {
		tool.AddErrorLogIntoFile("call " + walletAddress + " , but not found any Transaction may need to check")
		return tx, 0, nil
	}
	//change block sting to int 32
	i, errAfterParseInt := strconv.ParseInt(tx.Result[0].BlockNumber, 10, 32)
	if errAfterParseInt != nil {
		return tx, 0, errAfterParseInt
	}

	last = int32(i)
	return tx, last, nil
}

func MakePushTextForLineAlert(r resultTx, walletAddress string, AccountName string) (string, error) {
	if r.Result[0].Value == "" {
		return "", errors.New("No Result data for MakePushTextForLineAlert !")
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

	textGetNewToken := "Hi " + AccountName + " you get a new " + tokenAmountStr + " " + tokenText + " (~" + usdAmountStr + " USD) transfer from account" + fromAccount + "!"
	textSentNewToken := "Hi " + AccountName + " you have send " + tokenAmountStr + " " + tokenText + " (~" + usdAmountStr + " USD) transfer to account" + toAccount + "!"

	textSendAleartToLine := textGetNewToken
	if fromAccount == walletAddress {
		textSendAleartToLine = textSentNewToken
	}
	return textSendAleartToLine, nil
}
