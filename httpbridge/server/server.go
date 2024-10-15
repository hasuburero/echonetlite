package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	addr = ""
	port = "8080"
)

type Get_echonet_request struct {
	gw_id  string `json:"gw_id"`
	format string `json:"format"`
}

type Get_echonet_response struct {
	format string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	reqBody := make([]byte, length)
	var ctx Get_echonet_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		return
	}
	fmt.Println(ctx)

	var format Get_echonet_response{format:"1234657890"}
}

func main() {
	wait := make(chan bool)
	server := http.Server{
		Addr: addr + ":" + port,
	}

	http.HandleFunc("/", Handler)
	<-wait
}
