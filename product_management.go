package weemesh

import (
	"net/http"
)

func (c *Client) GetProductsOfMe() (*CommonResult[[]map[string]interface{}], error) {
	request, err := http.NewRequest(http.MethodGet, c.address+"/product", nil)
	if err != nil {
		return nil, err
	}

	return newRequest[[]map[string]interface{}](request, c)
}
