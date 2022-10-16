package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Client struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		conn:    nil,
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *Client) Connect() (err error) {
	conn, err := net.DialTimeout("tcp", c.address, timeout)
	if err != nil {
		return
	}
	c.conn = conn
	return
}

func (c *Client) Send() (err error) {
	_, err = io.Copy(c.conn, c.in)
	return
}

func (c *Client) Receive() (err error) {
	_, err = io.Copy(c.out, c.conn)
	return
}

func (c *Client) Close() (err error) {
	if c.conn == nil {
		return
	}
	err = c.conn.Close()
	c.conn = nil
	return
}
