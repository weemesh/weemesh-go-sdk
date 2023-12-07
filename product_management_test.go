package weemesh

import (
	"fmt"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func TestProductManagement(t *testing.T) {
	TestingT(t)
}

type ProductManagementSuite struct{}

var _ = Suite(&ProductManagementSuite{})

func (s *ProductManagementSuite) TestGetProductsOfMe(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetProductsOfMe()
	if err != nil {
		c.Fatal(err)
	}

	fmt.Println(result)
	c.Log(result)

	c.Assert(result.Code, Equals, http.StatusOK)
}
