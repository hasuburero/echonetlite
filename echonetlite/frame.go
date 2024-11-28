package echonetlite

import (
	"errors"
)

type Echonetlite struct {
	Compo      Component
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
	if int(self.Compo.opc) != len(self.Compo.datactx) {
		return errors.New("opc not matches for datactx length")
	}
	var frame []byte
	frame = append(frame, self.Compo.ehd1)
	frame = append(frame, self.Compo.ehd2)
	frame = append(frame, self.Compo.tid[:]...)
	frame = append(frame, self.Compo.seoj[:]...)
	frame = append(frame, self.Compo.deoj[:]...)
	frame = append(frame, self.Compo.esv)
	frame = append(frame, self.Compo.opc)
	for i := 0; i < int(self.Compo.opc); i++ {
		if int(self.Compo.datactx[i].pdc) != len(self.Compo.datactx) {
			return errors.New("pdc don't match for edt length")
		}
		frame = append(frame, self.Compo.datactx[i].epc)
		frame = append(frame, self.Compo.datactx[i].pdc)
		frame = append(frame, self.Compo.datactx[i].edt...)
	}
	return nil
}

func MakeInstance(destIP, destPort string, ehd1, ehd2 byte, tid [2]byte, seoj, deoj [3]byte) Echonetlite {
	component := Component{destIP: destIP, destPort: destPort, ehd1: ehd1, ehd2: ehd2, tid: tid, seoj: seoj, deoj: deoj}
	return Echonetlite{Compo: component}
}
