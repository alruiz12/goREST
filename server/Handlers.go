package server
import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"

)


/*
listenAnnounce is called when a POST requests 8080/listenAnnounce.
Writes down the current peer announcing so it avoids deleting it
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request
 */
func ListenMessage(w http.ResponseWriter, r *http.Request){
	var m Message
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}




















/*

func addTorrent(w http.ResponseWriter, r *http.Request){
	fmt.Println("... addTorrent STARTS ...")
	var t vars.Torrent
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &t); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	ret,err:=addTorrentRepo(t)
	if err!=nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(208) //Already reported
		if err := json.NewEncoder(w).Encode(ret); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		panic(err)
	}
	fmt.Println("... addTorrent FINISHES ...")
}
*/
/*
showTorrents is called when a GET requests 8080/addTorrent.
Sends new json encoded torrent back to the sender
@param1 used by an HTTP handler to construct an HTTP response.
@param2 represents HTTP request
 */
/*
func showTorrents(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vars.TorrentMap); err != nil {
		panic(err)
	}
}
*/
