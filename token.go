package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func getAccessToken(config AppConfigProperties) string {
	token := ""
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "username")
	data.Set("password", "password")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(http.MethodPost, config["token.url"], strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth("client", "secret")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error occurred while getting token..", err)
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error occurred while reading response body.. ", err)
		} else {
			var tokenResponse TokenResponse
			json.Unmarshal(body, &tokenResponse)
			token = tokenResponse.AccessToken
		}
	}
	return token
}
