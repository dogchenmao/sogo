package main

import (
	"fmt"
	"time"

	"github.com/dogchenmao/sogo/tp-group/robot"
)

func main() {
	fmt.Println("Test Start...")

	robot := robot.NewRobotMgr()
	robot.Init()

	for {
		time.Sleep(10 * time.Second)
	}
}
