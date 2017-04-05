package client
import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

/*
Router using gorilla/mux
*/
func NewClientRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	return router
}
