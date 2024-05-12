package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Server struct {
	id          string
	capacity    int
	cache       map[string]string
	database    map[string]string
	connections int
}

// only a server struct can call this function
func (s *Server) healthCheck() bool {
	var isOnline bool = rand.Intn(10) < 8
	return isOnline
}

func (s *Server) processRequest(req *MyRequest) string {
	s.connections++
	if _, exists := s.cache[req.payload]; exists {
		return s.cache[req.payload]
	} else {
		time.Sleep(time.Second)
		key := strconv.Itoa(rand.Intn(10000))
		s.database[key] = req.payload
		fmt.Println("Inserted ", req.payload, " in DB at ", key)
	}
	s.connections--
	return req.payload
}
