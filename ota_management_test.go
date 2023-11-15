package weemesh

import (
	"fmt"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func TestOtaManagement(t *testing.T) {
	TestingT(t)
}

type OtaManagementSuite struct{}

var _ = Suite(&OtaManagementSuite{})

func (s *OtaManagementSuite) TestGetFirmwareVersionsByProductKey(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" || *productKey == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetFirmwareVersionsByProductKey(*productKey)
	if err != nil {
		c.Fatal(err)
	}

	fmt.Println(result)

	c.Assert(result.Code, Equals, http.StatusOK)
}
