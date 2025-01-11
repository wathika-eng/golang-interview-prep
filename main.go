package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/matthewjamesboyle/golang-interview-prep/cmd/api"
)

func main() {
	log.Println(time.Now().Format("13:01:46"))
	// Start the server in a goroutine
	go api.StartServer()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	log.Println("Server shutting down...")
}
