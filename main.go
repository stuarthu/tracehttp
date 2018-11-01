package main

import (
	"./proxy"
	"net/http"
)

func main() {
	s := &proxy.Server{}
	http.ListenAndServe(":8080", s)
}
