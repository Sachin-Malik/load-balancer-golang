package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type LoadBalancer struct {
	servers []Server
}

func initLoadBalancer() LoadBalancer {
	var dummyServers []Server
	for i := 0; i < 10; i++ {
		server := Server{
			capacity:    5,
			cache:       map[string]string{},
			database:    map[string]string{},
			connections: 0,
		}
		dummyServers = append(dummyServers, server)
	}
	lb := LoadBalancer{
		servers: dummyServers,
	}
	return lb
}

func (lb *LoadBalancer) startHealthCheck(done <-chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return

		case <-ticker.C:
			for serverIndex, server := range lb.servers {
				isOnline := server.healthCheck()
				if !isOnline {
					fmt.Println("Server at ", serverIndex, " is offline Updating available server List")
					lb.servers = lb.removeServer(serverIndex)
					// adding a new server, with default Configs
					lb.addServer()
					fmt.Println("Number of server online ", len(lb.servers))
				}
			}
		}
	}
}

func (lb *LoadBalancer) addServer() {
	newServer := Server{
		id:          strconv.Itoa(rand.Intn(10000000)),
		capacity:    5,
		cache:       map[string]string{},
		database:    map[string]string{},
		connections: 0,
	}
	lb.servers = append(lb.servers, newServer)
}

func (lb *LoadBalancer) sendRequest(req *MyRequest) string {
	if len(lb.servers) == 0 {
		return "No server Online"
	}
	var minConnectionServerIndex int = 0
	var minConnections int = 5
	for index, server := range lb.servers {
		if server.connections < minConnections {
			minConnections = server.connections
			minConnectionServerIndex = index
		}
	}
	result := lb.servers[minConnectionServerIndex].processRequest(req)

	return result
}

func (lb *LoadBalancer) removeServer(serverIndex int) []Server {
	var newServerList []Server
	for index, item := range lb.servers {
		if index != serverIndex {
			newServerList = append(newServerList, item)
		}
	}
	return newServerList
}
