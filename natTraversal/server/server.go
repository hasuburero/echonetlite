package server

import (
	"net"
)

type ServerInstance struct {
	ip   string
	port string
}

func InitInstance(ip string, port string) ServerInstance {
	servinst := ServerInstance{ip: ip, port: port}
	return servinst
}

func receiveTCPConnection(listener *net.TCPListener){
  for{
    conn, err := listener.AcceptTCP()
    if err != nil{
      fmt.Println(err)
      continue
    }
    go func(conn *net.TCPConn){
      
    }
  }
}

func (self *ServerInstance) InitServer() error {
	tcpaddr, err := net.ResolveTCPAddr("tcp", self.ip+":"+self.port)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		return err
	}

	go func() {
		receiveTCPConnection(listener)
	}()

}
