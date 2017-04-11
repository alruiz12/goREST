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
	serverPort:="8888"
	clientPort:="8080"
	messsage:="hello!"
	var interval time.Duration=2
	var finishTime time.Duration=9

	// server starts listening
	go func() {
		log.Fatal(http.ListenAndServe(":"+serverPort, router))
	}()

	// client starts sending
	go func() {
		var quit = make(chan int)
		client.StartSendingMessages(interval,IP,serverPort,messsage,quit)
		time.AfterFunc(finishTime * time.Second, func(){close(quit)})
	}()

	// client starts listening
	log.Fatal(http.ListenAndServe(":"+clientPort, router))
}
