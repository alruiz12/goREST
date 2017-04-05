package client

import (
	"testing"
	"github.com/alruiz12/goREST/config"
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/alruiz12/goREST/server"
)

func TestClient(t *testing.T) {

	svRouter := server.NewServerRouter()
	IP:=config.GetMyIP("lo")
	fmt.Println(IP)
	// server starts listening
	srv:=&http.Server{Addr: ":8888", Handler:svRouter}
	go func(){
		if err := srv.ListenAndServe(); err!=nil{
			log.Printf("ListenAndServe error", err)
		}
	}()

	time.Sleep(3*time.Second)

	// client starts sending
	go func() {
		var quit = make(chan int)
		StartSendingMessages(2,IP,"8888","hello!",quit)
		time.AfterFunc(9 * time.Second, func() {close(quit)})
	}()

	// client starts listening
	cliRouter := NewClientRouter()
	client:=&http.Server{Addr: ":8080", Handler:cliRouter}
	go func() {
		if err := client.ListenAndServe(); err!=nil{
			log.Printf("ListenAndServe error", err)
		}
	}()

	time.Sleep(3*time.Second)

	incomingURL:=  "http://"+IP+":8888/showMessages"
	fmt.Println(incomingURL)
	request, err := http.NewRequest("GET", incomingURL, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("Failure expected: %d", res.StatusCode)
	}



	time.AfterFunc(19 * time.Second, func(){
		if err:= srv.Shutdown(nil); err!=nil{
			panic(err)
		}
		if err:= client.Shutdown(nil); err!=nil{
			panic(err)
		}
	})

	//avoid ending execution until shutdowns
	time.Sleep(22*time.Second)







}
