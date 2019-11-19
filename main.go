package main

import (
	"fmt"
	"time"

	"github.com/dogchenmao/sogo/sonet"
)

func main() {
	fmt.Println("Test Start...")

	client := sonet.NewConnect("127.0.0.1", 42002)

	go client.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}
