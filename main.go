package main

import (
	"fmt"
	"net/http"
)

type CalcServer bool

func (a CalcServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	var k CalcServer
	http.ListenAndServe(":3333", k)
}
