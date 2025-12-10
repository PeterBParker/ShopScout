package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResponse struct {
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func GetToken(clientID, clientSecret string) (*string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "product.compact")
	dataString := data.Encode()
	dataReader := strings.NewReader(dataString)
	req, err := http.NewRequest(http.MethodPost, "https://api.kroger.com/v1/connect/oauth2/token", dataReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tokenResp := &TokenResponse{}
	err = json.Unmarshal(bodyBytes, tokenResp)
	if err != nil {
		return nil, err
	}

	return &tokenResp.AccessToken, nil
}
