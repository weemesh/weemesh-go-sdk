package weemesh

import (
	"net/http"
	"net/url"
)

func (c *Client) GetAppProductKeys(appName string) (*CommonResult[[]string], error) {

	params := url.Values{}
	params.Set("app_name", appName)

	request, err := http.NewRequest(http.MethodGet, c.address+"/app/productKeys?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[[]string](request, c)
}
