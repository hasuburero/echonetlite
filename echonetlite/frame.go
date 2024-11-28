package echonetlite

import (
	"errors"
)

type Echonetlite struct {
	Component  Component
	Frame      []byte
	Frame_size int
}

type Component struct {
	destIP   string
	destPort string
	ehd1     byte
	ehd2     byte
	tid      [2]byte
	seoj     [3]byte
	deoj     [3]byte
	esv      byte
	opc      byte
	datactx  []Datactx
}

type Datactx struct {
	epc byte
	pdc byte
	edt []byte
}

func (self *Echonetlite) MakeFrame() error {
	if int(self.component.opc) != len(self.component.datactx) {
		return errors.New("opc not matches for datactx length")
	}
	var frame []byte
	frame = append(frame, self.component.ehd1)
	frame = append(frame, self.component.ehd2)
	frame = append(frame, self.component.tid[:]...)
	frame = append(frame, self.component.seoj[:]...)
	frame = append(frame, self.component.deoj[:]...)
	frame = append(frame, self.component.esv)
	frame = append(frame, self.component.opc)
	for i := 0; i < int(self.component.opc); i++ {
		if int(self.component.datactx[i].pdc) != len(self.component.datactx) {
			return errors.New("pdc don't match for edt length")
		}
		frame = append(frame, self.component.datactx[i].epc)
		frame = append(frame, self.component.datactx[i].pdc)
		frame = append(frame, self.component.datactx[i].edt...)
	}
	return nil
}

func MakeInstance(destIP, destPort string, ehd1, ehd2 byte, tid [2]byte, seoj, deoj [3]byte) Echonetlite {
	component := Component{destIP: destIP, destPort: destPort, ehd1: ehd1, ehd2: ehd2, tid: tid, seoj: seoj, deoj: deoj}
	return Echonetlite{component: component}
}
