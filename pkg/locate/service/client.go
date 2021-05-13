package service

import (
	"log"
	"net/rpc"
	"time"

	"github.com/1000Delta/wifi-locate/pkg/common/utils"
	"github.com/1000Delta/wifi-locate/pkg/locate"
)

const (
	initReconnectDelay time.Duration = 1
	maxReconnectDelay  time.Duration = 4
	connectTimeout     time.Duration = 20

	listenAddress = "wifi-locate-locate:52201"
)

type Client struct {
	conn *rpc.Client
}

// Close client rpc connection
func (c Client) Close() error {
	return c.conn.Close()
}

// Locate your location by the wifi scan list
func (c Client) Locate(req LocateReq, resp *locate.LocationInfo) error {
	if err := c.conn.Call("LocateService.Locate", req, resp); err != nil {
		return err
	}

	return nil
}

func (c Client) Collect(req CollectReq, resp *CollectResp) error {
	if err := c.conn.Call("LocateService.Collect", req, resp); err != nil {
		return err
	}

	return nil
}

func (c Client) CreateMap(req CreateMapReq, resp *CreateMapResp) error {
	if err := c.conn.Call("LocateService.CreateMap", req, resp); err != nil {
		return err
	}

	return nil
}

// DefaultClient to process rpc call
func DefaultClient() *Client {
	var client *Client
	// 超时时间内尝试建立连接
	err := utils.CallUntilNoErrorWithTimeout(
		func(currentDelay time.Duration) error {
			c, err := rpc.DialHTTP("tcp", listenAddress)
			if err != nil {
				log.Print(err) // 此处不能中止，否则会影响外部服务
				return err
			}
			client = &Client{c}
			return nil
		},
		initReconnectDelay, maxReconnectDelay, connectTimeout,
	)
	if err != nil {
		return nil
	}

	return client
}

// DefaultClient to process rpc call
func NewClient(address string) *Client {
	var client *Client
	// 超时时间内尝试建立连接
	err := utils.CallUntilNoErrorWithTimeout(
		func(currentDelay time.Duration) error {
			c, err := rpc.DialHTTP("tcp", address)
			if err != nil {
				log.Print(err) // 此处不能中止，否则会影响外部服务
				return err
			}
			client = &Client{c}
			return nil
		},
		initReconnectDelay, maxReconnectDelay, connectTimeout,
	)
	if err != nil {
		return nil
	}

	return client
}
