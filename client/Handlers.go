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
	"bytes"
	"mime/multipart"
	"log"
	"os"
	"path"
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


func SendFile(filePath string, IP string, port string){
	file, err := os.Open(os.Getenv("GOPATH")+filePath)
	if err != nil {
		fmt.Println("Opening file")
		log.Println(err)
	}
	defer file.Close()
	destinationURL:="http://"+IP+":"+port+"/receiveFile"
	fmt.Println(destinationURL)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", path.Base(filePath))
	if err != nil {
		fmt.Println("creating Form file")
		log.Println(err)
	}
	_, err = io.Copy(part, file)
	err=writer.Close()
	if err != nil {
		log.Println(err)
	}
	request, err := http.NewRequest("POST", destinationURL, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	if res.StatusCode != 200 {
		log.Println("Success expected: %d", res.StatusCode)
	}


}





