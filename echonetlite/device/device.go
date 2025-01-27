package device

import (
	"fmt"
	"net"
	"strconv"
)

type Recv_Context struct {
	IP    net.IP
	Frame []byte
	Err   error
}

type Device_Instance struct {
	MulticastAddr string
	MulticastPort int
	UnicastPort   int
	Conn          *net.UDPConn
	Recv_Channel  chan Recv_Context
}

const (
	DefaultMulticastAddr = "224.0.23.0"
	DefaultMulticastPort = 3610
	DefaultUnicastPort   = 3610
)

func (self *Device_Instance) Send(frame []byte, dst net.IP) error {
	address := net.UDPAddr{
		IP:   dst,
		Port: self.UnicastPort,
	}
	conn, err := net.Dial("udp", address.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(frame)
	if err != nil {
		return err
	}

	return nil
}

func (self *Device_Instance) Read() (net.IP, []byte, error) {
	recv_context := <-self.Recv_Channel
	return recv_context.IP, recv_context.Frame, recv_context.Err
}

func (self *Device_Instance) recvThread() {
	var buf = make([]byte, 2048)
	for {
		n, remoteAddr, err := self.Conn.ReadFromUDP(buf)
		if err != nil {
			self.Recv_Channel <- Recv_Context{Frame: nil, Err: err}
			continue
		}
		self.Recv_Channel <- Recv_Context{IP: remoteAddr.IP, Frame: buf[:n], Err: err}
	}
}

/*
Setting device multicastAddr and multicastPort
default addr: 224.0.23.0
default port: 3610
Setting controller unicastAddr and unicastPort
default port: 3610
*/
func Start(multicastaddr string, multicastport, unicastport int) (Device_Instance, error) {
	var device = Device_Instance{MulticastAddr: multicastaddr, MulticastPort: multicastport, UnicastPort: unicastport, Recv_Channel: make(chan Recv_Context, 10)}

	address := multicastaddr + ":" + strconv.Itoa(multicastport)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		fmt.Println("net.ResolveUDPAddr error")
		fmt.Println("using default settings")
		address = DefaultMulticastAddr + ":" + strconv.Itoa(DefaultMulticastPort)
		addr, err = net.ResolveUDPAddr("udp", address)
		if err != nil {
			fmt.Println("net.ListenMulticastUDP error")
			return Device_Instance{}, err
		} else {
			device.MulticastAddr = DefaultMulticastAddr
			device.MulticastPort = DefaultMulticastPort
		}
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("net.ListenMulticastUDP error")
		return Device_Instance{}, err
	}
	device.Conn = conn

	go func(device Device_Instance) {
		device.recvThread()
	}(device)

	return device, nil
}
