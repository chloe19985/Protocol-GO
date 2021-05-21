package main

import (
	//这里不是包名，而是相对路径名
	"fmt"
	"log"
	"os"
	"time"

	"./mc3e"
)

func main() {
	handler := mc3e.NewASCClientHandler("192.168.2.10:6000")
	handler.Timeout = 10 * time.Second
	//handler.Logger = log.New()
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer handler.Close()

	client := mc3e.NewClient(handler)

	results, err := client.ReadXCoils(6, 3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("results:% X\n", results)
}
