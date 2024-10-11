package controller

type Echonetlite struct {
	destIP   string
	destPort string
	format   []byte
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

func (self *Echonetlite) MakeFrame() {}
