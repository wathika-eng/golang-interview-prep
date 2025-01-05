package main

import (
	"log"
	"time"

	"github.com/matthewjamesboyle/golang-interview-prep/cmd/api"
)

func main() {
	log.Println(time.Now().Format("13:01:46"))
	api.StartServer()
}
