package main

import (
	"os"

	"github.com/PeterBParker/ShopScout/internal/auth"
	"github.com/PeterBParker/ShopScout/internal/kroger"
)

func main() {
	// Get store
	token, err := auth.GetToken(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		// TODO check error and retry?
		panic(err)
	}
	krogerClient := kroger.NewClient(*token)
	_, err = krogerClient.GetProducts("milk")
}
