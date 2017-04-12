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
	serverIP:=config.GetMyIP("lo")
	serverPort:="8888"
	clientPort:="8080"
	var interval time.Duration=2
	var finishTime time.Duration=9
	message:="hello!"
	filePath:="/src/github.com/alruiz12/goREST/FileToSend"

	// server starts listening
	go func() {
		log.Fatal(http.ListenAndServe(":"+serverPort, router))
	}()

	// client starts sending
	go func() {
		var quit = make(chan int)
		client.StartSendingMessages(interval,serverIP,serverPort,message,quit)
		time.AfterFunc(finishTime * time.Second, func(){close(quit)})
	}()

	// client sends file
	client.SendFile(filePath,serverIP,serverPort)

	// client starts listening
	log.Fatal(http.ListenAndServe(":"+clientPort, router))
}
