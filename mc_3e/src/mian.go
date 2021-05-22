package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./mc3e"
)

func main() {
	handler := mc3e.NewASCClientHandler("192.168.2.10:6001")
	handler.Timeout = 10 * time.Second
	//handler.Logger = log.New()
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer handler.Close()

	client := mc3e.NewClient(handler)

	results, err := client.ReadXCoils(05,1)
	//results2,err := client.ReadDRegisters(4,1)
	//results3,err := client.ReadDRegisters(8,1)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("results:% X\n", results)
	//fmt.Printf("results2:% X\n", results2)
	//fmt.Printf("results3:% x\n", results3)
}
