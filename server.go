
package main

import (
  "math/rand"
  "time"
  "fmt"
)

type Server struct {
  id string
  capacity    int
  cache       map[string]string
  database    map[string]string
  connections int
}

// only a server struct can call this function
func (s *Server) serverHeathCheck() bool {
  var isOnline bool = rand.Intn(10) > 8
  return isOnline
}

func (s *Server) processRequest(req Request) string {
  s.connections++;
  if _, exists := s.cache[req.payload]; exists {
    return s.cache[req.payload]
  }else{
    time.Sleep(time.Second)
    s.database[req.requestId]=req.payload
    fmt.Println("Inserted in DB", req.payload)
  }
  s.connections--;
  return req.payload;
}
