package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CustomerDebitRequest struct {
	TxnId          string  `json:"txnId"`
	FincleDate     string  `json:"fincleDate"`
	FromAccNumber  string  `json:"fromAccNumber"`
	ToAccNumber    string  `json:"toAccNumber"`
	Amount         float64 `json:"amount"`
	TranParticular string  `json:"tranParticular"`
	MerchantName   string  `json:"merchantName"`
}

type CustomerDebitResponse struct {
	ResponseCode    string  `json:"responseCode"`
	ResponseMessage string  `json:"responseMessage"`
	Content         Content `json:"content"`
}

type Content struct {
	TxnId           string `json:"txnId"`
	FinTxnId        string `json:"finTxnId"`
	Rrn             string `json:"rrn"`
	Remarks2        string `json:"remarks2"`
	FinResponseCode string `json:"finResponseCode"`
	SystemDate      string `json:"systemDate"`
}

func customerDebit(config AppConfigProperties) {
	customerDebitRequest := &CustomerDebitRequest{
		TxnId:          "IN20220523165821",
		FincleDate:     "20220523",
		FromAccNumber:  "001020001347",
		ToAccNumber:    "900103720001",
		TranParticular: "xxxxxxxxxxxxx",
		MerchantName:   "PickMe",
		Amount:         520.00,
	}

	obj, err := json.Marshal(customerDebitRequest)
	fmt.Println("Request : ", string(obj))
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		request, err := http.NewRequest(http.MethodPost, config["customerCredit.url"], bytes.NewBuffer(obj))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "Bearer "+getAccessToken(config))
		response, err := client.Do(request)

		if err != nil {
			fmt.Println("Error occurred while customer debit ", err)
		} else {
			body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println("Error occurred while reading the response", err)
			} else {
				var customerDebitResponse CustomerDebitResponse
				json.Unmarshal(body, &customerDebitResponse)
				fmt.Println("Response : ", customerDebitResponse)
				fmt.Println(customerDebitResponse.Content.TxnId)
			}
		}
	}
}
