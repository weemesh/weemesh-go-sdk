package weemesh

import (
	"net/http"
	"net/url"
)

func (c *Client) GetFirmwareVersionsByProductKey(productKey string) (*CommonResult[[]map[string]interface{}], error) {

	params := url.Values{}
	params.Set("product_key", productKey)

	request, err := http.NewRequest(http.MethodGet, c.address+"/"+APIVersion+"/firmware/versions?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[[]map[string]interface{}](request, c)
}
