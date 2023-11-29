package weemesh

import (
	"net/http"
	"net/url"
)

func (c *Client) GetFirmwareVersionsByProductKey(productKey string, tags []string) (*CommonResult[[]map[string]interface{}], error) {

	params := url.Values{}
	params.Set("product_key", productKey)

	for _, tag := range tags {
		params.Add("tags", tag)
	}

	request, err := http.NewRequest(http.MethodGet, c.address+"/"+APIVersion+"/firmware/versions?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[[]map[string]interface{}](request, c)
}

func (c *Client) GetOTALatestJob(sn string) (*CommonResult[map[string]interface{}], error) {
	request, err := http.NewRequest(http.MethodGet, c.address+"/"+APIVersion+"/ota/"+sn+"/latestJob", nil)
	if err != nil {
		return nil, err
	}

	return newRequest[map[string]interface{}](request, c)
}
