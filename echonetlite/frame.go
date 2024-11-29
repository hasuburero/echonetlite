package echonetlite

import (
	"errors"
)

type ClassDefinition struct {
}

var (
	Class_SmartMeter = [3]byte{0x02, 0x88, 0x01}
	Class_Battery    = [3]byte{0x02, 0x7d, 0x01}
)

const (
	EHD1 = byte(0x10)
	EHD2 = byte(0x81)
)

const (
	// request
	ESV_SetI    = byte(0x60) // write + no resopnse
	ESV_SetC    = byte(0x61) // write + response
	ESV_Get     = byte(0x62) // read
	ESV_INF_REQ = byte(0x63) // push (device send multicast packet)
	ESV_SetGet  = byte(0x6e) // write + read

	// response
	ESV_Set_Res    = byte(0x71) // SetC response
	ESV_Get_Res    = byte(0x72) // Get response
	ESV_INF        = byte(0x73) // push multicast packet
	ESV_INFC       = byte(0x74) // push response
	ESV_INFC_Res   = byte(0x7a) // push response
	ESV_SetGet_Res = byte(0x7f) // SetGet response
)

type Echonetlite struct {
	Frame      []byte
	Frame_size int
	ehd1       byte      // Echonet lite frame header1
	ehd2       byte      // Echonet lite frame header2
	tid        [2]byte   // transaction id
	seoj       [3]byte   // source echonetlite object
	deoj       [3]byte   // destination echonetlite object
	esv        byte      // echonetlite service
	opc        byte      // number of executing property
	datactx    []Datactx // property info
}

type Datactx struct {
	epc byte   // echonetlite property
	pdc byte   // EDT Bytes
	edt []byte // property value
}

func (self *Echonetlite) MakeFrame() error {
	if int(self.opc) != len(self.datactx) {
		return errors.New("opc not matches for datactx length")
	}
	var frame []byte
	frame = append(frame, self.ehd1)
	frame = append(frame, self.ehd2)
	frame = append(frame, self.tid[:]...)
	frame = append(frame, self.seoj[:]...)
	frame = append(frame, self.deoj[:]...)
	frame = append(frame, self.esv)
	frame = append(frame, self.opc)
	for i := 0; i < int(self.opc); i++ {
		if int(self.datactx[i].pdc) != len(self.datactx) {
			return errors.New("pdc don't match for edt length")
		}
		frame = append(frame, self.datactx[i].epc)
		frame = append(frame, self.datactx[i].pdc)
		frame = append(frame, self.datactx[i].edt...)
	}
	return nil
}

func MakeInstance(ehd1, ehd2 byte, tid [2]byte, seoj, deoj [3]byte) Echonetlite {
	echonetlite := Echonetlite{ehd1: ehd1, ehd2: ehd2, tid: tid, seoj: seoj, deoj: deoj}
	return echonetlite
}
