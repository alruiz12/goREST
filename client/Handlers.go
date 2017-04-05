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

/*
addTorrent is called when a POST requests 8080/addTorrent.
Adds a new Torrent to the Tracker file of torrents.
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request.
 */
func addTorrent(w http.ResponseWriter, r *http.Request){

}

/*
showTorrents is called when a GET requests 8080/addTorrent.
Sends new json encoded torrent back to the sender
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request
 */
func showTorrents(w http.ResponseWriter,r *http.Request) {

}


/*
upLoadFile is called when a POST requests 8080/upLoadFile.
Allow peer to upload a file
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request.
 */
func upLoadFile(w http.ResponseWriter, r *http.Request) {

}
