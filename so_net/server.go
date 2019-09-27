package so_net

import (
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      string
}

func (s *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%s", s.IP, s.Port))
		if err != nil {
			fmt.Sprintln("ResolveIPAddr err")
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Sprintln("ListenTCP err")
			return
		}

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Sprintln("AcceptTCP err")
				continue
			}

			go func() {
				for {
					buf := make([]byte, 50)
					cnt, err := conn.Read(buf)
					if err != nil {
						continue
					}

					if _, err = conn.Write(buf[:cnt]); err != nil {
						fmt.Sprintln("AcceptTCP err")
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("Stop server %s", s.Name)
}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) {

}
