package server

import (
	"net"
)

const (
	maxsize = 1024
)

type ServerInstance struct {
	ip   string
	port string
}

func InitInstance(ip string, port string) ServerInstance {
	servinst := ServerInstance{ip: ip, port: port}
	return servinst
}

func echonetHandler(conn *net.TCPConn) {
	defer conn.Close()
	buf := make([]byte, maxsize)

	// first greeting per each connection
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	gw_id := string(buf[:n])
	fmt.Println(gw_id)

	// echonetlite bridge default routine
	for {
		_, err = conn.Write([]byte("Hello world!!\n"))
		if err != nil {
			return
		}
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
	}
}

func receiveTCPConnection(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(conn *net.TCPConn) {
			echonetHandler(conn)
		}(conn)
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
