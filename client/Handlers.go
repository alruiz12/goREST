package client
import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"time"
	"strings"
	"errors"
)

func StartSendingMessages(interval time.Duration, IP string, port string,message string, quit chan int){
	ticker := time.NewTicker(interval * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				SendMessage(IP, port, message)

			case <-quit:
				fmt.Println("Stop sending messages")
				ticker.Stop()
				return
			}
		}
	}()

}



func SendMessage(IP string, port string, message string){
	var reader io.Reader
	trackerURL := "http://"+IP+":"+port+"/ListenMessage"
	jsonContent := `{"IP":"`+IP+`","port":"`+port+`","message":"`+message+`"}`
	reader = strings.NewReader(jsonContent)
	request, err := http.NewRequest("POST", trackerURL, reader)
	req, err := http.DefaultClient.Do(request)

	var serverAnswer string
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &serverAnswer);
	if err != nil {
		errors.New("error decoding")
	}
	fmt.Println("Server: ",req.Status)
}



