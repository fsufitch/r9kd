package channels

import "github.com/gorilla/mux"

// RegisterChannelRoutes registers the appropriate routes for REST channel endpoints
func RegisterChannelRoutes(router *mux.Router) {
	sub := router.PathPrefix("/channel").Subrouter()

	sub.HandleFunc("", postChannel).Methods("POST")
}
