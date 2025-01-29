package controller

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Recv_Context struct {
	Frame []byte
	Err   error
}

type Controller_Instance struct {
	MulticastAddr string
	MulticastPort int
	UnicastAddr   string
	UnicastPort   int
	Conn          *net.UDPConn
	Recv_Channel  chan Recv_Context
}

const (
	DefaultMulticastAddr = "224.0.23.0"
	DefaultMulticastPort = 3610
	DefaultUnicastAddr   = "localhost"
	DefaultUnicastPort   = 3611
)

func (self *Controller_Instance) Send(frame []byte) error {
	address := self.MulticastAddr + ":" + strconv.Itoa(self.MulticastPort)
	fmt.Println("sending to ", address)
	conn, err := net.Dial("udp", address)
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

func (self *Controller_Instance) Read() ([]byte, error) {
	recv_context := <-self.Recv_Channel
	return recv_context.Frame, recv_context.Err
}

func (self *Controller_Instance) recvThread() {
	var buf = make([]byte, 2048)
	for {
		n, _, err := self.Conn.ReadFromUDP(buf)
		if err != nil {
			self.Recv_Channel <- Recv_Context{Frame: nil, Err: err}
			continue
		}
		self.Recv_Channel <- Recv_Context{Frame: buf[:n], Err: nil}
	}
}

/*
Setting device multicastAddr and multicastPort
default addr: 224.0.23.0
default port: 3610
Setting controller unicastAddr and unicastPort
default addr: localhost
default port: 3610
*/
func Start(multicastAddr string, multicastPort int, unicastAddr string, unicastPort int) (Controller_Instance, error) {
	var controller = Controller_Instance{MulticastAddr: multicastAddr, MulticastPort: multicastPort, Recv_Channel: make(chan Recv_Context)}
	address := net.UDPAddr{
		IP:   net.ParseIP(unicastAddr),
		Port: unicastPort,
	}
	conn_server, err := net.ListenUDP("udp", &address)
	if err != nil {
		address = net.UDPAddr{
			IP:   net.ParseIP(DefaultUnicastAddr),
			Port: DefaultUnicastPort,
		}
		conn_server, err = net.ListenUDP("udp", &address)
		if err != nil {
			return Controller_Instance{}, errors.New("net.ListenUDP error")
		} else {
			controller.UnicastAddr = DefaultUnicastAddr
			controller.UnicastPort = DefaultUnicastPort
		}
	} else {
		controller.UnicastAddr = unicastAddr
		controller.UnicastPort = unicastPort
	}
	controller.Conn = conn_server

	go func(controller Controller_Instance) {
		controller.recvThread()
	}(controller)

	return controller, nil
}
