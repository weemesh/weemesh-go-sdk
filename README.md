# [weemesh/weemesh-go-sdk](https://github.com/weemesh/weemesh-go-sdk)

微麦石物联网平台SDK

## 功能
1. 获取指定产品的物模型
2. 获取指定设备的状态
3. 获取指定设备的概况
4. 获取指定设备的最新属性
5. 根据产品key获取到固件版本列表
6. 上传固件
7. 新增固件
8. 新增固件版本
9. 查看最新ota任务
10. 查看指定应用的product_key列表


## 使用方法

```go
package main

import (
	"gitee.com/winstar-smart/weemesh-go-sdk"
	"gitee.com/winstar-smart/weemesh-go-sdk/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

const (
	accessKey    = "xxxx"
	accessSecret = "xxxx"
	address      = "xxxx"
	productKey   = "xxxx"
	orgId        = "xxxx"
	sn           = "xxxx"
	appName = "xxxx"
	filePath = "/xxx/xxx.bin"
)

var (
	tags             = []string{"xxxx"}
	newFirmwareParam = model.NewFirmwareParams{
		ProductKey:   "xxxx",
		FirmwareName: "xxxx",
		Version:      "xxxx",
		Remark:       "xxxx",
		FirmwareType: 0,
		URL:          "xxxx",
		SignMethod:   0,
		Enable:       false,
		Tags:         tags,
	}

	newFirmwareVersionsParam = model.NewFirmwareVersionParams{
		FirmwareInfoID: primitive.ObjectID{},
		Version:        "xxxx",
		Remark:         "xxxx",
		URL:            "xxxx",
		Enable:         false,
		SignMethod:     0,
		Tags:           tags,
	}
)

func main() {
	c := weemesh.NewClient(accessKey, accessSecret, address)

	// 获取指定产品的物模型
	thingModel, err := c.GetThingModel(productKey, orgId)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(thingModel)

	// 获取指定设备的状态
	deviceState, err := c.GetDeviceState(sn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(deviceState)

	// 获取指定设备的概况
	device, err := c.GetDevice(sn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(device)

	// 获取指定设备的最新属性
	deviceLatestProperty, err := c.GetDeviceLatestProperty(sn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(deviceLatestProperty)

	// 根据产品key获取到固件版本列表
	firmwareVersions, err := c.GetFirmwareVersionsByProductKey(productKey, tags)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(firmwareVersions)

	// 上传固件
	url, err := c.UploadFirmware(filePath)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(url)

	// 新增固件
	firmware, err := c.NewFirmware(&newFirmwareParam)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(firmware)

	// 新增固件版本
	firmwareVersion, err := c.NewFirmwareVersion(&newFirmwareVersionsParam)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(firmwareVersion)

	// 查看最新ota任务
	latestJob, err := c.GetOTALatestJob(sn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(latestJob)

	// 获取我的产品列表
	products, err := c.GetProductsOfMe()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(products)

	// 查看指定应用的product_key列表
	productKeys, err := c.GetAppProductKeys(appName)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(productKeys)

}

```




