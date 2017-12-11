package main

import (
	"bytes"
	"log"
	"net"
	"fmt"
	
)
func check_result(tip string,err error)bool{
	if err := recover(); err != nil {
		log.Println(tip, err)
		return false
	}
	return true
}
func listenUDP() {
	log.Println("listening for stats UDP on port " + *udpListenAddress)
	serverAddr, err := net.ResolveUDPAddr("udp", *udpListenAddress)
	if err != nil {
		log.Println("Error: ", err)
	}

	serverConn, err := net.ListenUDP("udp", serverAddr)
	check_result("Error:",err)
    defer func() {  
        if err := recover(); err != nil {  
            fmt.Println("Critical: ",err) 
            fmt.Println("Panic,try to restart UDP Service") 
            go listenUDP()
        }  
    }()  
	defer serverConn.Close()

	buf := make([]byte, 8192)

	for {
		n, _, err := serverConn.ReadFromUDP(buf)
		if !check_result("Error reading from UDP: ",err){
			continue
		}
		udpPacketCount.Inc()
		if *debug {
			log.Printf("new udp package: %s", string(buf[ 0:n]))
		}
	    
		
		deltas, err := NewDelta(bytes.NewBuffer(buf[0:n]))
		if !check_result("Error creating delta: ",err){
			continue
		}

        for _,delta := range deltas{
			err = delta.Apply()
			if !check_result("Error applying delta: ",err){
				continue
			}
		}
	
	}
}
