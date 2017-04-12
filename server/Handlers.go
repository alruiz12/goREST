package server
import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/alruiz12/goREST/vars"
	"log"
	"fmt"
	"os"
)


/*
ListenMessage is called when a POST requests ServerIP:ServerPort/ListenMessage.
Writes down the current message to list of Messages
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request
 */
func ListenMessage(w http.ResponseWriter, r *http.Request){
	var m vars.Message
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &m); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	vars.MessageSlice=append(vars.MessageSlice,m)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}



/*
ShowMessages is called when a GET requests ServerIP:ServerPort/ShowMessages.
Sends new json encoded message back to the requester
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request
 */
func ShowMessages(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vars.MessageSlice); err != nil {
		panic(err)
	}
}


/*
receiveFile is called when a POST requests serverPort/upLoadFile.
Allows peer to send a file to server
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request.
 */
func receiveFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** addFile STARTS ***")
	var file string
	if r.Method == http.MethodPost{
		f, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error opening form file")
			log.Println(err)
			http.Error(w, "Error uploading file", 404)
			return
		}
		//if (Exists("../uploadedFiles/"+header.Filename)){ }
		defer f.Close()
		fileName:=header.Filename
		destination, err := os.Create(os.Getenv("GOPATH")+"/src/github.com/alruiz12/goREST/receivedFiles/"+fileName)
		if err != nil {
			http.Error(w,err.Error(), 501) //internal server error
			return
		}
		defer destination.Close()
		io.Copy(destination,f)

		body, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		//file filled  with body
		file = string(body)

	}
	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintln(w,`
	<form action="/upLoadFile" method="post" enctype="multipart/form-data">
	    upload a file<br>
	    <input type="file" name="file"><br>
	    <input type="submit">
	</form>
	<br>
	<br>
	`,file)

	fmt.Println("*** addFile FINISHES ***")
}




