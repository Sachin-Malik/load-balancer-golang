

package main

import (
  "strconv"
  "math/rand"
)

type LoadBalancer struct {
  servers []Server
}

func (lb *LoadBalancer) addServer() {
  newServer := Server{
    id:strconv.Itoa(rand.Intn(10000000)),
    capacity:5, 
    cache: map[string]string{}, 
    database: map[string]string{}, 
    connections:0,
  }
  lb.servers = append(lb.servers, newServer)
}

func (lb *LoadBalancer) sendRequest(req Request) string {
  var minConnectionServerIndex int = 0
  var minConnections int = 5;
  for index, server := range lb.servers{
    if(server.connections<minConnections){
      minConnections=server.connections
      minConnectionServerIndex=index
    }
  }
  result:=lb.servers[minConnectionServerIndex].processRequest(req);
  return result
}

func (lb *LoadBalancer) removeServer(serverId string) []Server{
  var newServerList []Server
  for _, item := range lb.servers {
    if(item.id!=serverId){
      newServerList = append(newServerList,item)
    }
  }
  return newServerList
}



