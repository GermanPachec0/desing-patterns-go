package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lightning connector into computer")
	com.InsertIntoLightningPort()
}
