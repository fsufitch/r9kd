package message

import "github.com/gorilla/mux"

// RegisterMessageRoutes registers the appropriate routes for REST channel endpoints
func RegisterMessageRoutes(router *mux.Router) {
	router.HandleFunc("/channel/{channel_id}/message", postMessage).Methods("POST")
}
