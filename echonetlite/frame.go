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
	EHD1       byte      // Echonet lite frame header1
	EHD2       byte      // Echonet lite frame header2
	Tid        [2]byte   // transaction id
	SEOJ       [3]byte   // source echonetlite object
	DEOJ       [3]byte   // destination echonetlite object
	ESV        byte      // echonetlite service
	OPC        byte      // number of executing property
	Datactx    []Datactx // property info
}

type Datactx struct {
	EPC byte   // echonetlite property
	PDC byte   // EDT Bytes
	EDT []byte // property value
}

func Tidinc(tid [2]byte) [2]byte {
	var int_buf int = 0
	int_buf = int(tid[0]) << 8
	int_buf += int(tid[1])
	int_buf += 1
	tid[1] = byte(int_buf)
	tid[0] = byte(int_buf >> 8)
	return tid
}

func (self *Echonetlite) MakeFrame() error {
	if int(self.OPC) != len(self.Datactx) {
		return errors.New("opc not matches for datactx length")
	}
	var frame []byte
	frame = append(frame, self.EHD1)
	frame = append(frame, self.EHD2)
	frame = append(frame, self.Tid[:]...)
	frame = append(frame, self.SEOJ[:]...)
	frame = append(frame, self.DEOJ[:]...)
	frame = append(frame, self.ESV)
	frame = append(frame, self.OPC)
	for i := 0; i < int(self.OPC); i++ {
		if int(self.Datactx[i].PDC) != len(self.Datactx[i].EDT) {
			return errors.New("pdc don't match for edt length")
		}
		frame = append(frame, self.Datactx[i].EPC)
		frame = append(frame, self.Datactx[i].PDC)
		frame = append(frame, self.Datactx[i].EDT...)
	}
	self.Frame = frame
	self.Frame_size = len(frame)
	return nil
}

func MakeInstance(ehd1, ehd2 byte, tid [2]byte, seoj, deoj [3]byte, esv byte, opc byte, datactx []Datactx) Echonetlite {
	echonetlite := Echonetlite{EHD1: ehd1, EHD2: ehd2, Tid: tid, SEOJ: seoj, DEOJ: deoj, ESV: esv, OPC: opc, Datactx: datactx}
	return echonetlite
}
