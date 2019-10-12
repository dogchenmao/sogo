package main

import (
	"fmt"
	"time"

	"github.com/dogchenmao/sogo/soclient"
	"github.com/dogchenmao/sogo/sonet"
)

func main() {
	fmt.Println("Test Start...")

	soclient.NewClient("", 1234)

	server := sonet.NewServer("te")
	server.Serve()

	for {
		time.Sleep(10 * time.Second)
	}
}
