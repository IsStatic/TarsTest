package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Println(w, "hello")
}

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &MyHandler{},
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")

}
