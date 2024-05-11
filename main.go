package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"sync"
)

type Request struct {
	time int
	requestId string
	payload string
}

func initLoadBalancer() LoadBalancer{
	var dummyServers []Server
	for i:=0;i<10;i++ {
		server := Server{
				capacity:5,
				cache:map[string]string{},
				database:map[string]string{},
				connections:0,
			}
		dummyServers = append(dummyServers,server);
		}
		lb := LoadBalancer{
			servers:dummyServers,
		}
	return lb
}
func main() {
	lb := initLoadBalancer()
	payload := []string {"Sachin","malik","nitin","jatin","ekta"}

	fmt.Println("Processing requests....");

	// if we dont use go routines, then we won't be able to process request concurrently
	// below example will take more than 5 seconds (each request X 1sec)


	/*###################################
	
	Uncomment the below lines to see 
	how it runs without conurrency
	
	####################################*/

	// start := time.Now()
	// var dummyRequest []Request
	// for i:=0;i<5;i++{
	// 	id := strconv.Itoa(rand.Intn(1000000))
	// 	req := Request{
	// 		time:1,
	// 		requestId: id,
	// 		payload:payload[i],
	// 	}
	// 	dummyRequest = append(dummyRequest,req);
	// }
	// for _, request := range dummyRequest {
	// 	func (req Request){
	// 		 lb.sendRequest(req);
	// 	}(request)
	// }
	// fmt.Println("all requests took", time.Since(start));

	// Here we are using go routines, so 5 request each taking 1 second will only take 1 second.
	start := time.Now()
	var dummyRequest []Request
	for i:=0;i<5;i++{
		id := strconv.Itoa(rand.Intn(1000000))
		req := Request{
			time:1,
			requestId: id,
			payload:payload[i],
		}
		dummyRequest = append(dummyRequest,req);
	}
	wg := sync.WaitGroup{}
	for _, request := range dummyRequest {
     wg.Add(1);
		go func (req Request){
			 defer wg.Done()
			 lb.sendRequest(req);
		}(request)
	}
	wg.Wait();
	fmt.Println("all requests took", time.Since(start));
}
