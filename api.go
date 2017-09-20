package template

/*
The api.go defines the methods that can be called from the outside. Most
of the methods will take a roster so that the service knows which nodes
it should work with.

This part of the service runs on the client or the app.
*/

import (
	"gopkg.in/dedis/onet.v1"
	"gopkg.in/dedis/onet.v1/log"
	"gopkg.in/dedis/onet.v1/network"
)

// ServiceName is used for registration on the onet.
const ServiceName = "Decenarch"

// Client is a structure to communicate with the CoSi
// service
type Client struct {
	*onet.Client
}

// NewClient instantiates a new cosi.Client
func NewClient() *Client {
	return &Client{Client: onet.NewClient(ServiceName)}
}

// Clock will return the time in seconds it took to run the protocol.
func (c *Client) Clock(r *onet.Roster) (*ClockResponse, onet.ClientError) {
	dst := r.RandomServerIdentity()
	log.Lvl4("Sending message to", dst)
	reply := &ClockResponse{}
	err := c.SendProtobuf(dst, &ClockRequest{r}, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Count will return the number of times `Clock` has been called on this
// service-node.
func (c *Client) Count(si *network.ServerIdentity) (int, error) {
	reply := &CountResponse{}
	err := c.SendProtobuf(si, &CountRequest{}, reply)
	if err != nil {
		return -1, err
	}
	return reply.Count, nil
}

// Save will record the website requested in the conodes
func (c *Client) Save(r *onet.Roster, url string) onet.ClientError {
	dst = r.RandomServerIdentity()
	log.Lvl4("Sending message to", dst)
	resp := &SaveResponse{}
	err := c.SendProtobuf(dst, &SaveRequest{url}, resp)
	return err
}

// Retrieve will send the website requested to the client
func (c *Client) Retrieve(si *network.ServerIdentity, url string) (*RetrieveResponse, onet.ClientError) {
	resp := &RetrieveResponse{}
	err := c.SendProtobuf(si, &RetrieveRequest{url}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
