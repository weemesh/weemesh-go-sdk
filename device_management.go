package weemesh

import (
	"net/http"
	"net/url"
)

func (c *Client) GetDeviceState(sn string) (*CommonResult[bool], error) {

	params := url.Values{}
	params.Set("sn", sn)

	request, err := http.NewRequest(http.MethodGet, c.address+"/device/state?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[bool](request, c)
}

func (c *Client) GetDeviceLatestProperty(sn string) (*CommonResult[map[string]interface{}], error) {
	request, err := http.NewRequest(http.MethodGet, c.address+"/device/"+sn+"/latestProperty", nil)
	if err != nil {
		return nil, err
	}

	return newRequest[map[string]interface{}](request, c)
}

func (c *Client) GetDevice(sn string) (*CommonResult[map[string]interface{}], error) {
	request, err := http.NewRequest(http.MethodGet, c.address+"/device/"+sn, nil)
	if err != nil {
		return nil, err
	}

	return newRequest[map[string]interface{}](request, c)
}
