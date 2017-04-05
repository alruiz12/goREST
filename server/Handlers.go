package server
import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"

	"github.com/alruiz12/goREST/vars"
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


