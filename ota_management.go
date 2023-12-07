package weemesh

import (
	"bytes"
	"encoding/json"
	"gitee.com/winstar-smart/weemesh-go-sdk/model"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func (c *Client) GetFirmwareVersionsByProductKey(productKey string, tags []string) (*CommonResult[[]map[string]interface{}], error) {

	params := url.Values{}
	params.Set("product_key", productKey)

	for _, tag := range tags {
		params.Add("tags", tag)
	}

	request, err := http.NewRequest(http.MethodGet, c.address+"/firmware/versions?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[[]map[string]interface{}](request, c)
}

func (c *Client) GetOTALatestJob(sn string) (*CommonResult[map[string]interface{}], error) {
	request, err := http.NewRequest(http.MethodGet, c.address+"/ota/"+sn+"/latestJob", nil)
	if err != nil {
		return nil, err
	}

	return newRequest[map[string]interface{}](request, c)
}

func (c *Client) NewFirmware(params *model.NewFirmwareParams) (*CommonResult[any], error) {
	b, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.address+"/firmware", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	return newRequest[any](request, c)
}

func (c *Client) NewFirmwareVersion(params *model.NewFirmwareVersionParams) (*CommonResult[any], error) {
	b, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.address+"/firmware/versions", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	return newRequest[any](request, c)
}

func (c *Client) UploadFirmware(filePath string) (*CommonResult[string], error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filePath)
	defer file.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	_, err = io.Copy(part1, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.address+"/firmware/upload", payload)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	return newRequest[string](request, c)
}
