package weemesh

import (
	"net/http"
	"net/url"
)

func (c *Client) GetThingModel(productKey, organizationID string) (*CommonResult[any], error) {

	params := url.Values{}
	params.Set("product_key", productKey)
	params.Set("organization_id", organizationID)

	request, err := http.NewRequest(http.MethodGet, c.address+"/thing_models?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[any](request, c)
}
