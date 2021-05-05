package service

import (
	"log"
	"net/rpc"

	"github.com/1000Delta/wifi-locate/pkg/locate"
)

var (
	listenAddress = ":52201"
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
	if err := c.conn.Call("Locate.Locate", req, resp); err != nil {
		return err
	}

	return nil
}

func (c Client) Collect(req CollectReq, resp *CollectResp) error {
	if err := c.conn.Call("Locate.Collect", req, resp); err != nil {
		return err
	}

	return nil
}

func (c Client) CreateMap(req CreateMapReq, resp *CreateMapResp) error {
	if err := c.conn.Call("Locate.CreateMap", req, resp); err != nil {
		return err
	}

	return nil
}

// NewClient to process rpc call
func NewClient() *Client {
	client, err := rpc.DialHTTP("tcp", listenAddress)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{client}
}
