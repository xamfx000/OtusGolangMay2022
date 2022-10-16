/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
)

var (
	host    string
	port    string
	timeout time.Duration
	err     error
)

func init() {
	flag.StringVar(&host, "host", "localhost", "destination address")
	flag.StringVar(&port, "port", "4242", "port")
	flag.DurationVar(&timeout, "timeout", time.Duration(50)*time.Second, "connection timeout")
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	signal.NotifyContext(ctx, syscall.SIGINT)
	client := NewTelnetClient(host+":"+port, timeout, os.Stdin, os.Stdout)
	err = client.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		err = client.Receive()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	go func() {
		err = client.Send()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	<-ctx.Done()
	cancel()
}
