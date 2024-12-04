package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
)

const (
	bytesize = 256
)

func main() {
	var a, b, c int = 100, 1000, 256*256 - 1
	fmt.Printf("%d:%d, %d:%d, %d:%d\n", a, echonetlite.ByteSize(uint(a)), b, echonetlite.ByteSize(uint(b)), c, echonetlite.ByteSize(uint(c)))

	size := echonetlite.ByteSize(uint(a))
	buf, err := echonetlite.Int2Byte(a, size)
	if err != nil {
		fmt.Println(err)
		fmt.Println("echonetlite.Int2Byte error")
		return
	}
	fmt.Printf("%d bytes\n", size)
	fmt.Printf("%d:%v\n", a, buf)

	size = echonetlite.ByteSize(uint(b))
	buf, err = echonetlite.Int2Byte(b, size)
	if err != nil {
		fmt.Println(err)
		fmt.Println("echonetlite.Int2Byte error")
		return
	}
	fmt.Printf("%d bytes\n", size)
	fmt.Printf("%d:%v\n", b, buf)

	size = echonetlite.ByteSize(uint(c))
	buf, err = echonetlite.Int2Byte(c, echonetlite.ByteSize(uint(c)))
	if err != nil {
		fmt.Println(err)
		fmt.Println("echonetlite.Int2Byte error")
		return
	}
	fmt.Printf("%d bytes\n", size)
	fmt.Printf("%d:%v\n", c, buf)

	return
}
