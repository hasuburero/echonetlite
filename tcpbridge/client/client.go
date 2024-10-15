package client

import (
	"fmt"
	"net"
)

const (
	host  = "localhost:8080" //
	gw_id = "1"              // echonetliteとは関係ない実用上の識別id
)

type Bridge struct {
	addr  string
	port  string
	gw_id string
	conn  net.Conn
}

func MakeBridgeInstance(addr string, port string, gw_id string) (Bridge, error) {
	bridge := Bridge{addr: addr, port: port, gw_id: gw_id}
	conn, err := net.Dial("tcp", bridge.addr+":"+bridge.port)
	net.Conn
	if err != nil {
		return Bridge{}, err
	}
	bridge.conn = conn

	_, err = bridge.conn.Write([]byte(bridge.gw_id))
	if err != nil {
		return Bridge{}, err
	}

	return bridge, err
}

func main() {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Dial error")
		return
	}

	_, err = conn.Write([]byte(gw_id))
	if err != nil {
		fmt.Println(err)
		fmt.Println("conn.Write error")
		return
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}

}
