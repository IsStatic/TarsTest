package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		fmt.Println(name)
		handler(w, r)
	}
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080,"}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
