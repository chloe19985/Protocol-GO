package main

import (
	//这里不是包名，而是相对路径名
	"fmt"
	"log"
	"os"
	"time"

	"./mcbin"
)

func main() {
	handler := mcbin.NewBINClientHandler("192.168.2.10:6000")
	handler.Timeout = 10 * time.Second
	//handler.Logger = log.New()
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer handler.Close()

	client := mcbin.NewClient(handler)

	results, err := client.ReadXCoils(1777,2 )
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("results:% X\n", results)
}
