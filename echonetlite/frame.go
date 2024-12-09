package echonetlite

import (
	"errors"
	"fmt"
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

func CalcByte(frame []byte) {
}

func (self *Echonetlite) ShowInstanceFrame() {
	length := 12
	fmt.Printf("EHD1:%x EHD2:%x ", self.EHD1, self.EHD2)
	fmt.Printf("Tid:%x%x ", self.Tid[0], self.Tid[1])
	fmt.Printf("SEOJ:%x%x%x ", self.SEOJ[0], self.SEOJ[1], self.SEOJ[2])
	fmt.Printf("DEOJ:%x%x%x ", self.DEOJ[0], self.DEOJ[1], self.DEOJ[2])
	fmt.Printf("ESV:%x OPC:%x ", self.ESV, self.OPC)
	for _, ctx := range self.Datactx {
		fmt.Printf("EPC:%x PDC:%x ", ctx.EPC, ctx.PDC)
		length += 2
		if ctx.PDC == 0 {
			continue
		} else {
			fmt.Printf("EDT:")
			for _, ctx := range ctx.EDT {
				fmt.Printf("%x", ctx)
				length++
			}
			fmt.Printf(" ")
		}
	}
	fmt.Println("")
	fmt.Println(length)
}

func ShowByteFrame(frame []byte) {
	fmt.Printf("EHD1:%x EHD2:%x ", frame[0], frame[1])
	fmt.Printf("Tid:%x%x ", frame[2], frame[3])
	fmt.Printf("SEOJ:%x%x%x ", frame[4], frame[5], frame[6])
	fmt.Printf("DEOJ:%x%x%x ", frame[7], frame[8], frame[9])
	fmt.Printf("ESV:%x OPC:%x ", frame[10], frame[11])
	index := 0
	for _ = range int(frame[11]) {
		fmt.Printf("EPC:%x PDC:%x ", frame[12+index], frame[12+index+1])
		index += 2
		if frame[12+index-1] == 0 {
			continue
		} else {
			fmt.Printf("EDT:")
			i := 0
			for _ = range int(frame[index-1]) {
				fmt.Printf("%x", frame[index+i])
				i++
			}
			fmt.Printf(" ")
			index += i
		}
	}
	fmt.Println("")
}

func (self *Echonetlite) ReverseFrame() error {
	if len(self.Frame) == 0 {
		return errors.New("0 length frame")
	}
	index := 0
	self.Frame_size = len(self.Frame)
	self.EHD1 = self.Frame[0]
	self.EHD2 = self.Frame[1]
	self.Tid = [2]byte{self.Frame[2], self.Frame[3]}
	self.SEOJ = [3]byte{self.Frame[4], self.Frame[5], self.Frame[6]}
	self.DEOJ = [3]byte{self.Frame[7], self.Frame[8], self.Frame[9]}
	self.ESV = self.Frame[10]
	self.OPC = self.Frame[11]
	index += 12
	if self.OPC == 0 {
		return nil
	}

	for i := 12; i < self.Frame_size; {
		var datactx Datactx = Datactx{EPC: self.Frame[i], PDC: self.Frame[i+1]}
		index += 2
		for j := range int(datactx.PDC) {
			datactx.EDT = append(datactx.EDT, self.Frame[i+2+j])
			index += 1
		}
		i += 2 + int(datactx.PDC)
		self.Datactx = append(self.Datactx, datactx)
	}
	self.Frame_size = index
	return nil
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
	index := 12

	//wip
	if self.ESV == ESV_Get {
		frame = append(frame, self.Datactx[i].EPC)
		frame = append(frame, self.Datactx[i].PDC)
	}
	for i := 0; i < int(self.OPC); i++ {
		frame = append(frame, self.Datactx[i].EPC)
		frame = append(frame, self.Datactx[i].PDC)
		index += 2
		if int(self.Datactx[i].PDC) != len(self.Datactx[i].EDT) {
			return errors.New("pdc don't match for edt length")
		} else if self.Datactx[i].PDC == 0 {
			continue
		} else {
			index += int(self.Datactx[i].PDC)
			frame = append(frame, self.Datactx[i].EDT...)
		}
	}
	self.Frame = frame
	self.Frame_size = len(frame)
	fmt.Println(self.Frame_size)
	self.Frame_size = index
	fmt.Println(self.Frame_size)
	return nil
}

func MakeInstance(frame []byte) Echonetlite {
	var echonetlite_instance Echonetlite
	echonetlite_instance = Echonetlite{EHD1: frame[0], EHD2: frame[1], Tid: [2]byte(frame[2:4]), SEOJ: [3]byte(frame[4:7]), DEOJ: [3]byte(frame[7:10]), ESV: frame[10], OPC: frame[11]}
	index := 12
	for index < len(frame) {
		var datactx Datactx = Datactx{EPC: frame[index], PDC: frame[index+1]}
		switch echonetlite_instance.ESV {
		case ESV_Get:
			index += 2
		case ESV_SetC:
			datactx.EDT = frame[index+2 : index+2+int(datactx.PDC)]
			index += 2 + int(datactx.PDC)
		case ESV_Get_Res:
			if datactx.PDC != 0 {
				datactx.EDT = frame[index+2 : index+2+int(datactx.PDC)]
				index += 2 + int(datactx.PDC)
			} else {
				index += 2
			}
		case ESV_Set_Res:
			if datactx.PDC != 0 {
				datactx.EDT = frame[index+2 : index+2+int(datactx.PDC)]
				index += 2 + int(datactx.PDC)
			} else {
				index += 2
			}
		default:
		}
		echonetlite_instance.Datactx = append(echonetlite_instance.Datactx, datactx)
	}
	echonetlite_instance.Frame_size = index
	return echonetlite_instance
}
