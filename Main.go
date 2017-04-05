package main

import (
	"time"
	"github.com/alruiz12/goREST/client"
	"net/http"
	"log"
	"github.com/alruiz12/goREST/config"

	"github.com/alruiz12/goREST/server"
)

func main() {

	router := server.NewServerRouter()
	IP:=config.GetMyIP("lo")

	// server starts listening
	go func() {
		log.Fatal(http.ListenAndServe(":8888", router))
	}()

	// client starts sending
	go func() {
		var quit = make(chan int)
		client.StartSendingMessages(2,IP,"8888","hello!",quit)
		time.AfterFunc(9 * time.Second, func(){close(quit)})
	}()

	// client starts listening
	log.Fatal(http.ListenAndServe(":8080", router))
}
