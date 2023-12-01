# [weemesh/weemesh-go-sdk](https://github.com/weemesh/weemesh-go-sdk)

微麦石物联网平台SDK

## 功能
1. 获取指定产品的物模型
2. 获取指定设备的状态


## 使用方法

```go
package main

import (
	"fmt"
	"gitee.com/winstar-smart/weemesh-go-sdk"
	"log"
)

const (
	accessKey    = "xxxx"
	accessSecret = "xxxx"
	address      = "xxxx"
	productKey   = "xxxx"
	orgId        = "xxxx"
	sn = "xxxx"
)

func main() {
	c := weemesh.NewClient(accessKey, accessSecret, address)

	// 获取指定产品的物模型
	result, err := c.GetThingModel(productKey, orgId)

	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(result)
	
	// 获取指定设备的状态
	result, err = c.GetDeviceState(sn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
	
}

```


