package main

import (
	"log"
	"net"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	addressArg          = kingpin.Arg("address", "address to dial").Required().String()
	outstandingCountArg = kingpin.Arg("count", "number of outstanding connections").Required().Int()
	connectionPeriodArg = kingpin.Arg("period", "ms to wait between connection attempts").Default("1000").Int()
)

func main() {
	kingpin.Parse()

	if addressArg == nil || outstandingCountArg == nil || connectionPeriodArg == nil {
		log.Fatalln("Must specify all arguments")
	}

	address := *addressArg
	outstandingCount := *outstandingCountArg
	connectionPeriod := *connectionPeriodArg

	for i := 0; i < outstandingCount; i++ {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("connection failed: %+v\n", err)
			continue
		}
		readDest := make([]byte, 1024)
		go conn.Read(readDest)
		time.Sleep(time.Millisecond * time.Duration(connectionPeriod))
	}

	time.Sleep(time.Minute * 60)
}
