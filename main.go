package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

var payloads []string = []string{"Sachin", "malik", "nitin", "jatin", "ekta"}

func main() {
	lb := initLoadBalancer()
	done := make(chan bool)
	go lb.startHealthCheck(done)
	fmt.Println("Processing requests....")
	// registering request
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", lbHandler(&lb))

	err := http.ListenAndServe(":6000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	time.Sleep(100 * time.Second)
	// closing done will stop our healthcheck
	// here we are taking down our server after 100 secs, so we close all channels as well.
	close(done)
}
