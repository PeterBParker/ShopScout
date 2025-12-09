package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type TokenResponse struct {
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func main() {
	// Get store
	req, err := http.NewRequest(http.MethodPost, "https://api.kroger.com/v1/connect/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response status:", resp.Status)

	tokenResp := &TokenResponse{}
	err = json.Unmarshal(bodyBytes, tokenResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("token is: %s", tokenResp.AccessToken)
}
