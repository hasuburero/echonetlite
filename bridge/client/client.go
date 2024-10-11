package client

import (
  "net"
  "fmt"
)

const (
  host = "localhost:8080" // 
  gw_id = "1" // echonetliteとは関係ない実用上の識別id
)

func main(){
  conn, err := net.Dial("tcp", host)
  if err != nil{
    fmt.Println("Dial error")
    return
  }

  output := "1000"
  _, err = conn.Write([]byte(output))
  if err != nil{
    fmt.Println(err)
    fmt.Println(conn.Write error)
    return
  }
  buf 
}
