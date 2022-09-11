package main

import (
	"bufio"
	"fmt"
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
	scanner := bufio.NewScanner(c.in)
	for scanner.Scan() {
		_, err = c.conn.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return
}

func (c *Client) Receive() (err error) {
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		_, err = c.out.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			return err
		}
	}
	return
}

func (c *Client) Close() (err error) {
	return c.conn.Close()
}
