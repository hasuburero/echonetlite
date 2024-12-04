package echonetlite

import (
	"errors"
	"unsafe"
)

const (
	shift_byte = 8
	bytesize   = 256
)

func ByteSize(arg uint) int {
	var i int = 0
	for ; i < 4; i++ {
		arg = arg / bytesize
		if arg == 0 {
			break
		}
	}

	return i + 1
}

func Byte2Int(arg1 []byte, arg2 int) (int, error) {
	var int_size int = int(unsafe.Sizeof(int(0)))
	if arg2 > int_size || arg2 <= 0 {
		err := errors.New("arg2 is out of size int")
		return 0, err
	} else if len(arg1) > int_size {
		err := errors.New("size of arg1 is out of size int")
		return 0, err
	}
	var result int = 0
	for i := range arg2 {
		result += (int(arg1[i]) << (shift_byte * (arg2 - 1 - i)))
	}

	return result, nil
}

func Int2Byte(arg1 int, arg2 int) ([]byte, error) {
	var int_size int = int(unsafe.Sizeof(int(0)))
	if arg2 > int_size || arg2 <= 0 {
		err := errors.New("arg2 is out of size int")
		return nil, err
	}
	result := make([]byte, arg2)
	for i := range arg2 {
		result[arg2-1-i] = byte(arg1 & 0xff)
		arg1 = arg1 >> shift_byte
	}

	return result, nil
}
