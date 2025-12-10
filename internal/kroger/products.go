package kroger

import (
	"fmt"
	"io"
)

type ProductResponse struct {
}

// GetProducts fetches products from the Kroger API based on the search term.
func (c *Client) GetProducts(searchTerm string) (*ProductResponse, error) {
	resp, err := c.Get(fmt.Sprintf("https://api.kroger.com/v1/products?filter.term=%s", searchTerm))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Product response: ", string(bodyBytes))

	productResp := &ProductResponse{}
	//err = json.Unmarshal(bodyBytes, productResp)
	//if err != nil {
	//	return nil, err
	//}
	return productResp, nil

}
