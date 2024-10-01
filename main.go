package main

import (
	"log"
	"net"
	"time"

	"github.com/mgranderath/traceroute/methods"
	"github.com/mgranderath/traceroute/methods/tcp"
	"github.com/mgranderath/traceroute/methods/udp"
)

func main() {
	ip := net.ParseIP("75.2.98.97")
	tcpTraceroute := tcp.New(ip, methods.TracerouteConfig{
		MaxHops:          20,
		NumMeasurements:  3,
		ParallelRequests: 18,
		Port:             53,
		Timeout:          time.Second / 2,
	})
	res, err := tcpTraceroute.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	udpTraceroute := udp.New(ip, true, methods.TracerouteConfig{
		MaxHops:          20,
		NumMeasurements:  3,
		ParallelRequests: 24,
		Port:             784,
		Timeout:          2 * time.Second,
	})
	res, err = udpTraceroute.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
