package main

import (
	//这里不是包名，而是相对路径名
	"fmt"
	"time"

	"./mcbin"
)

func main() {
	handler := mcbin.NewBINClientHandler("192.168.0.120:1400")
	handler.Timeout = 10 * time.Second
	//handler.Logger = log.New()
	err := handler.Connect()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer handler.Close()

	client := mcbin.NewClient(handler)

	results, err := client.ReadXCoils(10, 2)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("results:% x\n", results)
}
