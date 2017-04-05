package config
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/alruiz12/goREST/server"
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
func MyNewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"ListenMessage",
		"POST",
		"/ListenMessage",
		server.ListenMessage,
	},
	Route{
		"showMessages",
		"GET",
		"/showMessages",
		server.ShowMessages,
	},

}