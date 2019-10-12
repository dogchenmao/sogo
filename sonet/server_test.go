package sonet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	time.Sleep(5 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7734")
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}

	fmt.Println("ClientTest Dial...")
	for {
		bufw := []byte("nihao")
		conn.Write(bufw)
		buf := make([]byte, 500)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read Error...")
			return
		}
		fmt.Println(buf[:cnt])
		time.Sleep(5 * time.Second)
	}
}
func TestServer(t *testing.T) {
	fmt.Println("Test Start...")

	server := NewServer("te")

	go ClientTest()

	server.Serve()

}
