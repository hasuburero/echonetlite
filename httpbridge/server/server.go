package server

import (
	"fmt"
	"net/http"
)

const (
	addr = ""
	port = "8080"
)

func Handler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	wait := make(chan bool)
	server := http.Server{
		Addr: addr + ":" + port,
	}

	http.HandleFunc("/", Handler)
	<-wait
}
