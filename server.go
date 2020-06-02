package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	portArg = kingpin.Arg("port", "port to use").Required().Int()
)

func main() {
	kingpin.Parse()
	if portArg == nil {
		log.Fatalln("Must specify a port")
	}
	port := *portArg

	// Log, but do not (explicitly) respond to any caught signals
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGKILL, os.Interrupt)
	go func() {
		for {
			select {
			case sig := <-sigs:
				fmt.Println("Received signal:", sig)
			}
		}
	}()

	if err := server(port); err != nil {
		log.Fatalln(err)
	}
}

func server(port int) error {
	http.HandleFunc("/", HelloServer)
	log.Printf("Starting server at :%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
