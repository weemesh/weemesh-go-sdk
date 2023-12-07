package weemesh

import (
	"flag"
	"fmt"
	"gitee.com/winstar-smart/weemesh-go-sdk/model"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func TestOtaManagement(t *testing.T) {
	TestingT(t)
}

type OtaManagementSuite struct{}

var _ = Suite(&OtaManagementSuite{})

var filePath = flag.String("filePath", "", "bin文件路径")

func (s *OtaManagementSuite) TestGetFirmwareVersionsByProductKey(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" || *productKey == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetFirmwareVersionsByProductKey(*productKey, []string{"release"})
	if err != nil {
		c.Fatal(err)
	}

	fmt.Println(result)

	c.Assert(result.Code, Equals, http.StatusOK)
}

func (s *OtaManagementSuite) TestNewFirmware(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.NewFirmware(&model.NewFirmwareParams{
		ProductKey:   "1",
		FirmwareName: "测试固件3",
		Version:      "v1",
		Remark:       "test",
		FirmwareType: 1,
		URL:          "xxx",
		SignMethod:   0,
		Enable:       true,
		Tags:         []string{"test"},
	})
	if err != nil {
		c.Fatal(err)
	}

	fmt.Println(result)

	c.Assert(result.Code, Equals, http.StatusOK)
}

func (s *OtaManagementSuite) TestUploadFirmware(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.UploadFirmware(*filePath)
	if err != nil {
		c.Fatal(err)
	}

	fmt.Println(result)

	c.Assert(result.Code, Equals, http.StatusOK)
}
