package sonet

// import (
// 	"fmt"
// 	"net"
// 	"time"

// 	"github.com/dogchenmao/sogo/soiface"
// )

// type Server struct {
// 	Name      string
// 	IPVersion string
// 	IP        string
// 	Port      int
// }

// func CallBackClientConnect(con *net.TCPConn, buf []byte, cnt int) error {
// 	fmt.Println("CallBackClientConnect")
// 	str := []byte("CallBackClientConnect: ")
// 	_, err := con.Write(append(str, buf...))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *Server) Start() {
// 	fmt.Println("[START] Server listenner addr:", fmt.Sprintf("%s:%d", s.IP, s.Port))
// 	go func() {
// 		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
// 		if err != nil {
// 			fmt.Println("ResolveIPAddr err")
// 			return
// 		}

// 		listenner, err := net.ListenTCP(s.IPVersion, addr)
// 		if err != nil {
// 			fmt.Println("ListenTCP err", err)
// 			return
// 		}

// 		fmt.Println("start Sogo server ", s.Name, "success, now listenning...")

// 		var sessionID uint32 = 1
// 		for {
// 			conn, err := listenner.AcceptTCP()
// 			if err != nil {
// 				fmt.Println("AcceptTCP err")
// 				continue
// 			}

// 			dealConnect := NewConnect(conn, sessionID, CallBackClientConnect, nil)
// 			sessionID++

// 			go dealConnect.Start()
// 		}
// 	}()
// }

// func (s *Server) Stop() {
// 	fmt.Println("Stop server ", s.Name)
// }

// func (s *Server) Serve() {
// 	s.Start()
// 	for {
// 		time.Sleep(10 * time.Second)
// 	}
// }

// func NewServer(name string) soiface.IServer {
// 	st := &Server{
// 		Name:      name,
// 		IPVersion: "tcp4",
// 		IP:        "127.0.0.1",
// 		Port:      7734,
// 	}
// 	return st
// }
