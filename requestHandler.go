package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type MyRequest struct {
	request *http.Request
	payload string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is me Sachin!\n")
}
func getHello(w http.ResponseWriter, r *MyRequest, lb *LoadBalancer) {
	lb.sendRequest(r)
	io.WriteString(w, "Hello, HTTP!\n")
}

func lbHandler(lb *LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myreq := &MyRequest{
			request: r,
			payload: payloads[rand.Intn(len(payloads)-1)],
		}

		result := lb.sendRequest(myreq)
		io.WriteString(w, result)
	}
}
