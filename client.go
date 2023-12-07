package weemesh

import (
	"strconv"
	"sync"
	"time"
)

const (
	Version230503 = "2023-05-03"
	MethodSha1    = "sha1"
)

type Client struct {
	accessKey    string
	accessSecret string
	address      string
	sign         string
	timestamp    time.Time
	mu           sync.Mutex
}

func NewClient(accessKey, accessSecret, address string) *Client {
	return &Client{
		accessKey:    accessKey,
		accessSecret: accessSecret,
		address:      address,
	}
}

func (c *Client) signature() (string, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	// 判断是否需要重新生成签名
	if c.sign == "" || c.timestamp.Before(now) {
		et := now.Add(1 * time.Hour)
		timestamp := strconv.Itoa(int(et.Unix()))

		sign, err := generateSignatureWithTimestamp(c.accessKey, c.accessSecret, MethodSha1, Version230503, timestamp)
		if err != nil {
			return "", err
		}

		c.sign = sign
		c.timestamp = et
	}

	return c.sign, nil
}
